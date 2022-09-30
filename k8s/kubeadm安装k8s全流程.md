# kubeadm部署k8s集群

## 1. 环境要求

~~~~
IP                       角色                      安装软件
~~~~

~~~~
172.16.0.227            k8s-master        kube-apiserver kube-schduler kube-controller-manager docker flannel kubelet
~~~~

~~~~
172.16.0.228            k8s-node1         kubelet kube-proxy docker flannel
~~~~

~~~~
172.16.0.229            k8s-node2         kubelet kube-proxy docker flannel
~~~~

## 2. 环境初始化

* 以下所有操作，在三台节点全部执行

### 2.1 关闭防火墙及selinux

```shell
systemctl stop firewalld && systemctl disable firewalld
sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config && setenforce 0
```

### 2.2 关闭 swap 分区

```shell
swapoff -a # 临时
sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab # 永久
```

### 2.3 分别设置主机名

```shell
分别在三台机器上运行
hostnamectl set-hostname k8s-master
hostnamectl set-hostname k8s-node1
hostnamectl set-hostname k8s-node2
```

### 2.4 分别设置host

```shell
vi /etc/hosts

172.16.0.227 k8s-master
172.16.0.228 k8s-node1
172.16.0.229 k8s-node2
```

### 2.5 内核调整,将桥接的IPv4流量传递到iptables的链

```shell
cat >> /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
```

### 2.6 设置系统时区并同步时间服务器

```shell
yum install -y ntpdate
ntpdate time.windows.com
```

## 3. 安装docker

```shell
第一种方案:
1. 安装依赖
sudo yum install -y yum-utils  device-mapper-persistent-data  lvm2

sudo yum-config-manager  --add-repo   https://download.docker.com/linux/centos/docker-ce.repo

sudo yum install docker-ce docker-ce-cli containerd.io

2. 安装docker-ce和docker-ce-cli
sudo yum install docker-ce docker-ce-cli
 
sudo systemctl start docker
 
docker --version

3. 设置开机自启

sudo systemctl enable docker

第二种方案:

wget https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo -O /etc/yum.repos.d/docker-ce.repo
yum -y install docker-ce-20.10.8-3.el8
systemctl enable docker && systemctl start docker
docker --version
```

```shell
安装完docker,把配置文件也要配一下:
cat > /etc/docker/daemon.json <<EOF
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "storage-opts": [
    "overlay2.override_kernel_check=true"
  ]
}
EOF
```

## 4. 添加kubenetes 的yum源

```shell
cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
```

## 5. 安装kubeadm、kubelet和kubectl

```shell
yum install -y kubelet-1.21.0 kubeadm-1.21.0 kubectl-1.21.0
systemctl enable kubelet
```

## 6. 部署Kubenetes Master

只需要在Master 节点执行，这里的apiserve需要修改成自己的master地址

```shell
[root@k8s-master ~]# kubeadm init \
--apiserver-advertise-address=172.16.0.227 \
--image-repository registry.aliyuncs.com/google_containers \
--kubernetes-version v1.21.0 \
--service-cidr=10.1.0.0/16 \
--pod-network-cidr=10.244.0.0/16
```

```shell
输出结果:
[init] Using Kubernetes version: v1.21.0
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
	[WARNING Hostname]: hostname "k8s-masterk8s-node1k8s-node2" could not be reached
	[WARNING Hostname]: hostname "k8s-masterk8s-node1k8s-node2": lookup k8s-masterk8s-node1k8s-node2 on 100.100.2.138:53: no such host
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [k8s-masterk8s-node1k8s-node2 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.1.0.1 172.16.0.227]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [k8s-masterk8s-node1k8s-node2 localhost] and IPs [172.16.0.227 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [k8s-masterk8s-node1k8s-node2 localhost] and IPs [172.16.0.227 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.
[apiclient] All control plane components are healthy after 64.003664 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.21" in namespace kube-system with the configuration for the kubelets in the cluster
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node k8s-masterk8s-node1k8s-node2 as control-plane by adding the labels: [node-role.kubernetes.io/master(deprecated) node-role.kubernetes.io/control-plane node.kubernetes.io/exclude-from-external-load-balancers]
[mark-control-plane] Marking the node k8s-masterk8s-node1k8s-node2 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: v7glh0.osdxr7hnlyeuet0l
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.16.0.227:6443 --token v7glh0.osdxr7hnlyeuet0l \
	--discovery-token-ca-cert-hash sha256:28fe2722c1bd83da7305207a6de4fac79cc8eb70f73a9bf8dd1825b9329ddaf5
```

