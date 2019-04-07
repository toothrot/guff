// package: guff.proto
// file: divisions.proto

var divisions_pb = require("./divisions_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var DivisionsService = (function () {
  function DivisionsService() {}
  DivisionsService.serviceName = "guff.proto.DivisionsService";
  return DivisionsService;
}());

DivisionsService.GetDivisions = {
  methodName: "GetDivisions",
  service: DivisionsService,
  requestStream: false,
  responseStream: false,
  requestType: divisions_pb.GetDivisionsRequest,
  responseType: divisions_pb.GetDivisionsResponse
};

exports.DivisionsService = DivisionsService;

function DivisionsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DivisionsServiceClient.prototype.getDivisions = function getDivisions(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(DivisionsService.GetDivisions, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.DivisionsServiceClient = DivisionsServiceClient;

