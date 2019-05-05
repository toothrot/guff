/**
 * @fileoverview gRPC-Web generated client stub for guff.proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  GetDivisionsRequest,
  GetDivisionsResponse} from './divisions_pb';

export class DivisionsServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; }) {
    if (!options) options = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetDivisions = new grpcWeb.AbstractClientBase.MethodInfo(
    GetDivisionsResponse,
    (request: GetDivisionsRequest) => {
      return request.serializeBinary();
    },
    GetDivisionsResponse.deserializeBinary
  );

  getDivisions(
    request: GetDivisionsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: GetDivisionsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/guff.proto.DivisionsService/GetDivisions',
      request,
      metadata || {},
      this.methodInfoGetDivisions,
      callback);
  }

}

