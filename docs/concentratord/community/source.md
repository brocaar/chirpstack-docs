---
description: How to obtain the ChirpStack Concentratord source-code and compile it.
---

# Source

The repository containing the source-code for ChirpStack Concentratord is
located at [https://github.com/brocaar/chirpstack-concentratord](https://github.com/brocaar/chirpstack-concentratord).

## Building from source

The easiest way to (cross)compile ChirpStack Concentratord for various targets
is using the provided [Docker Compose](https://docs.docker.com/compose/) environment.
Before executing one of the following commands, make sure that Docker Compose
is installed.

```bash
# Compile ARMv5 binary
make build-armv5-release

# Compile ARMv7hf binary
make build-armv7hf-release

# Create .ipk for Kerlink iFemtoCell
make package-kerlink-ifemtocell

# Create .ipk for Multitech Conduit
make package-multitech-conduit

# Create .ipk for Multitech Conduit AP
make package-multitech-conduit-ap
```

* Binaries are located under `target/{ARCHITECTURE}/release`
* `.ipk` packages are located under `dist/{VENDOR}/{MODEL}`

### Compile optimizations

The provided `...-release` commands are using the default Rust `release`
mode profile. Several options can be set to minimize the size of the final
binary (usually at the cost of features or compile time).
See https://github.com/johnthagen/min-sized-rust for more information.

