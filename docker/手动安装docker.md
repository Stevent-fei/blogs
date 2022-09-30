一、安装命令
```shell
1、查看linux内核版本。（使用 root 权限登录 Centos ）： 
uname -r
[root@iZgia1btkivmb2Z ~]# uname -r #查看linux内核版本
3.10.0-957.21.3.el7.x86_64         #结果
[root@iZgia1btkivmb2Z ~]#
2、确保yum包更新到最新：
yum -y update

[root@iZgia1btkivmb2Z ~]# yum -y update
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务
---> 软件包 GeoIP.x86_64.0.1.5.0-13.el7 将被 升级
---> 软件包 GeoIP.x86_64.0.1.5.0-14.el7 将被 更新
......
完毕！
[root@iZgia1btkivmb2Z ~]#
3、卸载旧版本(如果安装过旧版本的话)：
sudo yum remove -y docker*

[root@iZgia1btkivmb2Z ~]# sudo yum remove -y docker*
已加载插件：fastestmirror
参数 docker* 没有匹配
不删除任何软件包
[root@iZgia1btkivmb2Z ~]#
4、安装需要的软件包
yum-util 提供yum-config-manager功能，另外两个是devicemapper驱动依赖的：

yum install -y yum-utils

[root@iZgia1btkivmb2Z ~]# yum install -y yum-utils
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务

完毕！
[root@iZgia1btkivmb2Z ~]#
5、设置yum源，并更新 yum 的包索引
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

[root@iZgia1btkivmb2Z ~]# sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
已加载插件：fastestmirror
adding repo from: http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
grabbing file http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo to /etc/yum.repos.d/docker-ce.repo
repo saved to /etc/yum.repos.d/docker-ce.repo
[root@iZgia1btkivmb2Z ~]#
yum makecache fast

[root@iZgia1btkivmb2Z ~]# yum makecache fast
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
......
元数据缓存已建立
[root@iZgia1btkivmb2Z ~]#
6、查看所有仓库中所有docker版本，并选择特定版本安装
yum list docker‐ce ‐‐showduplicates | sort ‐r

[root@iZgia1btkivmb2Z ~]# yum list docker-ce --showduplicates | sort -r
已加载插件：fastestmirror
可安装的软件包
Loading mirror speeds from cached hostfile
docker-ce.x86_64            3:20.10.9-3.el7                     docker-ce-stable
docker-ce.x86_64            3:20.10.8-3.el7                     docker-ce-stable
docker-ce.x86_64            3:20.10.7-3.el7                     docker-ce-stable
docker-ce.x86_64            3:20.10.6-3.el7                     docker-ce-stable
[root@iZgia1btkivmb2Z ~]#
7、安装docker
yum install -y docker-ce-3:19.03.9-3.el7.x86_64

[root@iZgia1btkivmb2Z ~]# yum install -y docker-ce-3:19.03.9-3.el7.x86_64
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务
---> 软件包 docker-ce.x86_64.3.19.03.9-3.el7 将被 安装

完毕！
[root@iZgia1btkivmb2Z ~]#
8、启动docker
systemctl start docker

注意：如果在start出现配置出错的问题，修改/etc/docker/daemon.json文件

[root@iZgia1btkivmb2Z ~]# systemctl start docker
9、设置开机自动启动
systemctl enable docker

10、查看版本信息：
（有client和service两部分表示docker安装启动都成功了 ）
docker version
```

```shell
11-3、如果有 daemon.json 文件，修改即可，如果没有，以下命令操作保存后会自动生成此文件。
[root@iZgia1btkivmb2Z docker]# ls  #查看文件                                   

[root@iZgia1btkivmb2Z docker]# ls
key.json                  #没有指定文件
[root@iZgia1btkivmb2Z docker]# 
  vim daemon.json

[root@iZgia1btkivmb2Z docker]# vim daemon.json
因为我这里是新建，所以执行上述命令后文件内容为空，直接写入以下内容后，

:wq 就可以

{
        "registry-mirrors": ["https://xxxx.mirror.aliyuncs.com"]
}
12、重启docker
systemctl daemon‐reload  

systemctl restart docker 

[root@iZgia1btkivmb2Z docker]# ls
daemon.json  key.json
[root@iZgia1btkivmb2Z docker]# systemctl daemon-reload
[root@iZgia1btkivmb2Z docker]# systemctl restart docker
 13、卸载docker
1 、yum remove ‐y docker*

2 、 rm ‐ rf / etc / systemd / system / docker . service . d
3 、 rm ‐ rf / var / lib / docker
4 、 rm ‐ rf / var / run / docker
```