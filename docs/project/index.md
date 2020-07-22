# The ChirpStack project

## LoRaWAN

LoRaWAN<sup>&reg;</sup> is a Low Power, Wide Area (LPWA)
networking protocol designed to wirelessly connect battery operated
‘things’ to the internet in regional, national or global networks, and targets
key Internet of Things (IoT) requirements such as bi-directional communication,
end-to-end security, mobility and localization services.

* [LoRa Alliance](https://www.lora-alliance.org/)
* [LoRaWAN specification](https://lora-alliance.org/lorawan-for-developers)

## About the ChirpStack LoRaWAN Network Server stack

ChirpStack provides open-source components for LoRaWAN
networks. Together they form a ready-to-use solution including an user-friendly
web-interface for device management and APIs for integration. The modular
architecture makes it possible to integrate within existing infrastructures.
All components are licensed under the MIT license and can be used for commercial
purposes. The following components are provided:

* [ChirpStack Gateway Bridge](../gateway-bridge/index.md): _handles the communication with the LoRaWAN gateways_
* [ChirpStack Network Server](../network-server/index.md): _a LoRaWAN Network Server implementation_
* [ChirpStack Application Server](../application-server/index.md): _a LoRaWAN Application Server implementation_
* [ChirpStack Gateway OS](../gateway-os/index.md): _Linux-based OS to run the (full) ChirpStack stack on a Raspberry Pi based LoRa gateway_

For a more technical understanding of the ChirpStack stack components and how
they work together, please refer to the
[Architecture](architecture.md) page.
