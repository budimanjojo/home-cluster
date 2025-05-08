<div align="center">

### My home Kubernetes Talos cluster :sailboat:

... managed with Flux :robot:

</div>

## :book:&nbsp; Overview

This repository _is_ my home Kubernetes cluster in a declarative state.
[Flux](https://github.com/fluxcd/flux2) watches my [cluster](./cluster/) directory and makes the changes to my cluster based on the YAML manifests.

Feel free to open a [Github issue](https://github.com/budimanjojo/home-cluster/issues/new/choose) if you have any questions.

This repository is built off the [k8s-at-home/template-cluster-k3s](https://github.com/k8s-at-home/template-cluster-k3s) repository.

---

## :art:&nbsp; Cluster components

### Cluster management

- [Talos](https://www.talos.dev): Built using [talhelper](https://github.com/budimanjojo/talhelper)
- [fluxcd](https://fluxcd.io/): Sync kubernetes cluster with this repository.
- [SOPS](https://toolkit.fluxcd.io/guides/mozilla-sops/): Encrypts secrets which is safe to store - even to a public repository.

### Networking

- [Cilium](https://cilium.io): For internal cluster networking, also as load balancer to expose services.
- [cert-manager](https://cert-manager.io/docs/): Configured to create TLS certs for all ingress services automatically using LetsEncrypt.
- [authelia](https://www.authelia.com/): Full featured authentication server.

### Storage

- [rook-ceph](https://rook.io): Cloud native distributed block storage for Kubernetes
- [nfs-subdir-external-provisioner](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner): Provides persistent volumes from NFS server.

### Host devices access

- [Intel GPU plugin](https://github.com/intel/intel-device-plugins-for-kubernetes): Access intel gpu available on nodes.
- [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery): Discover features available on nodes.

### Metrics

- [Prometheus](https://prometheus.io/): Scraping metrics from the entire cluster
- [Grafana](https://grafana.com): Visualization for the metrics from Prometheus

---

## :open_file_folder:&nbsp; Repository structure

The Git repository contains the following directories under `cluster` and are ordered below by how Flux will apply them.

```
./cluster
├── ./base    # entrypoint to Flux
└── ./apps    # everything is here
```

Inside the [apps](./cluster/apps/) directory, I divided all the apps using their namespaces.
Every app will have its own "Fluxtomization" file that describe their manifests and dependencies.

---

## :satellite:&nbsp; Network structure

Incoming http and https traffics from outside of my network are forwarded from OPNSense firewall into `envoy-gateway` pod with a LoadBalancer service using MetalLB layer2 implementation.
So, basically this is how the http(s) traffic flows:
```
Internet -> OPNSense firewall -> envoy-gateway service -> Kubernetes pod
```
Ingress-nginx service is using `Local` `externalTrafficPolicy` so I can track the real IP of clients trying to access my services.
For important backend services like my OPNSense, I use `nginx.ingress.kubernetes.io/whitelist-source-range` annotation to only allow access from my internal networks.
My certificates are managed with cert-manager using LetsEncrypt as the CA.

---

## :lock_with_ink_pen:&nbsp; Secret and configmaps management

Secrets are encrypted using [sops](https://github.com/mozilla/sops) before being pushed into this repository.
The encrypted secrets are then decrypted by sops using the private key inside the cluster.
For encryption/decryption, I use [age](https://github.com/FiloSottile/age).
The public key to encrypt the secret is in [.sops.yaml](.sops.yaml).
Secrets environment variables for the cluster are in [cluster-secret-vars.yaml](./cluster/base/config/cluster-secret-vars.sops.yaml).
The non secret variables are in [cluster-vars.yaml](./cluster/base/config/cluster-vars.yaml).

---

## :bar_chart:&nbsp; Metrics and chart management

Metrics scraping for the cluster are done using Prometheus.

Dashboards included in my cluster are:

- The provided dashboard from [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)
- Fluxcd dashboard from [here](https://github.com/fluxcd/flux2-monitoring-example/tree/main/monitoring/configs/dashboards)
- Rook-ceph dashboards from [here](https://www.rook.io/docs/rook/v1.10/Storage-Configuration/Monitoring/ceph-monitoring/?h=grafana#grafana-dashboards)

To add your own dashboard, create a configmap with the data include the json file of the dashboard and add label `grafana_dashboard: "1"` to the manifest.
The sidecar container from this [image](https://github.com/kiwigrid/k8s-sidecar) will mount the dashboard into your grafana pod.

---

## :handshake:&nbsp; Thanks

A lot of inspiration for my cluster came from this [awesome template](https://github.com/onedr0p/flux-cluster-template)
---

## Todo List

- [ ] Use redis operator
