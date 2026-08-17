# ChirpStack documentation

**NOTE:** This repository has been archieved, as the ChirpStack v4 documentation
source has moved to https://github.com/chirpstack/chirpstack-docs.

This repository contains the global [chirpstack.io](https://www.chirpstack.io) documentation.

## Requirements

The docs are generated using [MkDocs](https://www.mkdocs.org/), which is a
Python based documentation generator. Make sure you have Python and pip installed
first. To install the MkDocs requirements:

```bash
pip install -r requirements.txt
```

## Test server

The following command starts a server for testing the documentation:

```bash
mkdocs serve -a 0.0.0.0:9090
```
