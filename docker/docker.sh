#!/usr/bin/env bash
#安装依赖
sudo yum install -y yum-utils device-mapper-persistent-data lvm2

sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

#sudo yum install docker-ce docker-ce-cli containerd.io

#安装docker-ce和docker-ce-cli
sudo yum install docker-ce docker-ce-cli -y

sudo systemctl start docker

docker --version

#设置开机自启

sudo systemctl enable docker

cat >/etc/docker/daemon.json <<EOF
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

#重新加载配置参数
systemctl daemon-reload
#重新启动docker服务
systemctl restart docker
