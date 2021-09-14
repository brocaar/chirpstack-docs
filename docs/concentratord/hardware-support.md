# Hardware support

ChirpStack Concentratord currently supports SX1301/8, SX1302 / SX1303 based gateways
and the SX1280 based 2.4 GHz reference-design gateway.
In order to make configuration as easy as possible, it comes with the calibration
values for many different gateway models embedded (based on calibration values
provided by the vendors). Model specific features can be turned on and off using flags.

## SX1301/8

The `chirpstack-concentratord-sx1301` binary implements the [SX1301 HAL](https://github.com/lora-net/lora_gateway).

| Vendor | Gateway / Shield | Model flags | Model |
| --- | --- | --- | --- |
| IMST | iC880A EU868 | | imst_ic880a_eu868 |
| IMST | iC880A IN865 | | imst_ic880a_in865 |
| IMST | iC880A RU864 | | imst_ic880a_ru864 |
| Kerlink | iFemtoCell | | kerlink_ifemtocell_eu868 |
| Multitech | Multitech Conduit AP EU868 | | multitech_mtcap_lora_868_eu868 |
| Multitech | Multitech Conduit AP US915 | | multitech_mtcap_lora_915_us915 |
| Multitech | Multitech Conduit MTAC-LORA-H-868 | AP1, AP2, GNSS | multitech_mtac_lora_h_868_eu868 |
| Multitech | Multitech Conduit MTAC-LORA-H-915 | AP1, AP2, GNSS | multitech_mtac_lora_h_915_us915 |
| Pi Supply | LoRa Gateway HAT AU915 | GNSS | pi_supply_lora_gateway_hat_au916 |
| Pi Supply | LoRa Gateway HAT EU868 | GNSS | pi_supply_lora_gateway_hat_eu868 |
| Pi Supply | LoRa Gateway HAT US915 | GNSS | pi_supply_lora_gateway_hat_us915 |
| RAK | RAK2245 AS923 | GNSS | rak_2245_as923 |
| RAK | RAK2245 AU915 | GNSS | rak_2245_au915 |
| RAK | RAK2245 CN470 | GNSS | rak_2245_cn470 |
| RAK | RAK2245 EU433 | GNSS | rak_2245_eu433 |
| RAK | RAK2245 EU868 | GNSS | rak_2245_eu868 |
| RAK | RAK2245 IN865 | GNSS | rak_2245_in865 |
| RAK | RAK2245 KR920 | GNSS | rak_2245_kr920 |
| RAK | RAK2245 RU864 | GNSS | rak_2245_ru864 |
| RAK | RAK2245 US915 | GNSS | rak_2245_us915 |
| RAK | RAK2246 AS923 | GNSS | rak_2246_as923 |
| RAK | RAK2246 AU915 | GNSS | rak_2246_au915 |
| RAK | RAK2246 EU868 | GNSS | rak_2246_eu868 |
| RAK | RAK2246 IN865 | GNSS | rak_2246_in865 |
| RAK | RAK2246 KR920 | GNSS | rak_2246_kr920 |
| RAK | RAK2246 RU864 | GNSS | rak_2246_ru864 |
| RAK | RAK2246 US915 | GNSS | rak_2246_us915 |
| RisingHF | RHF0M301 EU868 | | risinghf_rhf0m301_eu868 |
| RisingHF | RHF0M301 US915 | | risinghf_rhf0m301_us915 |
| Sandbox | LoRaGo Port EU868 | | sandbox_lorago_port_eu868 |
| Sandbox | LoRaGo Port US915 | | sandbox_lorago_port_us915 |
| Wifx | LORIX One EU868 | | wifx_lorix_one_eu868 |

## SX1302

The `chirpstack-concentratord-sx1302` binary implements the [SX1302 HAL](https://github.com/lora-net/sx1302_hal).

| Vendor | Gateway / Shield | Model flags | Model |
| --- | --- | --- | --- |
| RAK | RAK2287 AS923 | GNSS, USB | rak_2287_as923 |
| RAK | RAK2287 AU915 | GNSS, USB | rak_2287_au915 |
| RAK | RAK2287 EU868 | GNSS, USB | rak_2287_eu868 |
| RAK | RAK2287 IN865 | GNSS, USB | rak_2287_in865 |
| RAK | RAK2287 KR920 | GNSS, USB | rak_2287_kr920 |
| RAK | RAK2287 RU864 | GNSS, USB | rak_2287_ru864 |
| RAK | RAK2287 US915 | GNSS, USB | rak_2287_us915 |
| Semtech | CoreCell EU868 | | semtech_sx1302c868gw1_eu868 |
| Semtech | CoreCell US915 | | semtech_sx1302c915gw1_us915 |

## 2G4 (SX1280)

The `chirpstack-concentratord-2g4` binary implements the [2g4 HAL](https://github.com/Lora-net/gateway_2g4_hal/).

| Vendor | Gateway / Shield | Model flags | Model |
| --- | --- | --- | --- |
| Semtech | SX1280ZXXXXGW1 | | semtech_sx1280z3dsfgw1 |
