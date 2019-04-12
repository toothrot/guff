// package: guff.proto
// file: users.proto

var users_pb = require("./users_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var UsersService = (function () {
  function UsersService() {}
  UsersService.serviceName = "guff.proto.UsersService";
  return UsersService;
}());

UsersService.GetCurrentUser = {
  methodName: "GetCurrentUser",
  service: UsersService,
  requestStream: false,
  responseStream: false,
  requestType: users_pb.GetCurrentUserRequest,
  responseType: users_pb.GetCurrentUserResponse
};

exports.UsersService = UsersService;

function UsersServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UsersServiceClient.prototype.getCurrentUser = function getCurrentUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(UsersService.GetCurrentUser, {
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

exports.UsersServiceClient = UsersServiceClient;

