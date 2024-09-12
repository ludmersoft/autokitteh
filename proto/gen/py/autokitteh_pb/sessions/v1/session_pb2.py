# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/sessions/v1/session.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from autokitteh_pb.program.v1 import program_pb2 as autokitteh_dot_program_dot_v1_dot_program__pb2
from autokitteh_pb.values.v1 import values_pb2 as autokitteh_dot_values_dot_v1_dot_values__pb2
from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$autokitteh/sessions/v1/session.proto\x12\x16\x61utokitteh.sessions.v1\x1a#autokitteh/program/v1/program.proto\x1a!autokitteh/values/v1/values.proto\x1a\x1b\x62uf/validate/validate.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x88\x07\n\x0cSessionState\x12\x46\n\x07\x63reated\x18\n \x01(\x0b\x32,.autokitteh.sessions.v1.SessionState.CreatedR\x07\x63reated\x12\x46\n\x07running\x18\x0b \x01(\x0b\x32,.autokitteh.sessions.v1.SessionState.RunningR\x07running\x12@\n\x05\x65rror\x18\x0c \x01(\x0b\x32*.autokitteh.sessions.v1.SessionState.ErrorR\x05\x65rror\x12L\n\tcompleted\x18\r \x01(\x0b\x32..autokitteh.sessions.v1.SessionState.CompletedR\tcompleted\x12\x46\n\x07stopped\x18\x0e \x01(\x0b\x32,.autokitteh.sessions.v1.SessionState.StoppedR\x07stopped\x1a\t\n\x07\x43reated\x1aZ\n\x07Running\x12\x1e\n\x06run_id\x18\x01 \x01(\tB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x05runId\x12/\n\x04\x63\x61ll\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x04\x63\x61ll\x1a\\\n\x05\x45rror\x12\x16\n\x06prints\x18\x01 \x03(\tR\x06prints\x12;\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x1c.autokitteh.program.v1.ErrorB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x05\x65rror\x1a\xa7\x02\n\tCompleted\x12\x16\n\x06prints\x18\x01 \x03(\tR\x06prints\x12i\n\x07\x65xports\x18\x02 \x03(\x0b\x32;.autokitteh.sessions.v1.SessionState.Completed.ExportsEntryB\x12\xfa\xf7\x18\x0e\x9a\x01\x0b\"\x04r\x02\x10\x01*\x03\xc8\x01\x01R\x07\x65xports\x12>\n\x0creturn_value\x18\x03 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x0breturnValue\x1aW\n\x0c\x45xportsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value:\x02\x38\x01\x1a!\n\x07Stopped\x12\x16\n\x06reason\x18\x01 \x01(\tR\x06reason\"\xc3\x08\n\x04\x43\x61ll\x12>\n\x04spec\x18\x01 \x01(\x0b\x32!.autokitteh.sessions.v1.Call.SpecB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x04spec\x12N\n\x08\x61ttempts\x18\x02 \x03(\x0b\x32$.autokitteh.sessions.v1.Call.AttemptB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x08\x61ttempts\x1a\xcc\x02\n\x04Spec\x12@\n\x08\x66unction\x18\x01 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x08\x66unction\x12=\n\x04\x61rgs\x18\x02 \x03(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x04\x61rgs\x12Y\n\x06kwargs\x18\x03 \x03(\x0b\x32-.autokitteh.sessions.v1.Call.Spec.KwargsEntryB\x12\xfa\xf7\x18\x0e\x9a\x01\x0b\"\x04r\x02\x10\x01*\x03\xc8\x01\x01R\x06kwargs\x12\x10\n\x03seq\x18\x04 \x01(\rR\x03seq\x1aV\n\x0bKwargsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value:\x02\x38\x01\x1a\xdb\x04\n\x07\x41ttempt\x12@\n\x05start\x18\x01 \x01(\x0b\x32*.autokitteh.sessions.v1.Call.Attempt.StartR\x05start\x12I\n\x08\x63omplete\x18\x02 \x01(\x0b\x32-.autokitteh.sessions.v1.Call.Attempt.CompleteR\x08\x63omplete\x1ao\n\x06Result\x12\x31\n\x05value\x18\n \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value\x12\x32\n\x05\x65rror\x18\x0b \x01(\x0b\x32\x1c.autokitteh.program.v1.ErrorR\x05\x65rror\x1a]\n\x05Start\x12\x42\n\nstarted_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\tstartedAt\x12\x10\n\x03num\x18\x05 \x01(\rR\x03num\x1a\xf2\x01\n\x08\x43omplete\x12\x46\n\x0c\x63ompleted_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x0b\x63ompletedAt\x12@\n\x0eretry_interval\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\rretryInterval\x12\x17\n\x07is_last\x18\x03 \x01(\x08R\x06isLast\x12\x43\n\x06result\x18\x04 \x01(\x0b\x32+.autokitteh.sessions.v1.Call.Attempt.ResultR\x06result\"\x9e\x06\n\x10SessionLogRecord\x12(\n\x01t\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x01t\x12\x1d\n\nprocess_id\x18\x02 \x01(\tR\tprocessId\x12\x44\n\x05print\x18\n \x01(\x0b\x32..autokitteh.sessions.v1.SessionLogRecord.PrintR\x05print\x12>\n\tcall_spec\x18\x0b \x01(\x0b\x32!.autokitteh.sessions.v1.Call.SpecR\x08\x63\x61llSpec\x12X\n\x12\x63\x61ll_attempt_start\x18\x0c \x01(\x0b\x32*.autokitteh.sessions.v1.Call.Attempt.StartR\x10\x63\x61llAttemptStart\x12\x61\n\x15\x63\x61ll_attempt_complete\x18\r \x01(\x0b\x32-.autokitteh.sessions.v1.Call.Attempt.CompleteR\x13\x63\x61llAttemptComplete\x12:\n\x05state\x18\x0e \x01(\x0b\x32$.autokitteh.sessions.v1.SessionStateR\x05state\x12W\n\x0cstop_request\x18\x0f \x01(\x0b\x32\x34.autokitteh.sessions.v1.SessionLogRecord.StopRequestR\x0bstopRequest\x1a\x1b\n\x05Print\x12\x12\n\x04text\x18\x01 \x01(\tR\x04text\x1a%\n\x0bStopRequest\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\"\xa4\x01\n\x04Type\x12\x14\n\x10TYPE_UNSPECIFIED\x10\x00\x12\x0e\n\nTYPE_PRINT\x10\x01\x12\x12\n\x0eTYPE_CALL_SPEC\x10\x02\x12\x1b\n\x17TYPE_CALL_ATTEMPT_START\x10\x04\x12\x1e\n\x1aTYPE_CALL_ATTEMPT_COMPLETE\x10\x08\x12\x0e\n\nTYPE_STATE\x10\x10\x12\x15\n\x11TYPE_STOP_REQUEST\x10 \"^\n\nSessionLog\x12P\n\x07records\x18\x01 \x03(\x0b\x32(.autokitteh.sessions.v1.SessionLogRecordB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x07records\"\xf3\x05\n\x07Session\x12\x1d\n\nsession_id\x18\x01 \x01(\tR\tsessionId\x12\x19\n\x08\x62uild_id\x18\x02 \x01(\tR\x07\x62uildId\x12\x15\n\x06\x65nv_id\x18\x03 \x01(\tR\x05\x65nvId\x12L\n\nentrypoint\x18\x04 \x01(\x0b\x32#.autokitteh.program.v1.CodeLocationB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\nentrypoint\x12W\n\x06inputs\x18\x05 \x03(\x0b\x32+.autokitteh.sessions.v1.Session.InputsEntryB\x12\xfa\xf7\x18\x0e\x9a\x01\x0b\"\x04r\x02\x10\x01*\x03\xc8\x01\x01R\x06inputs\x12*\n\x11parent_session_id\x18\x06 \x01(\tR\x0fparentSessionId\x12=\n\x04memo\x18\x07 \x03(\x0b\x32).autokitteh.sessions.v1.Session.MemoEntryR\x04memo\x12\x39\n\ncreated_at\x18\n \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12>\n\x05state\x18\x0c \x01(\x0e\x32(.autokitteh.sessions.v1.SessionStateTypeR\x05state\x12#\n\rdeployment_id\x18\x14 \x01(\tR\x0c\x64\x65ploymentId\x12\x19\n\x08\x65vent_id\x18\x15 \x01(\tR\x07\x65ventId\x1aV\n\x0bInputsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value:\x02\x38\x01\x1a\x37\n\tMemoEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01*\xd6\x01\n\x10SessionStateType\x12\"\n\x1eSESSION_STATE_TYPE_UNSPECIFIED\x10\x00\x12\x1e\n\x1aSESSION_STATE_TYPE_CREATED\x10\x01\x12\x1e\n\x1aSESSION_STATE_TYPE_RUNNING\x10\x02\x12\x1c\n\x18SESSION_STATE_TYPE_ERROR\x10\x03\x12 \n\x1cSESSION_STATE_TYPE_COMPLETED\x10\x04\x12\x1e\n\x1aSESSION_STATE_TYPE_STOPPED\x10\x05\x42\xf1\x01\n\x1a\x63om.autokitteh.sessions.v1B\x0cSessionProtoP\x01ZKgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/sessions/v1;sessionsv1\xa2\x02\x03\x41SX\xaa\x02\x16\x41utokitteh.Sessions.V1\xca\x02\x16\x41utokitteh\\Sessions\\V1\xe2\x02\"Autokitteh\\Sessions\\V1\\GPBMetadata\xea\x02\x18\x41utokitteh::Sessions::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.sessions.v1.session_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\032com.autokitteh.sessions.v1B\014SessionProtoP\001ZKgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/sessions/v1;sessionsv1\242\002\003ASX\252\002\026Autokitteh.Sessions.V1\312\002\026Autokitteh\\Sessions\\V1\342\002\"Autokitteh\\Sessions\\V1\\GPBMetadata\352\002\030Autokitteh::Sessions::V1'
  _SESSIONSTATE_RUNNING.fields_by_name['run_id']._options = None
  _SESSIONSTATE_RUNNING.fields_by_name['run_id']._serialized_options = b'\372\367\030\003\310\001\001'
  _SESSIONSTATE_ERROR.fields_by_name['error']._options = None
  _SESSIONSTATE_ERROR.fields_by_name['error']._serialized_options = b'\372\367\030\003\310\001\001'
  _SESSIONSTATE_COMPLETED_EXPORTSENTRY._options = None
  _SESSIONSTATE_COMPLETED_EXPORTSENTRY._serialized_options = b'8\001'
  _SESSIONSTATE_COMPLETED.fields_by_name['exports']._options = None
  _SESSIONSTATE_COMPLETED.fields_by_name['exports']._serialized_options = b'\372\367\030\016\232\001\013\"\004r\002\020\001*\003\310\001\001'
  _CALL_SPEC_KWARGSENTRY._options = None
  _CALL_SPEC_KWARGSENTRY._serialized_options = b'8\001'
  _CALL_SPEC.fields_by_name['function']._options = None
  _CALL_SPEC.fields_by_name['function']._serialized_options = b'\372\367\030\003\310\001\001'
  _CALL_SPEC.fields_by_name['args']._options = None
  _CALL_SPEC.fields_by_name['args']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _CALL_SPEC.fields_by_name['kwargs']._options = None
  _CALL_SPEC.fields_by_name['kwargs']._serialized_options = b'\372\367\030\016\232\001\013\"\004r\002\020\001*\003\310\001\001'
  _CALL_ATTEMPT_START.fields_by_name['started_at']._options = None
  _CALL_ATTEMPT_START.fields_by_name['started_at']._serialized_options = b'\372\367\030\003\310\001\001'
  _CALL_ATTEMPT_COMPLETE.fields_by_name['completed_at']._options = None
  _CALL_ATTEMPT_COMPLETE.fields_by_name['completed_at']._serialized_options = b'\372\367\030\003\310\001\001'
  _CALL.fields_by_name['spec']._options = None
  _CALL.fields_by_name['spec']._serialized_options = b'\372\367\030\003\310\001\001'
  _CALL.fields_by_name['attempts']._options = None
  _CALL.fields_by_name['attempts']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _SESSIONLOG.fields_by_name['records']._options = None
  _SESSIONLOG.fields_by_name['records']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _SESSION_INPUTSENTRY._options = None
  _SESSION_INPUTSENTRY._serialized_options = b'8\001'
  _SESSION_MEMOENTRY._options = None
  _SESSION_MEMOENTRY._serialized_options = b'8\001'
  _SESSION.fields_by_name['entrypoint']._options = None
  _SESSION.fields_by_name['entrypoint']._serialized_options = b'\372\367\030\003\310\001\001'
  _SESSION.fields_by_name['inputs']._options = None
  _SESSION.fields_by_name['inputs']._serialized_options = b'\372\367\030\016\232\001\013\"\004r\002\020\001*\003\310\001\001'
  _globals['_SESSIONSTATETYPE']._serialized_start=3887
  _globals['_SESSIONSTATETYPE']._serialized_end=4101
  _globals['_SESSIONSTATE']._serialized_start=231
  _globals['_SESSIONSTATE']._serialized_end=1135
  _globals['_SESSIONSTATE_CREATED']._serialized_start=607
  _globals['_SESSIONSTATE_CREATED']._serialized_end=616
  _globals['_SESSIONSTATE_RUNNING']._serialized_start=618
  _globals['_SESSIONSTATE_RUNNING']._serialized_end=708
  _globals['_SESSIONSTATE_ERROR']._serialized_start=710
  _globals['_SESSIONSTATE_ERROR']._serialized_end=802
  _globals['_SESSIONSTATE_COMPLETED']._serialized_start=805
  _globals['_SESSIONSTATE_COMPLETED']._serialized_end=1100
  _globals['_SESSIONSTATE_COMPLETED_EXPORTSENTRY']._serialized_start=1013
  _globals['_SESSIONSTATE_COMPLETED_EXPORTSENTRY']._serialized_end=1100
  _globals['_SESSIONSTATE_STOPPED']._serialized_start=1102
  _globals['_SESSIONSTATE_STOPPED']._serialized_end=1135
  _globals['_CALL']._serialized_start=1138
  _globals['_CALL']._serialized_end=2229
  _globals['_CALL_SPEC']._serialized_start=1291
  _globals['_CALL_SPEC']._serialized_end=1623
  _globals['_CALL_SPEC_KWARGSENTRY']._serialized_start=1537
  _globals['_CALL_SPEC_KWARGSENTRY']._serialized_end=1623
  _globals['_CALL_ATTEMPT']._serialized_start=1626
  _globals['_CALL_ATTEMPT']._serialized_end=2229
  _globals['_CALL_ATTEMPT_RESULT']._serialized_start=1778
  _globals['_CALL_ATTEMPT_RESULT']._serialized_end=1889
  _globals['_CALL_ATTEMPT_START']._serialized_start=1891
  _globals['_CALL_ATTEMPT_START']._serialized_end=1984
  _globals['_CALL_ATTEMPT_COMPLETE']._serialized_start=1987
  _globals['_CALL_ATTEMPT_COMPLETE']._serialized_end=2229
  _globals['_SESSIONLOGRECORD']._serialized_start=2232
  _globals['_SESSIONLOGRECORD']._serialized_end=3030
  _globals['_SESSIONLOGRECORD_PRINT']._serialized_start=2797
  _globals['_SESSIONLOGRECORD_PRINT']._serialized_end=2824
  _globals['_SESSIONLOGRECORD_STOPREQUEST']._serialized_start=2826
  _globals['_SESSIONLOGRECORD_STOPREQUEST']._serialized_end=2863
  _globals['_SESSIONLOGRECORD_TYPE']._serialized_start=2866
  _globals['_SESSIONLOGRECORD_TYPE']._serialized_end=3030
  _globals['_SESSIONLOG']._serialized_start=3032
  _globals['_SESSIONLOG']._serialized_end=3126
  _globals['_SESSION']._serialized_start=3129
  _globals['_SESSION']._serialized_end=3884
  _globals['_SESSION_INPUTSENTRY']._serialized_start=3741
  _globals['_SESSION_INPUTSENTRY']._serialized_end=3827
  _globals['_SESSION_MEMOENTRY']._serialized_start=3829
  _globals['_SESSION_MEMOENTRY']._serialized_end=3884
# @@protoc_insertion_point(module_scope)
