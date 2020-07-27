---
description: Integrate with ChirpStack Application Server using the API interface.
---

# API

ChirpStack Application Server provides to API interface that can be used for
building integration. Both interfaces provide _exactly_ the same functionality.

## gRPC

[gRPC](https://grpc.io/) is a high-performance, open-source universal RPC
framework. [Protocol Buffers](https://developers.google.com/protocol-buffers)
definitions are used to define this API. All definitions are hosted in the
[chirpstack-api](https://github.com/brocaar/chirpstack-api) repository.

Using the gRPC toolsets, it is possible to generate API clients for various
programming languages. The following languages are officially supported by
gRPC but there are additional community-based implementations:

* C++
* Go
* Node.js
* Java
* Ruby
* Android Java
* PHP
* Python
* C#
* Objective-C

### Authentication

In order to use the gRPC API methods, you must provide per-RPC credentials.
The key for this metadata is `authorization`, the value `Bearer <API TOKEN>`
(replace **&lt;API TOKEN&gt;** with the API TOKEN obtained using the web-interface).

### ChirpStack SDKs

The ChirpStack project provides SDKs for the following programming languages:

* [Go](https://pkg.go.dev/github.com/brocaar/chirpstack-api/go/v3/)
* [Python](https://pypi.org/project/chirpstack-api/)
* [JavaScript](https://www.npmjs.com/package/@chirpstack/chirpstack-api)
* [Rust](https://crates.io/crates/chirpstack_api)

## RESTful JSON interface

The RESTful JSON API interface is provides by an embedded [Restful HTTP API to gRPC](https://github.com/grpc-ecosystem/grpc-gateway)
proxy. While it can be more convenient for simple use-cases, it is slightly
less performant due to the additional translation layer.

ChirpStack Application Server provides an API console containing all API
methods and their documentation. This console can be accessed at `/api`
(e.g. [http://localhost:8080/api](http://localhost:8080/api)).

### Authentication

In order to use the API methods within the API console, you must enter your
API token, which can be obtained using the web-interface, in the JWT TOKEN field
(no need to press _Enter_).

When using the API endpoints within your own integrations or using for example
cURL, you must set the `Grpc-Metadata-Authorization` header to `Bearer <API TOKEN>`.
