# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/users/v1/user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1e\x61utokitteh/users/v1/user.proto\x12\x13\x61utokitteh.users.v1\x1a\x1b\x62uf/validate/validate.proto\"}\n\x10UserAuthProvider\x12\x1c\n\x04name\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x04name\x12!\n\x07user_id\x18\x02 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x06userId\x12\x14\n\x05\x65mail\x18\x03 \x01(\tR\x05\x65mail\x12\x12\n\x04\x64\x61ta\x18\x04 \x01(\x0cR\x04\x64\x61ta\"\xa0\x01\n\x04User\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12#\n\rprimary_email\x18\x02 \x01(\tR\x0cprimaryEmail\x12Z\n\x0e\x61uth_providers\x18\x03 \x03(\x0b\x32%.autokitteh.users.v1.UserAuthProviderB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\rauthProvidersB\xd9\x01\n\x17\x63om.autokitteh.users.v1B\tUserProtoP\x01ZEgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/users/v1;usersv1\xa2\x02\x03\x41UX\xaa\x02\x13\x41utokitteh.Users.V1\xca\x02\x13\x41utokitteh\\Users\\V1\xe2\x02\x1f\x41utokitteh\\Users\\V1\\GPBMetadata\xea\x02\x15\x41utokitteh::Users::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.users.v1.user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\027com.autokitteh.users.v1B\tUserProtoP\001ZEgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/users/v1;usersv1\242\002\003AUX\252\002\023Autokitteh.Users.V1\312\002\023Autokitteh\\Users\\V1\342\002\037Autokitteh\\Users\\V1\\GPBMetadata\352\002\025Autokitteh::Users::V1'
  _USERAUTHPROVIDER.fields_by_name['name']._options = None
  _USERAUTHPROVIDER.fields_by_name['name']._serialized_options = b'\372\367\030\004r\002\020\001'
  _USERAUTHPROVIDER.fields_by_name['user_id']._options = None
  _USERAUTHPROVIDER.fields_by_name['user_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _USER.fields_by_name['auth_providers']._options = None
  _USER.fields_by_name['auth_providers']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _globals['_USERAUTHPROVIDER']._serialized_start=84
  _globals['_USERAUTHPROVIDER']._serialized_end=209
  _globals['_USER']._serialized_start=212
  _globals['_USER']._serialized_end=372
# @@protoc_insertion_point(module_scope)
