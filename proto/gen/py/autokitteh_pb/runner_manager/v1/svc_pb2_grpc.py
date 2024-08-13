# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from autokitteh_pb.runner_manager.v1 import svc_pb2 as autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2


class RunnerManagerServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Start = channel.unary_unary(
                '/autokitteh.runner_manager.v1.RunnerManagerService/Start',
                request_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartRequest.SerializeToString,
                response_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartResponse.FromString,
                )
        self.RunnerHealth = channel.unary_unary(
                '/autokitteh.runner_manager.v1.RunnerManagerService/RunnerHealth',
                request_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthRequest.SerializeToString,
                response_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthResponse.FromString,
                )
        self.Stop = channel.unary_unary(
                '/autokitteh.runner_manager.v1.RunnerManagerService/Stop',
                request_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopRequest.SerializeToString,
                response_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopResponse.FromString,
                )
        self.Health = channel.unary_unary(
                '/autokitteh.runner_manager.v1.RunnerManagerService/Health',
                request_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthRequest.SerializeToString,
                response_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthResponse.FromString,
                )


class RunnerManagerServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Start(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RunnerHealth(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Stop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Health(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RunnerManagerServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Start': grpc.unary_unary_rpc_method_handler(
                    servicer.Start,
                    request_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartRequest.FromString,
                    response_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartResponse.SerializeToString,
            ),
            'RunnerHealth': grpc.unary_unary_rpc_method_handler(
                    servicer.RunnerHealth,
                    request_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthRequest.FromString,
                    response_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthResponse.SerializeToString,
            ),
            'Stop': grpc.unary_unary_rpc_method_handler(
                    servicer.Stop,
                    request_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopRequest.FromString,
                    response_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopResponse.SerializeToString,
            ),
            'Health': grpc.unary_unary_rpc_method_handler(
                    servicer.Health,
                    request_deserializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthRequest.FromString,
                    response_serializer=autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'autokitteh.runner_manager.v1.RunnerManagerService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RunnerManagerService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Start(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.runner_manager.v1.RunnerManagerService/Start',
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartRequest.SerializeToString,
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StartResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RunnerHealth(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.runner_manager.v1.RunnerManagerService/RunnerHealth',
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthRequest.SerializeToString,
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.RunnerHealthResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Stop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.runner_manager.v1.RunnerManagerService/Stop',
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopRequest.SerializeToString,
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.StopResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Health(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.runner_manager.v1.RunnerManagerService/Health',
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthRequest.SerializeToString,
            autokitteh_dot_runner__manager_dot_v1_dot_svc__pb2.HealthResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
