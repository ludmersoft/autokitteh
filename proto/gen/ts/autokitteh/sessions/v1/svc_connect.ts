// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/sessions/v1/svc.proto (package autokitteh.sessions.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteRequest, DeleteResponse, GetLogRequest, GetLogResponse, GetRequest, GetResponse, ListRequest, ListResponse, ListSessionLogRecordsRequest, ListSessionLogRecordsResponse, StartRequest, StartResponse, StopRequest, StopResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autokitteh.sessions.v1.SessionsService
 */
export const SessionsService = {
  typeName: "autokitteh.sessions.v1.SessionsService",
  methods: {
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.Start
     */
    start: {
      name: "Start",
      I: StartRequest,
      O: StartResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.Stop
     */
    stop: {
      name: "Stop",
      I: StopRequest,
      O: StopResponse,
      kind: MethodKind.Unary,
    },
    /**
     * List returns events without their data.
     *
     * @generated from rpc autokitteh.sessions.v1.SessionsService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.GetLog
     */
    getLog: {
      name: "GetLog",
      I: GetLogRequest,
      O: GetLogResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.ListSessionLogRecords
     */
    listSessionLogRecords: {
      name: "ListSessionLogRecords",
      I: ListSessionLogRecordsRequest,
      O: ListSessionLogRecordsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.sessions.v1.SessionsService.Delete
     */
    delete: {
      name: "Delete",
      I: DeleteRequest,
      O: DeleteResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

