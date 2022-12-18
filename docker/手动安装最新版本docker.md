# 手动安装最新版本docker

## 直接复制粘贴命令即可，无需修改内容～

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

生效方法
1.重新加载配置参数
systemctl daemon-reload
2.重新启动docker服务
systemctl restart docker
```