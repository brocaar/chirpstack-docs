---
title: Ansible and Vagrant install
menu:
    main:
        parent: install
        weight: 5
---

# Ansible based deployment

[Ansible](https://docs.ansible.com/) is an open-source tool for automating
deployment and server-management related steps.

The LoRa Server project provides a so-called Ansible-playbook which can be
used to deploy LoRa Server, including the setup of all requirements.

The source of this playbook, including additional documentation how to
use this, can be found at [https://github.com/brocaar/loraserver-setup](https://github.com/brocaar/loraserver-setup).

## Local VM deployment using Vagrant

[Vagrant](https://www.vagrantup.com/) is a tool for automating the creation
of virtual machines. It can integrate with Ansible so that it not only create
the VM for you, but also will install all LoRa Server components.

After following the instructions mentioned in the [loraserver-setup](https://github.com/brocaar/loraserver-setup)
repository, this allows you to create a local test environment running
inside a VM:

```bash
vagrant up
```

As this is using exactly the same Ansible-playbook as for remote deployments,
this can also be used for testing before doing a remote deployment, e.g.
when making modifications to the playbook.

## Remote deployment

Ansible can also be used to do remote deployments. You need to setup a
so-called inventory of servers, to which Ansible will connect for executing
the deployment steps. After following the steps mentioned in
[loraserver-setup](https://github.com/brocaar/loraserver-setup), the following
would perform a remote deployment:

```bash
ansible-playbook -i inventory full_deploy.yml
```
