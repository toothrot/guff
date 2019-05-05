/**
 * @fileoverview gRPC-Web generated client stub for guff.proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  TestEchoRequest,
  TestEchoResponse} from './test_pb';

export class TestServiceClient {
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

  methodInfoTestEcho = new grpcWeb.AbstractClientBase.MethodInfo(
    TestEchoResponse,
    (request: TestEchoRequest) => {
      return request.serializeBinary();
    },
    TestEchoResponse.deserializeBinary
  );

  testEcho(
    request: TestEchoRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: TestEchoResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/guff.proto.TestService/TestEcho',
      request,
      metadata || {},
      this.methodInfoTestEcho,
      callback);
  }

}

