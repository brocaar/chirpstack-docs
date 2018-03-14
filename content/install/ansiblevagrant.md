---
title: Ansible and Vagrant install
menu:
    main:
        parent: install
        weight: 5
---

# Ansible and Vagrant install

## Vagrant (using VirtualBox)

In case you would like to setup a local test environment to experiment with
the LoRa Server project, an easy option is by using the LoRa Server setup
scripts, which can be used both for a local Vagrant based setup as for remote
deployments: [https://github.com/brocaar/loraserver-setup](https://github.com/brocaar/loraserver-setup).

See also: [https://www.vagrantup.com](https://www.vagrantup.com)

## Ansible based deployments

The LoRa Server setup playbook for Ansible will setup the LoRa Server project
for you, including its dependencies. The same playbook is used to provision
the Vagrant box (see above). See for instructions:
[https://github.com/brocaar/loraserver-setup](https://github.com/brocaar/loraserver-setup).

See also: [http://docs.ansible.com](http://docs.ansible.com)
