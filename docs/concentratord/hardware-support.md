# Hardware support

ChirpStack Concentratord currently supports SX1301/8 and SX1302 based gateways.
In order to make configuration as easy as possible, it comes with the calibration
values for many different gateway models embedded (based on calibration values
provided by the vendors). Model specific features can be turned on and off using flags.

## SX1301/8

The `chirpstack-concentratord-sx1301` binary implements the [SX1301 HAL](https://github.com/lora-net/lora_gateway).

| Vendor | Gateway / Shield | Model flags | Model |
| --- | --- | --- | --- |
| IMST | iC880A | | imst_ic880a_eu868 |
| Kerlink | iFemtoCell | | kerlink_ifemtocell_eu868 |
| Multitech | Multitech Conduit AP EU868 | | multitech_mtcap_lora_868_eu868 |
| Multitech | Multitech Conduit AP US915 | | multitech_mtcap_lora_915_us915 |
| Multitech | Multitech Conduit MTAC-LORA-H-868 | AP1, AP2, GNSS | multitech_mtac_lora_h_868_eu868 |
| Multitech | Multitech Conduit MTAC-LORA-H-915 | AP1, AP2, GNSS | multitech_mtac_lora_h_915_us915 |
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
| Sandbox | LoRaGo Port EU868 | | sandbox_lorago_port_eu868 |
| Sandbox | LoRaGo Port US915 | | sandbox_lorago_port_us915 |
| Wifx | LORIX One EU868 | | wifx_lorix_one_eu868 |

## SX1302

The `chirpstack-concentratord-sx1302` binary implements the [SX1302 HAL](https://github.com/lora-net/sx1302_hal).

| Vendor | Gateway / Shield | Model flags | Model |
| --- | --- | --- | --- |
| Semtech | CoreCell EU868 | | semtech_sx1302c868gw1_eu868 |
| Semtech | CoreCell US915 | | semtech_sx1302c915gw1_us915 |
