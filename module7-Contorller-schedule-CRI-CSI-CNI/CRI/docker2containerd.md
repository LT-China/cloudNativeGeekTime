### 参考： https://kubernetes.io/zh/docs/setup/production-environment/container-runtimes/#containerd

```
### 安装配置 containerd 先决条件
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# 设置必需的 sysctl 参数，这些参数在重新启动后仍然存在。
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
# 应用 sysctl 参数而无需重新启动
sudo sysctl --system
```

### Stop service

```sh
systemctl stop kubelet
systemctl stop docker
systemctl stop containerd
```

### Create containerd config folder

```sh
sudo mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
```

### Update default config

```sh
vi /etc/containerd/config.toml
sed -i s#k8s.gcr.io/pause:3.5#registry.aliyuncs.com/google_containers/pause:3.5#g /etc/containerd/config.toml
sed -i s#'SystemdCgroup = false'#'SystemdCgroup = true'#g /etc/containerd/config.toml
```

### Edit kubelet config and add extra args

```sh
vi /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
Environment="KUBELET_EXTRA_ARGS=--container-runtime=remote --container-runtime-endpoint=unix:///run/containerd/containerd.sock --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.5"
```

### Restart

```sh
systemctl daemon-reload
systemctl restart containerd
systemctl restart kubelet
```

### Config crictl to set correct endpoint

```sh
cat <<EOF | sudo tee /etc/crictl.yaml
runtime-endpoint: unix:///run/containerd/containerd.sock
EOF
```