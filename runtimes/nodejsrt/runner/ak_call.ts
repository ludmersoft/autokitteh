import {randomUUID} from "node:crypto";
import {EventEmitter, once} from "node:events";
import { Client } from "@connectrpc/connect";
import {HandlerService} from "./pb/autokitteh/user_code/v1/handler_svc_pb";


export interface Waiter {
    wait:  (f: Function, v: any, token: string) => Promise<any>
    execute_signal: (token: string) => Promise<any>
    replay_signal: (token: string, value: any) => Promise<void>
}

export class ActivityWaiter implements Waiter{
    event: EventEmitter;
    f: Function
    a: any
    token: string
    client: Client<typeof HandlerService>
    runnerId: string

    constructor(client: Client<typeof HandlerService>, runnerId: string) {
        this.client = client;
        this.event = new EventEmitter();
        this.f = () => {}
        this.token = ""
        this.runnerId = runnerId;
    }

    async execute_signal(token: string): Promise<any> {
        if (token != this.token) {
            throw new Error('tokens do not match')
        }

        const v = await this.f(...this.a)
        this.event.emit('return', v);
        return v
    }

    async replay_signal(token: string, value: any): Promise<void> {
        if (token != this.token) {
            throw new Error('tokens do not match')
        }

        this.event.emit('return', value);
    }

    async wait(f: Function, v: any, token: string): Promise<any> {
        this.f = f
        this.a = v
        this.token = token

        await this.client.activity({})

        if (v[0].runnerId != undefined) {
            this.runnerId = v[0].runnerId
        }

        const encoder = new TextEncoder()
        const resp = await this.client.activity({
            runnerId: this.runnerId,
            data: encoder.encode(JSON.stringify({token: token})),
            callInfo: {
                function: f.name,
                args: [],
            }
        });
        console.log("activity resp", resp)
        const r = (await once(this.event, 'return'))[0]
        console.log("got return value", r)
        return r
    }
}

export const ak_call = (waiter: Waiter) => {
    return async (...args: any) => {
        let f = args[0];
        let f_args: any = []
        if (args.length > 1) {
            f_args = args.slice(1);
        }


        if (f.ak_call == false) {
            return await f(...f_args);
        }

        const results = await waiter.wait(f, f_args, randomUUID());
        console.log("got results", results)
        return results;
    }
}
