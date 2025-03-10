import * as vm from 'vm';
import * as fs from 'fs';
import path from "path";
import {transformAsync} from "@babel/core";
import {patchCode} from "./ast_utils";
import {listFiles} from "./file_utils";
import {execSync} from "node:child_process";
import {createRequire} from "module";

async function transpile(code: string, filename: string): Promise<string> {
    const result = await transformAsync(code, {
        presets: ['@babel/preset-env', '@babel/preset-typescript'],
        filename: filename,
        minified: false,
        compact: false,
    });

    if (!result?.code) {
        throw new Error('Failed to transpile TypeScript code.');
    }

    return result.code;
}

async function patchDir(dir: string): Promise<void> {
    const files = await listFiles(dir)
    for (const file of files) {
        if (file.includes("/node_modules/")) {
            continue
        }

        if (!file.endsWith(".js") && !file.endsWith(".ts")) {
            continue
        }
        let code = await fs.promises.readFile(file, "utf8");
        try {
            code = await patchCode(code)
            if (file.endsWith(".ts")) {
                code = await transpile(code, file);
            }
            let newFile = file.replace(".ts", ".js")
            await fs.promises.writeFile(newFile, code, 'utf-8');
        } catch (err) {
            console.error(err);
        }
    }
}


export interface Context {
    [key: string]: any; // Allow dynamic properties
}

export class Sandbox {
    context: Context
    codeDir: string
    ak_call: Function

    constructor(codeDir: string, ak_call: Function) {
        this.codeDir = codeDir;
        this.context = {};
        this.ak_call = ak_call;
        // this.prepareCodeDir();
        this.initContext()
    }

    prepareCodeDir(): void {
        let output = execSync(`cd ${this.codeDir}; npm install`).toString();
        console.log(output);
    }

    setCodeDir(codeDir: string): void {
        this.codeDir = codeDir;
        this.prepareCodeDir();
    }

    initContext() {
        this.context.exports = {}
        this.context.ak_call = this.ak_call
        this.context.workingDir = "."
        this.context.console = {
            log: console.log,
        }


        this.context.process = {
            env: process.env,
            cwd: () => this.codeDir
        }

        this.context.require = (moduleName: string) => {

            if (moduleName.startsWith(".")) {
                moduleName = moduleName.endsWith(".js") ? moduleName : moduleName + ".js"

                const previousExecutingFile = this.context.__currentExecutingFile;
                const callerPath = previousExecutingFile || this.codeDir;
                moduleName = path.resolve(path.dirname(callerPath), moduleName);
                const code = fs.readFileSync(moduleName, "utf8")
                this.context.__currentExecutingFile = moduleName;

                this.context.__filename = moduleName;
                this.context.__dirname = path.dirname(moduleName);
                let o = vm.runInContext(code,this.context);
                if (typeof o === "function") {
                    o._ak_direct_call = true
                }
                return this.context
            }
            const customRequire = createRequire(path.resolve(this.codeDir, "package.json"));

            let mod: any
            console.log("resolving " + moduleName)
            try {
                mod = require(moduleName)
            } catch (e) {
                // mod =  require(`${this.codeDir}/node_modules/${moduleName}`)
                // let modPath = path.resolve(`${this.codeDir}/node_modules/${moduleName}`)
                // mod = require(modPath)
                // mod =  require(`${this.codeDir}/node_modules/${moduleName}`)
                mod = customRequire(moduleName);
            }
            return mod
        }
        vm.createContext(this.context);
    }

    async loadFile(filePath: string): Promise<void> {
        filePath = path.join(this.codeDir, filePath)
        await patchDir(path.dirname(filePath))
        filePath = filePath.replace(".ts", ".js");
        let code = fs.readFileSync(filePath, 'utf8');
        this.context.__currentExecutingFile = filePath;

        this.context.__filename = filePath;
        this.context.__dirname = path.dirname(filePath);
        vm.runInContext(code,this.context);
    }

    async run(f: string, args: any, callback: Function): Promise<any> {
        let code = `(async () => {
            return await ${f}(...[${JSON.stringify(args)}]);
        })();`
        const results = await vm.runInContext(code, this.context)
        callback(results)
    }
}
