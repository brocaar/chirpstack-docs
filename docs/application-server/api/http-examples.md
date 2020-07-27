---
description: Integrate with the ChirpStack Application Server API using the RESTful API.
---

# HTTP examples

!!! info

	Payload fields in the RESTful JSON API are / must be encoded as [base64](https://en.wikipedia.org/wiki/Base64).
	While the data looks different from [HEX](https://en.wikipedia.org/wiki/Hexadecimal) encoded payloads, it is
	just a different way to represent binary data. A website which can help to make the
	conversion is [ASCII to Hex](https://www.asciitohex.com/).

## Enqueue downlink

The following examples enqueues the base64 payload `AQID` (which is `010203` in HEX)
for the device with DevEUI `0101010101010101`. Do not forget to replace **&lt;API TOKEN&gt;**
with an API token retrieved using the web-interface.

```bash
curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Grpc-Metadata-Authorization: Bearer <API TOKEN>' -d '{ \ 
   "deviceQueueItem": { \ 
     "confirmed": false, \ 
     "data": "AQID", \ 
     "fPort": 10 \ 
   } \ 
 }' 'http://localhost:8080/api/devices/0101010101010101/queue'
```
