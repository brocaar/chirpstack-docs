---
title: Quickstart Docker-Compose
menu:
  main:
    parent: guides
    weight: 1
description: Quickstart guide on how to get started with the LoRa Server project using Docker Compose.
---

# Quickstart using Docker-Compose

[Docker Compose](https://docs.docker.com/compose/) (part of Docker) makes
it possible to orchestrate the configuration of multiple Docker containers
at once using a `docker-compose.yml` file.

## Requirements

### Install Docker

Please refer to the [Get Started with Docker](https://www.docker.com/get-started)
guide to install Docker for MacOS or Windows. When installing Docker on Linux,
please refer to one of the following guides:

* [CentOS](https://docs.docker.com/install/linux/docker-ce/centos/#install-docker-ce)
* [Debian](https://docs.docker.com/install/linux/docker-ce/debian/)
* [Fedora](https://docs.docker.com/install/linux/docker-ce/fedora/)
* [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

### Install Compose

To install Docker Compose on Linux, please refer to the
[Install Compose on Linux systems](https://docs.docker.com/compose/install/#install-compose)
guide. You can skip this step for MacOS and Windows.

## LoRa Server

### Configure

The LoRa Server project provides an example `docker-compose.yml` file that
you can use as a starting-point. This example can be found at
[https://github.com/brocaar/loraserver-docker](https://github.com/brocaar/loraserver-docker)
and also contains more documentation.

To clone this repository, you need to execute the following commands:

{{<highlight bash>}}
$ git clone https://github.com/brocaar/loraserver-docker.git
$ cd loraserver-docker
{{< /highlight >}}



### Start

After you have updated the configuration, you can run the following command
to start all Docker containers:

{{<highlight bash>}}
$ docker-compose up
{{< /highlight >}}

Please note that the first time you execute this command, there might be
some errors logged as the database needs to be initialized.

## Add network-server

As each container has its own hostname, you must use the hostname of the 
`loraserver` container when adding the network-server in the LoRa App Server
web-interface.

When using the above example, it means that you must enter `loraserver:8000`
as the network-server hostname:IP. See [network-servers](https://docs.loraserver.io/lora-app-server/use/network-servers/)
for more information.
