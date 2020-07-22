---
description: Finding the log files of the ChirpStack components.
---

# Log files

The ChirpStack components are writing their log output to [Syslog](https://en.wikipedia.org/wiki/Syslog),
which writes to `/var/log/messages`.

## All logs

To view the logs, run:

```bash
sudo tail -f /var/log/messages
```

## By service

To filter the messages by a specific service (e.g. `chirpstack-concentratord`),
run:

```bash
sudo tail -f /var/log/messages |grep chirpstack-concentratord
```
