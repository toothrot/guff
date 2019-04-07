// package: guff.proto
// file: divisions.proto

import * as divisions_pb from "./divisions_pb";
import {grpc} from "@improbable-eng/grpc-web";

type DivisionsServiceGetDivisions = {
  readonly methodName: string;
  readonly service: typeof DivisionsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof divisions_pb.GetDivisionsRequest;
  readonly responseType: typeof divisions_pb.GetDivisionsResponse;
};

export class DivisionsService {
  static readonly serviceName: string;
  static readonly GetDivisions: DivisionsServiceGetDivisions;
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

export class DivisionsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getDivisions(
    requestMessage: divisions_pb.GetDivisionsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: divisions_pb.GetDivisionsResponse|null) => void
  ): UnaryResponse;
  getDivisions(
    requestMessage: divisions_pb.GetDivisionsRequest,
    callback: (error: ServiceError|null, responseMessage: divisions_pb.GetDivisionsResponse|null) => void
  ): UnaryResponse;
}

