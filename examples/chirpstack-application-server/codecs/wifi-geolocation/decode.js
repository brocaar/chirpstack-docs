// bytes contains one or multiple times thw following encoding:
// [MAC Address: 6 bytes] | [Signal strength (int8): 1 byte]
function Decode(fPort, bytes, variables) {
  var i = 0;
  var len = bytes.length;
  var out = {access_points: []};
  
  for (; i < len ;) {
    out.access_points.push({macAddress: bytes.slice(i, i + 6), signalStrength: int8(bytes[i+6])});
    i += 7;
  }
  return out;
}

// convert a byte value to signed int8
function int8(byte) {
  var sign = byte & (1 << 7);
  if (sign) {
    return 0xFFFFFF00 | byte;
  }
  return byte;
}
