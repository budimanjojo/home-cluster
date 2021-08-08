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

### Cluster management
  - [kubeadm](https://kubernetes.io/docs/reference/setup-tools/kubeadm/): Create and update cluster (planning to move to kubespray soon).
  - [fluxcd](https://fluxcd.io/): Sync kubernetes cluster with this repository.
  - [SOPS](https://toolkit.fluxcd.io/guides/mozilla-sops/): Encrypts secrets which is safe to store - even to a public repository.
### Networking
  - [weave](https://www.weave.works/product/enterprise-kubernetes-platform/): For internal cluster networking.
  - [MetalLB](https://metallb.universe.tf/): Load balancer for bare metal Kubernetes cluster.
  - [keepalived](https://github.com/acassen/keepalived): Loadbalancing & high-availability for master nodes.
  - [haproxy](http://www.haproxy.org/): Load balancer for keepalived master nodes.
  - [cert-manager](https://cert-manager.io/docs/): Configured to create TLS certs for all ingress services automatically using LetsEncrypt.
  - [traefik](https://github.com/traefik/traefik): Ingress controller for services.
  - [authelia](https://www.authelia.com/): Full featured authentication server.
### Storage
  - [nfs-subdir-external-provisioner](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner): Provides persistent volumes from NFS server.
### Host devices access
  - [Intel GPU plugin](https://github.com/intel/intel-device-plugins-for-kubernetes): Access intel gpu available on nodes.
  - [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery): Discover features available on nodes.
  - [smarter-device-manager](https://gitlab.com/arm-research/smarter/smarter-device-manager): Access devices available on nodes.
### Metrics
  - [Prometheus](https://prometheus.io/): Scraping metrics from the entire cluster
  - [Grafana](https://grafana.com): Visualization for the metrics from Prometheus

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

## :lock_with_ink_pen:&nbsp; Secret and configmaps management

Secrets are encrypted using [sops](https://github.com/mozilla/sops) before being pushed into this repository. The encrypted secrets are then decrypted by sops using the private key inside the cluster. Encryption/decryption are done using [age](https://github.com/FiloSottile/age). The public key to encrypt the secret can be accessed in [.sops.yaml](.sops.yaml). Secrets environment variables for the cluster are done by the built-in fluxcd envsubs feature can be accessed in [cluster-secret-vars.yaml](.cluster/base/cluster-secret-vars.yaml). The non secret equivalent can be accessed in [cluster-vars.yaml](.cluster/base/cluster-vars.yaml).

---

## :bar_chart:&nbsp; Metrics and chart management

Metrics scraping for the cluster are done using Prometheus. The manifests are generated from my own maintained [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus), can be accessed [here](https://github.com/budimanjojo/kube-prometheus).

Dashboards included in my cluster are:
  - The provided dashboard from [kube-prometheus](https://github.com/prometheus-operator/kube-prometheus)
  - Traefik dashboard from [grafanalabs](https://grafana.com/grafana/dashboards/12250)
  - Fluxcd dashboard from [here](https://github.com/fluxcd/flux2/tree/main/manifests/monitoring/grafana/dashboards)

To add your own dashboard, create a configmap with the data include the json file of the dashboard and add label `grafana_dashboard: "1"` to the manifest. The sidecar container from this [image](https://github.com/kiwigrid/k8s-sidecar) will mount the dashboard into your grafana pod.

---

## :handshake:&nbsp; Thanks

A lot of inspiration for my cluster came from the people that have shared their clusters over at [awesome-home-kubernetes](https://github.com/k8s-at-home/awesome-home-kubernetes)

---

## Todo List

- [x] Move everything here
- [x] Use Renovate to auto PR updates
- [x] Document stuffs
- [x] Replace kubed with fluxcd envsubs
- [x] Replace non standard versioning containers
- [x] Use security context for permissions instead
- [ ] Planning to use kubespray to provision cluster
