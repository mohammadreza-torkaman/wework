/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "mohammadrezatorkaman.wework.recruiter";

export interface Job {
  id: number;
  title: string;
  description: string;
  tags: string[];
  postDeadline: string;
  jobDeadline: string;
  maxPrice: string;
  location: string;
  jobType: number;
  status: string;
  contractor: string;
}

const baseJob: object = {
  id: 0,
  title: "",
  description: "",
  tags: "",
  postDeadline: "",
  jobDeadline: "",
  maxPrice: "",
  location: "",
  jobType: 0,
  status: "",
  contractor: "",
};

export const Job = {
  encode(message: Job, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    for (const v of message.tags) {
      writer.uint32(34).string(v!);
    }
    if (message.postDeadline !== "") {
      writer.uint32(42).string(message.postDeadline);
    }
    if (message.jobDeadline !== "") {
      writer.uint32(50).string(message.jobDeadline);
    }
    if (message.maxPrice !== "") {
      writer.uint32(58).string(message.maxPrice);
    }
    if (message.location !== "") {
      writer.uint32(66).string(message.location);
    }
    if (message.jobType !== 0) {
      writer.uint32(72).uint64(message.jobType);
    }
    if (message.status !== "") {
      writer.uint32(82).string(message.status);
    }
    if (message.contractor !== "") {
      writer.uint32(90).string(message.contractor);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Job {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseJob } as Job;
    message.tags = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.tags.push(reader.string());
          break;
        case 5:
          message.postDeadline = reader.string();
          break;
        case 6:
          message.jobDeadline = reader.string();
          break;
        case 7:
          message.maxPrice = reader.string();
          break;
        case 8:
          message.location = reader.string();
          break;
        case 9:
          message.jobType = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.status = reader.string();
          break;
        case 11:
          message.contractor = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Job {
    const message = { ...baseJob } as Job;
    message.tags = [];
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
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(String(e));
      }
    }
    if (object.postDeadline !== undefined && object.postDeadline !== null) {
      message.postDeadline = String(object.postDeadline);
    } else {
      message.postDeadline = "";
    }
    if (object.jobDeadline !== undefined && object.jobDeadline !== null) {
      message.jobDeadline = String(object.jobDeadline);
    } else {
      message.jobDeadline = "";
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = String(object.maxPrice);
    } else {
      message.maxPrice = "";
    }
    if (object.location !== undefined && object.location !== null) {
      message.location = String(object.location);
    } else {
      message.location = "";
    }
    if (object.jobType !== undefined && object.jobType !== null) {
      message.jobType = Number(object.jobType);
    } else {
      message.jobType = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.contractor !== undefined && object.contractor !== null) {
      message.contractor = String(object.contractor);
    } else {
      message.contractor = "";
    }
    return message;
  },

  toJSON(message: Job): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    if (message.tags) {
      obj.tags = message.tags.map((e) => e);
    } else {
      obj.tags = [];
    }
    message.postDeadline !== undefined &&
      (obj.postDeadline = message.postDeadline);
    message.jobDeadline !== undefined &&
      (obj.jobDeadline = message.jobDeadline);
    message.maxPrice !== undefined && (obj.maxPrice = message.maxPrice);
    message.location !== undefined && (obj.location = message.location);
    message.jobType !== undefined && (obj.jobType = message.jobType);
    message.status !== undefined && (obj.status = message.status);
    message.contractor !== undefined && (obj.contractor = message.contractor);
    return obj;
  },

  fromPartial(object: DeepPartial<Job>): Job {
    const message = { ...baseJob } as Job;
    message.tags = [];
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
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.tags !== undefined && object.tags !== null) {
      for (const e of object.tags) {
        message.tags.push(e);
      }
    }
    if (object.postDeadline !== undefined && object.postDeadline !== null) {
      message.postDeadline = object.postDeadline;
    } else {
      message.postDeadline = "";
    }
    if (object.jobDeadline !== undefined && object.jobDeadline !== null) {
      message.jobDeadline = object.jobDeadline;
    } else {
      message.jobDeadline = "";
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = object.maxPrice;
    } else {
      message.maxPrice = "";
    }
    if (object.location !== undefined && object.location !== null) {
      message.location = object.location;
    } else {
      message.location = "";
    }
    if (object.jobType !== undefined && object.jobType !== null) {
      message.jobType = object.jobType;
    } else {
      message.jobType = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.contractor !== undefined && object.contractor !== null) {
      message.contractor = object.contractor;
    } else {
      message.contractor = "";
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
