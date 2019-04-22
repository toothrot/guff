/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.guff.proto.GetCurrentUserResponse')

goog.require('jspb.BinaryReader')
goog.require('jspb.BinaryWriter')
goog.require('jspb.Message')
goog.require('proto.guff.proto.AppStatus')
goog.require('proto.guff.proto.UserRoles')

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
proto.guff.proto.GetCurrentUserResponse = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null)
}
goog.inherits(proto.guff.proto.GetCurrentUserResponse, jspb.Message)
if (goog.DEBUG && !COMPILED) {
  proto.guff.proto.GetCurrentUserResponse.displayName = 'proto.guff.proto.GetCurrentUserResponse'
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
  proto.guff.proto.GetCurrentUserResponse.prototype.toObject = function (opt_includeInstance) {
    return proto.guff.proto.GetCurrentUserResponse.toObject(opt_includeInstance, this)
  }

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Whether to include the JSPB
   *     instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.guff.proto.GetCurrentUserResponse} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.guff.proto.GetCurrentUserResponse.toObject = function (includeInstance, msg) {
    var f, obj = {
      appStatus: (f = msg.getAppStatus()) && proto.guff.proto.AppStatus.toObject(includeInstance, f),
      userRoles: (f = msg.getUserRoles()) && proto.guff.proto.UserRoles.toObject(includeInstance, f)
    }

    if (includeInstance) {
      obj.$jspbMessageInstance = msg
    }
    return obj
  }
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.guff.proto.GetCurrentUserResponse}
 */
proto.guff.proto.GetCurrentUserResponse.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes)
  var msg = new proto.guff.proto.GetCurrentUserResponse
  return proto.guff.proto.GetCurrentUserResponse.deserializeBinaryFromReader(msg, reader)
}

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.guff.proto.GetCurrentUserResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.guff.proto.GetCurrentUserResponse}
 */
proto.guff.proto.GetCurrentUserResponse.deserializeBinaryFromReader = function (msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break
    }
    var field = reader.getFieldNumber()
    switch (field) {
      case 1:
        var value = new proto.guff.proto.AppStatus
        reader.readMessage(value, proto.guff.proto.AppStatus.deserializeBinaryFromReader)
        msg.setAppStatus(value)
        break
      case 2:
        var value = new proto.guff.proto.UserRoles
        reader.readMessage(value, proto.guff.proto.UserRoles.deserializeBinaryFromReader)
        msg.setUserRoles(value)
        break
      default:
        reader.skipField()
        break
    }
  }
  return msg
}

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter()
  proto.guff.proto.GetCurrentUserResponse.serializeBinaryToWriter(this, writer)
  return writer.getResultBuffer()
}

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.guff.proto.GetCurrentUserResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.guff.proto.GetCurrentUserResponse.serializeBinaryToWriter = function (message, writer) {
  var f = undefined
  f = message.getAppStatus()
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.guff.proto.AppStatus.serializeBinaryToWriter
    )
  }
  f = message.getUserRoles()
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.guff.proto.UserRoles.serializeBinaryToWriter
    )
  }
}

/**
 * optional AppStatus app_status = 1;
 * @return {?proto.guff.proto.AppStatus}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.getAppStatus = function () {
  return /** @type{?proto.guff.proto.AppStatus} */ (
    jspb.Message.getWrapperField(this, proto.guff.proto.AppStatus, 1))
}

/** @param {?proto.guff.proto.AppStatus|undefined} value */
proto.guff.proto.GetCurrentUserResponse.prototype.setAppStatus = function (value) {
  jspb.Message.setWrapperField(this, 1, value)
}

proto.guff.proto.GetCurrentUserResponse.prototype.clearAppStatus = function () {
  this.setAppStatus(undefined)
}

/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.hasAppStatus = function () {
  return jspb.Message.getField(this, 1) != null
}

/**
 * optional UserRoles user_roles = 2;
 * @return {?proto.guff.proto.UserRoles}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.getUserRoles = function () {
  return /** @type{?proto.guff.proto.UserRoles} */ (
    jspb.Message.getWrapperField(this, proto.guff.proto.UserRoles, 2))
}

/** @param {?proto.guff.proto.UserRoles|undefined} value */
proto.guff.proto.GetCurrentUserResponse.prototype.setUserRoles = function (value) {
  jspb.Message.setWrapperField(this, 2, value)
}

proto.guff.proto.GetCurrentUserResponse.prototype.clearUserRoles = function () {
  this.setUserRoles(undefined)
}

/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.guff.proto.GetCurrentUserResponse.prototype.hasUserRoles = function () {
  return jspb.Message.getField(this, 2) != null
}

