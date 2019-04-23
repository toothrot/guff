/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.guff.proto.GetCurrentUserRequest', null, global);
goog.exportSymbol('proto.guff.proto.GetCurrentUserResponse', null, global);
goog.exportSymbol('proto.guff.proto.GoogleOAuthConfig', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.guff.proto.GetCurrentUserRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.guff.proto.GetCurrentUserRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.guff.proto.GetCurrentUserRequest.displayName = 'proto.guff.proto.GetCurrentUserRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.guff.proto.GetCurrentUserRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.guff.proto.GetCurrentUserRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.guff.proto.GetCurrentUserRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GetCurrentUserRequest.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.guff.proto.GetCurrentUserRequest}
 */
proto.guff.proto.GetCurrentUserRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.guff.proto.GetCurrentUserRequest;
  return proto.guff.proto.GetCurrentUserRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.guff.proto.GetCurrentUserRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.guff.proto.GetCurrentUserRequest}
 */
proto.guff.proto.GetCurrentUserRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.guff.proto.GetCurrentUserRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.guff.proto.GetCurrentUserRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.guff.proto.GetCurrentUserRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GetCurrentUserRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.guff.proto.GetCurrentUserResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.guff.proto.GetCurrentUserResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.guff.proto.GetCurrentUserResponse.displayName = 'proto.guff.proto.GetCurrentUserResponse';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.guff.proto.GetCurrentUserResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.guff.proto.GetCurrentUserResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GetCurrentUserResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    email: jspb.Message.getFieldWithDefault(msg, 1, ""),
    googleOauthConfig: (f = msg.getGoogleOauthConfig()) && proto.guff.proto.GoogleOAuthConfig.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.guff.proto.GetCurrentUserResponse}
 */
proto.guff.proto.GetCurrentUserResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.guff.proto.GetCurrentUserResponse;
  return proto.guff.proto.GetCurrentUserResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.guff.proto.GetCurrentUserResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.guff.proto.GetCurrentUserResponse}
 */
proto.guff.proto.GetCurrentUserResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 2:
      var value = new proto.guff.proto.GoogleOAuthConfig;
      reader.readMessage(value,proto.guff.proto.GoogleOAuthConfig.deserializeBinaryFromReader);
      msg.setGoogleOauthConfig(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.guff.proto.GetCurrentUserResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.guff.proto.GetCurrentUserResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GetCurrentUserResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getGoogleOauthConfig();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.guff.proto.GoogleOAuthConfig.serializeBinaryToWriter
    );
  }
};


/**
 * optional string email = 1;
 * @return {string}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.guff.proto.GetCurrentUserResponse.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional GoogleOAuthConfig google_oauth_config = 2;
 * @return {?proto.guff.proto.GoogleOAuthConfig}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.getGoogleOauthConfig = function() {
  return /** @type{?proto.guff.proto.GoogleOAuthConfig} */ (
    jspb.Message.getWrapperField(this, proto.guff.proto.GoogleOAuthConfig, 2));
};


/** @param {?proto.guff.proto.GoogleOAuthConfig|undefined} value */
proto.guff.proto.GetCurrentUserResponse.prototype.setGoogleOauthConfig = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.guff.proto.GetCurrentUserResponse.prototype.clearGoogleOauthConfig = function() {
  this.setGoogleOauthConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.hasGoogleOauthConfig = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.guff.proto.GoogleOAuthConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.guff.proto.GoogleOAuthConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.guff.proto.GoogleOAuthConfig.displayName = 'proto.guff.proto.GoogleOAuthConfig';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.guff.proto.GoogleOAuthConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.guff.proto.GoogleOAuthConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.guff.proto.GoogleOAuthConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GoogleOAuthConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    clientId: jspb.Message.getFieldWithDefault(msg, 1, ''),
    loginurl: jspb.Message.getFieldWithDefault(msg, 2, '')
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.guff.proto.GoogleOAuthConfig}
 */
proto.guff.proto.GoogleOAuthConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.guff.proto.GoogleOAuthConfig;
  return proto.guff.proto.GoogleOAuthConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.guff.proto.GoogleOAuthConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.guff.proto.GoogleOAuthConfig}
 */
proto.guff.proto.GoogleOAuthConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientId(value);
      break
      case 2:
        var value = /** @type {string} */ (reader.readString())
        msg.setLoginurl(value)
        break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.guff.proto.GoogleOAuthConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.guff.proto.GoogleOAuthConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.guff.proto.GoogleOAuthConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GoogleOAuthConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLoginurl()
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string client_id = 1;
 * @return {string}
 */
proto.guff.proto.GoogleOAuthConfig.prototype.getClientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.guff.proto.GoogleOAuthConfig.prototype.setClientId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};

/**
 * optional string loginURL = 2;
 * @return {string}
 */
proto.guff.proto.GoogleOAuthConfig.prototype.getLoginurl = function () {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ''))
}

/** @param {string} value */
proto.guff.proto.GoogleOAuthConfig.prototype.setLoginurl = function (value) {
  jspb.Message.setProto3StringField(this, 2, value)
}


goog.object.extend(exports, proto.guff.proto);
