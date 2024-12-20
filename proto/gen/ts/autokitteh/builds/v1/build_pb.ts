// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/builds/v1/build.proto (package autokitteh.builds.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message autokitteh.builds.v1.Build
 */
export class Build extends Message<Build> {
  /**
   * @generated from field: string build_id = 1;
   */
  buildId = "";

  /**
   * @generated from field: string project_id = 2;
   */
  projectId = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 3;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<Build>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.builds.v1.Build";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "build_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Build {
    return new Build().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Build {
    return new Build().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Build {
    return new Build().fromJsonString(jsonString, options);
  }

  static equals(a: Build | PlainMessage<Build> | undefined, b: Build | PlainMessage<Build> | undefined): boolean {
    return proto3.util.equals(Build, a, b);
  }
}

