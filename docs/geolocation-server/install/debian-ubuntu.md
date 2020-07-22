---
description: Instructions on how to install ChirpStack Geolocation Server on a Debian or Ubuntu based Linux installation.
---

# Debian / Ubuntu installation

These steps have been tested on:

* Ubuntu 18.04 LTS
* Debian 10 (Buster)

## ChirpStack Debian repository

ChirpStack provides pre-compiled binaries packaged as Debian (.deb)
packages. In order to activate this repository, execute the following
commands:

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

sudo echo "deb https://artifacts.chirpstack.io/packages/3.x/deb stable main" | sudo tee /etc/apt/sources.list.d/chirpstack.list
sudo apt-get update
```

## Install ChirpStack Geolocation Server

In order to install ChirpStack Geolocation Server, execute the following command:

```bash
sudo apt-get install chirpstack-geolocation-server
```

After installation, modify the configuration file which is located at
`/etc/chirpstack-geolocation-server/chirpstack-geolocation-server.toml`.

Settings you probably want to set / change:

* `geo_server.backend.collos.subscription_key`

## Starting ChirpStack Geolocation Server

How you need to (re)start and stop ChirpStack Geolocation Server depends on if your
distribution uses init.d or systemd.

### init.d

```bash
sudo /etc/init.d/chirpstack-geolocation-server [start|stop|restart|status]
```

### systemd

```bash
sudo systemctl [start|stop|restart|status] chirpstack-geolocation-server
```

## ChirpStack Geolocation Server log output

Now you've setup ChirpStack Geolocation Server, it is a good time to verify that it
is actually up-and-running. This can be done by looking at the ChirpStack Geolocation Server
log output.

Like the previous step, which command you need to use for viewing the
log output depends on if your distribution uses init.d or systemd.

### init.d

All logs are written to `/var/log/chirpstack-geolocation-server/chirpstack-geolocation-server.log`.
To view and follow this logfile:

```bash
tail -f /var/log/chirpstack-geolocation-server/chirpstack-geolocation-server.log
```

### systemd

```bash
journalctl -u chirpstack-geolocation-server -f -n 50
```
