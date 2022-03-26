/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.blog.blog";

export interface Art {
  creator: string;
  id: number;
  title: string;
  body: string;
  price: string;
  oldArt: string;
  authorArt: string;
  authorSender: string;
  published: string;
  image: string;
}

const baseArt: object = {
  creator: "",
  id: 0,
  title: "",
  body: "",
  price: "",
  oldArt: "",
  authorArt: "",
  authorSender: "",
  published: "",
  image: "",
};

export const Art = {
  encode(message: Art, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.body !== "") {
      writer.uint32(34).string(message.body);
    }
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    if (message.oldArt !== "") {
      writer.uint32(50).string(message.oldArt);
    }
    if (message.authorArt !== "") {
      writer.uint32(58).string(message.authorArt);
    }
    if (message.authorSender !== "") {
      writer.uint32(66).string(message.authorSender);
    }
    if (message.published !== "") {
      writer.uint32(74).string(message.published);
    }
    if (message.image !== "") {
      writer.uint32(82).string(message.image);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Art {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseArt } as Art;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.body = reader.string();
          break;
        case 5:
          message.price = reader.string();
          break;
        case 6:
          message.oldArt = reader.string();
          break;
        case 7:
          message.authorArt = reader.string();
          break;
        case 8:
          message.authorSender = reader.string();
          break;
        case 9:
          message.published = reader.string();
          break;
        case 10:
          message.image = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Art {
    const message = { ...baseArt } as Art;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = String(object.body);
    } else {
      message.body = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.oldArt !== undefined && object.oldArt !== null) {
      message.oldArt = String(object.oldArt);
    } else {
      message.oldArt = "";
    }
    if (object.authorArt !== undefined && object.authorArt !== null) {
      message.authorArt = String(object.authorArt);
    } else {
      message.authorArt = "";
    }
    if (object.authorSender !== undefined && object.authorSender !== null) {
      message.authorSender = String(object.authorSender);
    } else {
      message.authorSender = "";
    }
    if (object.published !== undefined && object.published !== null) {
      message.published = String(object.published);
    } else {
      message.published = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = String(object.image);
    } else {
      message.image = "";
    }
    return message;
  },

  toJSON(message: Art): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.body !== undefined && (obj.body = message.body);
    message.price !== undefined && (obj.price = message.price);
    message.oldArt !== undefined && (obj.oldArt = message.oldArt);
    message.authorArt !== undefined && (obj.authorArt = message.authorArt);
    message.authorSender !== undefined &&
      (obj.authorSender = message.authorSender);
    message.published !== undefined && (obj.published = message.published);
    message.image !== undefined && (obj.image = message.image);
    return obj;
  },

  fromPartial(object: DeepPartial<Art>): Art {
    const message = { ...baseArt } as Art;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = object.body;
    } else {
      message.body = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.oldArt !== undefined && object.oldArt !== null) {
      message.oldArt = object.oldArt;
    } else {
      message.oldArt = "";
    }
    if (object.authorArt !== undefined && object.authorArt !== null) {
      message.authorArt = object.authorArt;
    } else {
      message.authorArt = "";
    }
    if (object.authorSender !== undefined && object.authorSender !== null) {
      message.authorSender = object.authorSender;
    } else {
      message.authorSender = "";
    }
    if (object.published !== undefined && object.published !== null) {
      message.published = object.published;
    } else {
      message.published = "";
    }
    if (object.image !== undefined && object.image !== null) {
      message.image = object.image;
    } else {
      message.image = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
