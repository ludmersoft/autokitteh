# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/sessions/v1/svc.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from autokitteh_pb.sessions.v1 import session_pb2 as autokitteh_dot_sessions_dot_v1_dot_session__pb2
from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n autokitteh/sessions/v1/svc.proto\x12\x16\x61utokitteh.sessions.v1\x1a$autokitteh/sessions/v1/session.proto\x1a\x1b\x62uf/validate/validate.proto\"\xe8\x02\n\x0cStartRequest\x12\x42\n\x07session\x18\x01 \x01(\x0b\x32\x1f.autokitteh.sessions.v1.SessionB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x07session\x12U\n\x0bjson_inputs\x18\x02 \x03(\x0b\x32\x34.autokitteh.sessions.v1.StartRequest.JsonInputsEntryR\njsonInputs\x1a=\n\x0fJsonInputsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01:~\xfa\xf7\x18z\x1ax\n session.session_id_must_be_empty\x12 session_id must not be specified\x1a\x32has(this.session) && this.session.session_id == \'\'\"8\n\rStartResponse\x12\'\n\nsession_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\tsessionId\"l\n\x0bStopRequest\x12\'\n\nsession_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\tsessionId\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\x12\x1c\n\tterminate\x18\x03 \x01(\x08R\tterminate\"\x0e\n\x0cStopResponse\"\x96\x03\n\x0bListRequest\x12#\n\rdeployment_id\x18\x01 \x01(\tR\x0c\x64\x65ploymentId\x12\x1d\n\nproject_id\x18\x02 \x01(\tR\tprojectId\x12\x19\n\x08\x65vent_id\x18\x03 \x01(\tR\x07\x65ventId\x12\x19\n\x08\x62uild_id\x18\x04 \x01(\tR\x07\x62uildId\x12R\n\nstate_type\x18\x05 \x01(\x0e\x32(.autokitteh.sessions.v1.SessionStateTypeB\t\xfa\xf7\x18\x05\x82\x01\x02\x10\x01R\tstateType\x12\x15\n\x06org_id\x18\x06 \x01(\tR\x05orgId\x12\x1d\n\ncount_only\x18\n \x01(\x08R\tcountOnly\x12\x1b\n\tpage_size\x18\x14 \x01(\x05R\x08pageSize\x12G\n\x04skip\x18\x15 \x01(\x05\x42\x33\xfa\xf7\x18/\xba\x01,\n\x11session.list.skip\x12\x0cMust be >= 0\x1a\tthis >= 0R\x04skip\x12\x1d\n\npage_token\x18\x16 \x01(\tR\tpageToken\"\x97\x01\n\x0cListResponse\x12I\n\x08sessions\x18\x01 \x03(\x0b\x32\x1f.autokitteh.sessions.v1.SessionB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x08sessions\x12\x14\n\x05\x63ount\x18\x02 \x01(\x03R\x05\x63ount\x12&\n\x0fnext_page_token\x18\n \x01(\tR\rnextPageToken\"V\n\nGetRequest\x12\'\n\nsession_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\tsessionId\x12\x1f\n\x0bjson_values\x18\x02 \x01(\x08R\njsonValues\"Q\n\x0bGetResponse\x12\x42\n\x07session\x18\x01 \x01(\x0b\x32\x1f.autokitteh.sessions.v1.SessionB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x07session\"\xc1\x02\n\rGetLogRequest\x12\'\n\nsession_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\tsessionId\x12\x1f\n\x0bjson_values\x18\x02 \x01(\x08R\njsonValues\x12\x43\n\x05types\x18\x03 \x01(\x0e\x32-.autokitteh.sessions.v1.SessionLogRecord.TypeR\x05types\x12\x1c\n\tascending\x18\x0b \x01(\x08R\tascending\x12\x1b\n\tpage_size\x18\x14 \x01(\x05R\x08pageSize\x12G\n\x04skip\x18\x15 \x01(\x05\x42\x33\xfa\xf7\x18/\xba\x01,\n\x11session.list.skip\x12\x0cMust be >= 0\x1a\tthis >= 0R\x04skip\x12\x1d\n\npage_token\x18\x16 \x01(\tR\tpageToken\"\x8d\x01\n\x0eGetLogResponse\x12=\n\x03log\x18\x01 \x01(\x0b\x32\".autokitteh.sessions.v1.SessionLogB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x03log\x12\x14\n\x05\x63ount\x18\x02 \x01(\x03R\x05\x63ount\x12&\n\x0fnext_page_token\x18\n \x01(\tR\rnextPageToken\"8\n\rDeleteRequest\x12\'\n\nsession_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\tsessionId\"\x10\n\x0e\x44\x65leteResponse2\x8f\x04\n\x0fSessionsService\x12T\n\x05Start\x12$.autokitteh.sessions.v1.StartRequest\x1a%.autokitteh.sessions.v1.StartResponse\x12Q\n\x04Stop\x12#.autokitteh.sessions.v1.StopRequest\x1a$.autokitteh.sessions.v1.StopResponse\x12Q\n\x04List\x12#.autokitteh.sessions.v1.ListRequest\x1a$.autokitteh.sessions.v1.ListResponse\x12N\n\x03Get\x12\".autokitteh.sessions.v1.GetRequest\x1a#.autokitteh.sessions.v1.GetResponse\x12W\n\x06GetLog\x12%.autokitteh.sessions.v1.GetLogRequest\x1a&.autokitteh.sessions.v1.GetLogResponse\x12W\n\x06\x44\x65lete\x12%.autokitteh.sessions.v1.DeleteRequest\x1a&.autokitteh.sessions.v1.DeleteResponseB\xed\x01\n\x1a\x63om.autokitteh.sessions.v1B\x08SvcProtoP\x01ZKgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/sessions/v1;sessionsv1\xa2\x02\x03\x41SX\xaa\x02\x16\x41utokitteh.Sessions.V1\xca\x02\x16\x41utokitteh\\Sessions\\V1\xe2\x02\"Autokitteh\\Sessions\\V1\\GPBMetadata\xea\x02\x18\x41utokitteh::Sessions::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.sessions.v1.svc_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\032com.autokitteh.sessions.v1B\010SvcProtoP\001ZKgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/sessions/v1;sessionsv1\242\002\003ASX\252\002\026Autokitteh.Sessions.V1\312\002\026Autokitteh\\Sessions\\V1\342\002\"Autokitteh\\Sessions\\V1\\GPBMetadata\352\002\030Autokitteh::Sessions::V1'
  _STARTREQUEST_JSONINPUTSENTRY._options = None
  _STARTREQUEST_JSONINPUTSENTRY._serialized_options = b'8\001'
  _STARTREQUEST.fields_by_name['session']._options = None
  _STARTREQUEST.fields_by_name['session']._serialized_options = b'\372\367\030\003\310\001\001'
  _STARTREQUEST._options = None
  _STARTREQUEST._serialized_options = b'\372\367\030z\032x\n session.session_id_must_be_empty\022 session_id must not be specified\0322has(this.session) && this.session.session_id == \'\''
  _STARTRESPONSE.fields_by_name['session_id']._options = None
  _STARTRESPONSE.fields_by_name['session_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _STOPREQUEST.fields_by_name['session_id']._options = None
  _STOPREQUEST.fields_by_name['session_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _LISTREQUEST.fields_by_name['state_type']._options = None
  _LISTREQUEST.fields_by_name['state_type']._serialized_options = b'\372\367\030\005\202\001\002\020\001'
  _LISTREQUEST.fields_by_name['skip']._options = None
  _LISTREQUEST.fields_by_name['skip']._serialized_options = b'\372\367\030/\272\001,\n\021session.list.skip\022\014Must be >= 0\032\tthis >= 0'
  _LISTRESPONSE.fields_by_name['sessions']._options = None
  _LISTRESPONSE.fields_by_name['sessions']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _GETREQUEST.fields_by_name['session_id']._options = None
  _GETREQUEST.fields_by_name['session_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _GETRESPONSE.fields_by_name['session']._options = None
  _GETRESPONSE.fields_by_name['session']._serialized_options = b'\372\367\030\003\310\001\001'
  _GETLOGREQUEST.fields_by_name['session_id']._options = None
  _GETLOGREQUEST.fields_by_name['session_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _GETLOGREQUEST.fields_by_name['skip']._options = None
  _GETLOGREQUEST.fields_by_name['skip']._serialized_options = b'\372\367\030/\272\001,\n\021session.list.skip\022\014Must be >= 0\032\tthis >= 0'
  _GETLOGRESPONSE.fields_by_name['log']._options = None
  _GETLOGRESPONSE.fields_by_name['log']._serialized_options = b'\372\367\030\003\310\001\001'
  _DELETEREQUEST.fields_by_name['session_id']._options = None
  _DELETEREQUEST.fields_by_name['session_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _globals['_STARTREQUEST']._serialized_start=128
  _globals['_STARTREQUEST']._serialized_end=488
  _globals['_STARTREQUEST_JSONINPUTSENTRY']._serialized_start=299
  _globals['_STARTREQUEST_JSONINPUTSENTRY']._serialized_end=360
  _globals['_STARTRESPONSE']._serialized_start=490
  _globals['_STARTRESPONSE']._serialized_end=546
  _globals['_STOPREQUEST']._serialized_start=548
  _globals['_STOPREQUEST']._serialized_end=656
  _globals['_STOPRESPONSE']._serialized_start=658
  _globals['_STOPRESPONSE']._serialized_end=672
  _globals['_LISTREQUEST']._serialized_start=675
  _globals['_LISTREQUEST']._serialized_end=1081
  _globals['_LISTRESPONSE']._serialized_start=1084
  _globals['_LISTRESPONSE']._serialized_end=1235
  _globals['_GETREQUEST']._serialized_start=1237
  _globals['_GETREQUEST']._serialized_end=1323
  _globals['_GETRESPONSE']._serialized_start=1325
  _globals['_GETRESPONSE']._serialized_end=1406
  _globals['_GETLOGREQUEST']._serialized_start=1409
  _globals['_GETLOGREQUEST']._serialized_end=1730
  _globals['_GETLOGRESPONSE']._serialized_start=1733
  _globals['_GETLOGRESPONSE']._serialized_end=1874
  _globals['_DELETEREQUEST']._serialized_start=1876
  _globals['_DELETEREQUEST']._serialized_end=1932
  _globals['_DELETERESPONSE']._serialized_start=1934
  _globals['_DELETERESPONSE']._serialized_end=1950
  _globals['_SESSIONSSERVICE']._serialized_start=1953
  _globals['_SESSIONSSERVICE']._serialized_end=2480
# @@protoc_insertion_point(module_scope)
