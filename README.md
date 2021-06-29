<div align="center">

<img src="https://camo.githubusercontent.com/5b298bf6b0596795602bd771c5bddbb963e83e0f/68747470733a2f2f692e696d6775722e636f6d2f7031527a586a512e706e67" align="center" width="144px" height="144px"/>

### My home Kubernetes cluster :sailboat:
... managed with Flux :robot:

</div>

## :book:&nbsp; Overview

This repository _is_ my home Kubernetes cluster in a declarative state. [Flux](https://github.com/fluxcd/flux2) watches my [cluster](./cluster/) folder and makes the changes to my cluster based on the YAML manifests.

Feel free to open a [Github issue](https://github.com/budimanjojo/home-cluster/issues/new/choose) if you have any questions.

This repository is built off the [k8s-at-home/template-cluster-k3s](https://github.com/k8s-at-home/template-cluster-k3s) repository but reconfigured to not using helm at all.

---

## :art:&nbsp; Cluster components

  - [weave](https://www.weave.works/product/enterprise-kubernetes-platform/): For internal cluster networking.
  - [nfs-subdir-external-provisioner](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner): Provides persistent volumes from NFS server.
  - [SOPS](https://toolkit.fluxcd.io/guides/mozilla-sops/): Encrypts secrets which is safe to store - even to a public repository.
  - [traefik](https://github.com/traefik/traefik): Ingress controller for services.
  - [authelia](https://www.authelia.com/): Full featured authentication server.
  - [cert-manager](https://cert-manager.io/docs/): Configured to create TLS certs for all ingress services automatically using LetsEncrypt.
  - [Kubed](https://appscode.com/products/kubed/): Configmap and Secrets syncronization.
  - [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery): Discover features available on nodes.
  - [smarter-device-manager](https://gitlab.com/arm-research/smarter/smarter-device-manager): Access devices available on nodes.
  - [MetalLB](https://metallb.universe.tf/): Load balancer for bare metal Kubernetes cluster.
  - [keepalived](https://github.com/acassen/keepalived): Loadbalancing & high-availability for master nodes.
  - [haproxy](http://www.haproxy.org/): Load balancer for keepalived master nodes.
  
---

## :open_file_folder:&nbsp; Repository structure

The Git repository contains the following directories under `cluster` and are ordered below by how Flux will apply them.

- **base** directory is the entrypoint to Flux
- **crds** directory contains custom resource definitions (CRDs) that need to exist globally in your cluster before anything else exists
- **core** directory (depends on **crds**) are important infrastructure applications that should never be pruned by Flux
- **apps** directory (depends on **core**) is where your common applications could be placed, Flux will prune resources here if they are not tracked by Git anymore

```
./cluster
├── ./apps
├── ./base
├── ./core
└── ./crds
```

---

## :handshake:&nbsp; Thanks

A lot of inspiration for my cluster came from the people that have shared their clusters over at [awesome-home-kubernetes](https://github.com/k8s-at-home/awesome-home-kubernetes)

---

## Todo List

- [ ] Move everything here
- [v] Use Renovate to auto PR updates
- [ ] Document stuffs
- [ ] Pre-commit checks
