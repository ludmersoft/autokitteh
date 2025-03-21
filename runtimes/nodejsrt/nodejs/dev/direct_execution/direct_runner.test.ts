// import Runner from "../../runtime/runner";
import Runner from "../../testdata/simple-test/runtime/.ak/runtime/runner";
import { DirectHandlerClient } from "./direct_client";
import * as path from "path";
import fs from "fs";
import { TextEncoder, TextDecoder } from "util";
import { test } from '@jest/globals';

test('full flow', async () => {
  console.log("Starting direct runner test...");

  // Create absolute path to the test directory
  const testDir = path.resolve('testdata/simple-test/runtime');
  // const testDir = path.resolve('examples/simple-test');
  const entryPoint = 'test.ts:test';
  if (!fs.existsSync(testDir)) {
    throw new Error(`Test dir not found at ${testDir}`);
  }

  console.log(`Using absolute test directory: ${testDir}`);

  // Create direct client - no need for separate waiter
  const directClient = new DirectHandlerClient();
  console.log("Created DirectHandlerClient");

  // Create runner with direct client - it will create its own ActivityWaiter internally
  console.log("Creating runner...");
  const runner = new Runner(
      "test-runner-id",
      testDir,
      directClient as any
  );
  console.log("Runner created");

  // Start the runner
  console.log("Starting runner...");
  await runner.start();
  console.log("Runner started");

  // Create a mock router to capture the service implementation
  console.log("Creating service...");
  const mockServiceImplementation: any = {};
  const mockRouter = {
    service: (service: any, implementation: any) => {
      console.log("Capturing service implementation");
      Object.assign(mockServiceImplementation, implementation);
      return mockRouter;
    }
  };

  // Register the runner service on our mock router using a type assertion to bypass type checking
  console.log("Registering runner service...");
  (runner.createService() as any)(mockRouter);
  console.log("Available runner service methods:", Object.keys(mockServiceImplementation));

  // Connect the runner service back to the direct client
  console.log("Connecting runner service to direct client...");
  directClient.setRunnerService(mockServiceImplementation);
  console.log("Service connected");

  // Emit the started event to complete initialization
  console.log("Emitting started event...");
  try {
    // Try accessing the events property safely
    const runnerAny = runner as any;
    if (runnerAny.events && typeof runnerAny.events.emit === 'function') {
      runnerAny.events.emit('started');
      console.log("Started event emitted");
    } else {
      console.log("Could not access events property or emit method");
    }
  } catch (err) {
    console.log("Could not emit started event, continuing anyway:", err);
  }

  // Create a simple event object with data
  const encoder = new TextEncoder();
  const eventData = {
    token: "test-token",
    args: ["TestUser"]
  };
  console.log("Created event data:", eventData);

  // Create the start request with entry point and event
  console.log("Executing test function...");
  const startRequest = {
    entryPoint: entryPoint,
    event: {
      data: encoder.encode(JSON.stringify(eventData))
    }
  };
  console.log("Created start request with entry point:", startRequest.entryPoint);

  // Add timeout in case the call hangs
  console.log("Calling start method directly on client...");
  const result = await directClient.start(startRequest);

  // Log the result
  console.log("Execution result:", result);

  // Clean up the runner
  // console.log("Stopping runner...");
  // runner.stop();
  // console.log("Runner stopped");

  // Set the client's active state to false so the runner can exit
  console.log("Setting client isActive to false...");
  directClient.setActive(false);
  console.log("Client isActive set to false");

  // // Force exit any pending health check intervals
  // console.log("Forcing health check cleanup...");
  // const runnerAny = runner as any;
  // if (runnerAny.healthcheckTimer) {
  //   clearInterval(runnerAny.healthcheckTimer);
  //   runnerAny.healthcheckTimer = undefined;
  //   console.log("Health check timer cleared");
  // }
  //
  // // Set internal flag to prevent further health checks
  // if (typeof runnerAny._isRunningHealthCheck !== 'undefined') {
  //   runnerAny._isRunningHealthCheck = false;
  //   console.log("Running health check flag set to false");
  // }

  console.log("Test completed successfully!");
});
