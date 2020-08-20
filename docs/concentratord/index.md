# Introduction

ChirpStack Concentratord is an open-source LoRa(WAN) concentrator daemon, part
of the [ChirpStack](../index.md) project. It exposes a [ZeroMQ](https://zeromq.org/)
based API that can be used by one or multiple applications to interact with
gateway hardware. By implementing and abstracting the the hardware specifics
in a separate daemon and exposing this over a ZeroMQ based API, the packet
forwarding application can be completely decoupled from the gateway hardware.
It also allows for running multiple packet forwarding applications simultaniously.
