/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "mohammadrezatorkaman.wework.recruiter";

export interface MsgSubmitJob {
  creator: string;
  title: string;
  description: string;
  tags: string;
  postDeadline: string;
  jobDeadline: string;
  maxPrice: string;
  location: string;
  jobtype: number;
}

export interface MsgSubmitJobResponse {}

const baseMsgSubmitJob: object = {
  creator: "",
  title: "",
  description: "",
  tags: "",
  postDeadline: "",
  jobDeadline: "",
  maxPrice: "",
  location: "",
  jobtype: 0,
};

export const MsgSubmitJob = {
  encode(message: MsgSubmitJob, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.tags !== "") {
      writer.uint32(34).string(message.tags);
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
    if (message.jobtype !== 0) {
      writer.uint32(72).uint64(message.jobtype);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitJob {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitJob } as MsgSubmitJob;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.tags = reader.string();
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
          message.jobtype = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitJob {
    const message = { ...baseMsgSubmitJob } as MsgSubmitJob;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
      message.tags = String(object.tags);
    } else {
      message.tags = "";
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
    if (object.jobtype !== undefined && object.jobtype !== null) {
      message.jobtype = Number(object.jobtype);
    } else {
      message.jobtype = 0;
    }
    return message;
  },

  toJSON(message: MsgSubmitJob): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.tags !== undefined && (obj.tags = message.tags);
    message.postDeadline !== undefined &&
      (obj.postDeadline = message.postDeadline);
    message.jobDeadline !== undefined &&
      (obj.jobDeadline = message.jobDeadline);
    message.maxPrice !== undefined && (obj.maxPrice = message.maxPrice);
    message.location !== undefined && (obj.location = message.location);
    message.jobtype !== undefined && (obj.jobtype = message.jobtype);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSubmitJob>): MsgSubmitJob {
    const message = { ...baseMsgSubmitJob } as MsgSubmitJob;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
      message.tags = object.tags;
    } else {
      message.tags = "";
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
    if (object.jobtype !== undefined && object.jobtype !== null) {
      message.jobtype = object.jobtype;
    } else {
      message.jobtype = 0;
    }
    return message;
  },
};

const baseMsgSubmitJobResponse: object = {};

export const MsgSubmitJobResponse = {
  encode(_: MsgSubmitJobResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitJobResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitJobResponse } as MsgSubmitJobResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitJobResponse {
    const message = { ...baseMsgSubmitJobResponse } as MsgSubmitJobResponse;
    return message;
  },

  toJSON(_: MsgSubmitJobResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSubmitJobResponse>): MsgSubmitJobResponse {
    const message = { ...baseMsgSubmitJobResponse } as MsgSubmitJobResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitJob(request: MsgSubmitJob): Promise<MsgSubmitJobResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SubmitJob(request: MsgSubmitJob): Promise<MsgSubmitJobResponse> {
    const data = MsgSubmitJob.encode(request).finish();
    const promise = this.rpc.request(
      "mohammadrezatorkaman.wework.recruiter.Msg",
      "SubmitJob",
      data
    );
    return promise.then((data) =>
      MsgSubmitJobResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