```shell
根据输出提示操作:
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

```shell
默认token的有效期为24小时,当过期之后,该token就不可用了

如果后续有nodes节点加入,解决方法如下:

重新生成新的token: kubeadm token create

[root@k8s-master ~]# kubeadm token create
0w3a92.ijgba9ia0e3scicg
[root@k8s-master ~]# kubeadm token list
TOKEN                     TTL         EXPIRES                     USAGES                   DESCRIPTION                                                EXTRA GROUPS
v7glh0.osdxr7hnlyeuet0l   22h         2022-09-30T15:14:35+08:00   authentication,signing   The default bootstrap token generated by 'kubeadm init'.   system:bootstrappers:kubeadm:default-node-token

[root@k8s-master ~]# kubeadm token create --print-join-command
kubeadm join 172.16.0.227:6443 --token qr5qe9.vghud1nyga6fb1tk --discovery-token-ca-cert-hash sha256:28fe2722c1bd83da7305207a6de4fac79cc8eb70f73a9bf8dd1825b9329ddaf5
```

```shell
获取ca证书sha256编码hash值
[root@k8s-master ~]# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'

ce07a7f5b259961884c55e3ff8784b1eda6f8b5931e6fa2ab0b30b6a4234c09a
```

## 7. 加入Kubernetes Node

在两个 Node 节点执行

使用kubeadm join 注册Node节点到Matser

kubeadm join 的内容，在上面kubeadm init 已经生成好了

```shell
[root@k8s-node01 ~]# kubeadm join 172.16.0.227:6443 --token qr5qe9.vghud1nyga6fb1tk --discovery-token-ca-cert-hash sha256:28fe2722c1bd83da7305207a6de4fac79cc8eb70f73a9bf8dd1825b9329ddaf5
[preflight] Running pre-flight checks
	[WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at https://kubernetes.io/docs/setup/cri/
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```

## 8. 安装网络查件flannel

```shell
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

查看集群的node状态，安装完网络工具之后，只有显示如下状态，所有节点全部都Ready好了之后才能继续后面的操作

```shell
[root@k8s-master ~]# kubectl get node
NAME                           STATUS   ROLES                  AGE     VERSION
k8s-masterk8s-node1k8s-node2   Ready    control-plane,master   3h16m   v1.21.0
k8s-node1                      Ready    <none>                 101m    v1.21.0
k8s-node2                      Ready    <none>                 101m    v1.21.0
```

## 9. 测试Kubernetes集群

```shell
[root@k8s-master ~]# kubectl create deployment nginx --image=nginx
deployment.apps/nginx created
 
[root@k8s-master ~]# kubectl expose deployment nginx --port=80 --type=NodePort
service/nginx exposed
 
[root@k8s-master ~]# kubectl get pods,svc
NAME                         READY   STATUS    RESTARTS   AGE
pod/nginx-554b9c67f9-wf5lm   1/1     Running   0          24s
 
NAME                 TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
service/kubernetes   ClusterIP   10.1.0.1       <none>        443/TCP        39m
service/nginx        NodePort    10.1.224.251   <none>        80:30328/TCP   9

访问地址:http://NodeIP:Port,此例就是:http://192.168.112.128:30328
```

## 10. 部署rancher

```shell
docker run -d --restart=unless-stopped -p 8080:80 -p 8443:443 -v /home/rancher/rancher:/var/lib/rancher -v /home/rancher/auditlog:/var/log/auditlog --privileged --name rancher2 rancher/rancher:latest
```