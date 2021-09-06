---
description: Configuring the ChirpStack Gateway OS using the gateway-config utility.
---

# Configuration

ChirpStack Gateway OS comes with an utility called `gateway-config` for
the configuration of the gateway and the services running on the gateway.

This utility can be accessed using SSH. Use the following command to SSH
into the gateway (in this example, the IP of the gateway is `192.168.0.1`):

```bash
ssh admin@192.168.0.1
```

The default username is admin, the default password is admin.

Then start the `gateway-config` utility as `root` user:

```bash
sudo gateway-config
```
