---
title: Ansible and Vagrant
menu:
  main:
    parent: guides
    weight: 2
description: Setting up the ChirpStack stack using Ansible (and optionally using Vagrant).
---

# Ansible based deployments

[Ansible](https://docs.ansible.com/) is an open-source tool for automating
deployment and server-management related steps.
ChirpStack provides a so-called Ansible-playbook which can be
used to deploy the ChirpStack stack, including all requirements.
The source of this playbook, including additional documentation how to
use this, can be found at [https://github.com/brocaar/chirpstack-ansible-playbook](https://github.com/brocaar/chirpstack-ansible-playbook).

## Local VM deployment using Vagrant

[Vagrant](https://www.vagrantup.com/) is a tool for automating the creation
of virtual machines. It can integrate with Ansible so that it not only create
the VM for you, but also will install all ChirpStack stack components.

After following the instructions mentioned in the [chirpstack-ansible-playbook](https://github.com/brocaar/chirpstack-ansible-playbook)
repository, this allows you to create a local test environment running
inside a VM:

{{<highlight bash>}}
vagrant up
{{< /highlight >}}

As this is using exactly the same Ansible-playbook as for remote deployments,
this can also be used for testing before doing a remote deployment, e.g.
when making modifications to the playbook.

## Remote deployment

Ansible can also be used to perform remote deployments. You need to setup a
so-called inventory of servers, to which Ansible will connect for executing
the deployment steps. After following the steps mentioned in
[chirpstack-ansible-playbook](https://github.com/brocaar/chirpstack-ansible-playbook), the following
would perform a remote deployment:

{{<highlight bash>}}
ansible-playbook -i inventory full_deploy.yml
{{< /highlight >}}
