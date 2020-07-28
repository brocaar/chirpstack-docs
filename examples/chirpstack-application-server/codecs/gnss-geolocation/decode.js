// bytes has the following encoding:
// [Sat. count (uint8): 1 byte] | [LR1110 GNSS payload: remaining bytes ]
function Decode(fPort, bytes, variables) {
  return {
    "sat_count": bytes[0],
    "lr1110_gnss": bytes.slice(2, bytes.length)
  };
}
