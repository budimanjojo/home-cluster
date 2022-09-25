<div align="center">

<img src="https://raw.githubusercontent.com/siderolabs/talos/main/website/assets/icons/logo.svg" align="center" width="144px"/>

### Talos Linux cluster

... managed with Talhelper :robot:

</div>

## :book:&nbsp; Overview

This directory contains my [Talos](https://www.talos.dev/) Kubernetes cluster in declarative state.
I use my own tool [Talhelper](https://github.com/budimanjojo/talhelper) to create the `machineconfig` files of all my nodes.
The secrets are encrypted with [SOPS](https://toolkit.fluxcd.io/guides/mozilla-sops/).

Feel free to open a [Github issue](https://github.com/budimanjojo/home-cluster/issues/new/choose) if you have any questions.

---

## :scroll:&nbsp; How to apply

1. Prepare your nodes with `Talos Linux`
2. Install `talhelper`.
3. Create your own [talconfig.yaml](https://github.com/budimanjojo/home-cluster/blob/main/talos/talconfig.yaml).
4. Run `talhelper gensecret > talsecret.sops.yaml` if you don't have `machineconfig` before or `talhelper gensecret -f <your-machineconfig.yaml> > talsecret.sops.yaml` if you already have one.
5. Run `sops -e -i talsecret.sops.yaml` to encrypt your secrets (make sure you already have your own `.sops.yaml`) file.
6. Run `talhelper genconfig` and the files will be generated in `./clusterconfig` directory by default.
7. Copy the generated `./clusterconfig/talosconfig` to your `~/.talos/config`.
8. Run `talosctl -n <node-ip> apply-config --insecure ./clusterconfig/<clustername>-<hostname>.yaml` on each of your node. Don't forget to run `talosctl -n <node-ip> bootstrap` on one of your controlplane node.
9. Push your current directory to your git repository of choice. :wink:

---

## :memo:&nbsp; After bootstrap

After you're done with bootstrapping, you can now install your `Kubernetes CNI` of your choice.
If you want to use cilium, you can look at my [cni](./cni) directory.
You can do `kubectl kustomize --enable-helm ./cni | kubectl apply -f -` to do this.

If you also want to deploy [kubelet-csr-approver](https://github.com/postfinance/kubelet-csr-approver) like I do, you can also do the above step to my [kubelet-csr-approver](./kubelet-csr-approver) directory.

Now, you can continue to work on your cluster.
Check out my [cluster](../cluster) directory to see how I manage my cluster with [Flux](https://github.com/fluxcd/flux2).
