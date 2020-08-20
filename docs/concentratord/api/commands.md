# Commands

# Commands

Commands can be sent to Concentratord using a [ZeroMQ REQ](http://zguide.zeromq.org/page:all#toc52)
socket. The first data-frame holds the command type (string), the second
data-frame holds the command payload encoded using Protobuf.

## `gateway_id`

Request the Gateway ID (the data-frame is empty). The response is the 8byte
Gateway ID.

## `down`

Request to enqueue the given frame for downlink (`DownlinkFrame` Protobuf
message). A downlink command is responded by a `DownlinkTXAck` message.

## `config`

Request to re-configure the channel-configuration (`GatewayConfiguration`
Protobuf message). The response is empty.

!!! info

	The re-configuration happens in-memory only. This means that when
	Concentratord is restarted, it will revert to the configuration as specified
	in the configuration file. This allows for restarting Concentratord in case of
	a re-configuration error (in which case the process will terminate) and
	reverting back to the original configuration.
