// package: guff.proto
// file: test.proto

import * as test_pb from "./test_pb";
import {grpc} from "@improbable-eng/grpc-web";

type TestServiceTestEcho = {
  readonly methodName: string;
  readonly service: typeof TestService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof test_pb.TestEchoRequest;
  readonly responseType: typeof test_pb.TestEchoResponse;
};

export class TestService {
  static readonly serviceName: string;
  static readonly TestEcho: TestServiceTestEcho;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class TestServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  testEcho(
    requestMessage: test_pb.TestEchoRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: test_pb.TestEchoResponse|null) => void
  ): UnaryResponse;
  testEcho(
    requestMessage: test_pb.TestEchoRequest,
    callback: (error: ServiceError|null, responseMessage: test_pb.TestEchoResponse|null) => void
  ): UnaryResponse;
}

