/**
 * protobuf のワイヤ形式の最小限の読み書き（ポータルの API 用）。
 *
 * 型定義は持たず、フィールド番号で値を引く。番号はポータルの JS に埋め込まれた
 * descriptor（isuxportal/resources/*.proto）から読み取ったもの。
 */

export type Fields = Map<number, (bigint | Uint8Array)[]>;

const readVarint = (bytes: Uint8Array, pos: { i: number }): bigint => {
  let result = 0n;
  let shift = 0n;
  for (;;) {
    if (pos.i >= bytes.length) throw new Error("protobuf: varint が途中で切れています");
    const byte = bytes[pos.i++]!;
    result |= BigInt(byte & 0x7f) << shift;
    if ((byte & 0x80) === 0) return result;
    shift += 7n;
  }
};

export const decode = (bytes: Uint8Array): Fields => {
  const fields: Fields = new Map();
  const pos = { i: 0 };
  while (pos.i < bytes.length) {
    const key = Number(readVarint(bytes, pos));
    const field = key >> 3;
    let value: bigint | Uint8Array;
    switch (key & 7) {
      case 0:
        value = readVarint(bytes, pos);
        break;
      case 1:
        value = bytes.subarray(pos.i, pos.i + 8);
        pos.i += 8;
        break;
      case 2: {
        const length = Number(readVarint(bytes, pos));
        value = bytes.subarray(pos.i, pos.i + length);
        pos.i += length;
        break;
      }
      case 5:
        value = bytes.subarray(pos.i, pos.i + 4);
        pos.i += 4;
        break;
      default:
        throw new Error(`protobuf: 読めない wire type ${key & 7}`);
    }
    if (pos.i > bytes.length) throw new Error("protobuf: メッセージが途中で切れています");
    fields.set(field, [...(fields.get(field) ?? []), value]);
  }
  return fields;
};

const last = (fields: Fields, n: number): bigint | Uint8Array | undefined => fields.get(n)?.at(-1);

export const int = (fields: Fields, n: number): number => {
  const value = last(fields, n);
  return typeof value === "bigint" ? Number(BigInt.asIntN(64, value)) : 0;
};

export const bool = (fields: Fields, n: number): boolean => int(fields, n) !== 0;

export const str = (fields: Fields, n: number): string => {
  const value = last(fields, n);
  return value instanceof Uint8Array ? Buffer.from(value).toString("utf8") : "";
};

export const msg = (fields: Fields, n: number): Fields => {
  const value = last(fields, n);
  return value instanceof Uint8Array ? decode(value) : new Map();
};

export const msgs = (fields: Fields, n: number): Fields[] =>
  (fields.get(n) ?? []).flatMap((value) => (value instanceof Uint8Array ? [decode(value)] : []));

const varint = (value: bigint): number[] => {
  const out: number[] = [];
  let rest = BigInt.asUintN(64, value);
  while (rest >= 0x80n) {
    out.push(Number(rest & 0x7fn) | 0x80);
    rest >>= 7n;
  }
  out.push(Number(rest));
  return out;
};

export const encodeInt = (n: number, value: number): Uint8Array =>
  Uint8Array.from([...varint(BigInt(n << 3)), ...varint(BigInt(value))]);

export const encodeStr = (n: number, value: string): Uint8Array => {
  const bytes = Buffer.from(value, "utf8");
  return Uint8Array.from([...varint(BigInt((n << 3) | 2)), ...varint(BigInt(bytes.length)), ...bytes]);
};

export const encodeMsg = (n: number, ...parts: Uint8Array[]): Uint8Array => {
  const body = Buffer.concat(parts);
  return Uint8Array.from([...varint(BigInt((n << 3) | 2)), ...varint(BigInt(body.length)), ...body]);
};

const printable = (bytes: Uint8Array): string | undefined => {
  try {
    const text = new TextDecoder("utf-8", { fatal: true }).decode(bytes);
    return /[\x00-\x08\x0b\x0c\x0e-\x1f]/.test(text) ? undefined : text;
  } catch {
    return undefined;
  }
};

/** 型定義なしで中身を見るためのダンプ（isuh portal get 用）。文字列として読めなければ入れ子として展開する。 */
export const dumpFields = (fields: Fields): Record<string, unknown> => {
  const result: Record<string, unknown> = {};
  for (const [n, values] of fields) {
    const shown = values.map((value) => {
      if (typeof value === "bigint") return value.toString();
      const text = printable(value);
      if (text !== undefined) return text;
      try {
        const nested = decode(value);
        if (nested.size > 0) return dumpFields(nested);
      } catch {
        /* 入れ子のメッセージでなければ文字列として出す */
      }
      return Buffer.from(value).toString("utf8");
    });
    result[n] = shown.length === 1 ? shown[0] : shown;
  }
  return result;
};
