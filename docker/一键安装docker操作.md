在 Linux 上安装 Docker 是常见的安装场景，并且安装过程非常简单。

通常难点在于 Linux 不同发行版之间的轻微区别，比如 Ubuntu 和 CentOS 之间的差异。

接下来的示例基于 Ubuntu 版本 Linux，同样适用于更低或者更高的版本。

理论上，下面的示例在 CentOS 的各种版本上也是可以执行的。至于 Linux 操作系统是安装在自己的数据中心，还是第三方公有云，或是笔记本的虚拟机上，都没有任何的区别。唯一需求就是这台机器是 Linux 操作系统，并且能够访问 https://get.docker.com。

首先读者需要选择安装的 Docker 版本。当前有两个版本可供选择：社区版（Community Edition，CE）和企业版（Enterprise Edition，EE）。

Docker CE 是免费的，并且是接下来示例中将要使用的版本。

Docker EE 包含 Docker CE 中的全部功能，还包括了商业支持以及与其他 Docker 产品的集成，比如 Docker 可信镜像库和通用控制面板。

下面的例子使用wget命令来运行一个 Shell 脚本，完成 Docker CE 的安装。

更多其他在 Linux 上安装 Docker 的方式，可以打开 Docker 主页面（www.docker.com），单击页面中 Get Started 按钮来获取。

注：在开始下面的步骤之前，要确认系统升级到最新的包，并且打了相应的安全补丁。

1) 在 Linux 机器上打开一个新的 Shell。

2) 使用wget从 https://get.docker.com 获取并运行 Docker 安装脚本，然后采用 Shell 中管道（pipe）的方式来执行这个脚本。

```shell
$ wget -qO- https://get.docker.com/ | sh

modprobe: FATAL: Module aufs not found /lib/modules/4.4.0-36-generic
+ sh -c 'sleep 3; yum -y -q install docker-engine'
<Snip>
If you would like to use Docker as a non-root user, you should
now consider adding your user to the "docker" group with
something like:

sudo usermod -aG docker your-user

Remember that you will have to log out and back in...
```

3) 最好通过非 root 用户来使用 Docker。这时需要添加非 root 用户到本地 Docker Unix 组当中。

下面的命令展示了如何把名为 npoulton 的用户添加到 Docker 组中，以及如何确认操作是否执行成功。
```shell
$ sudo usermod -aG docker npoulton

$ cat /etc/group | grep docker
docker:x:999:npoulton
```
如果当前登录用户就是要添加到 Docker 组中的用户的话，则需要重新登录，组权限设置才会生效。

至此 Docker 已经在 Linux 上安装成功。运行下面命令来确认安装结果。
```shell
$ docker --version
Docker version 18.01.0-ce, build 03596f5

$ docker system info
Containers: 0
Running: 0
Paused: 0
Stopped: 0
Images: 0
Server Version: 18.01.0-ce
Storage Driver: overlay2
Backing Filesystem: extfs
<Snip>
```
如果上述步骤在自己的 Linux 发行版中无法成功执行，可以访问 Docker Docs 网站（https://docs.docker.com）并单击与自己的版本相关的那个链接。

接下来页面会跳转到 Docker 官方提供的适合当前版本的安装指南页面，这个安装指南通常会保持更新。

但是需要注意，Docker 网站上提供的指令使用了包管理器，相比前面的例子需要更多的步骤才能完成安装操作。

实际上，如果读者使用浏览器打开网页 https://get.docker.com，会发现这其实就是一个 Shell 脚本，脚本中已经帮读者定义好了安装相关的指令，包括设置 Docker 为系统开启自启动。

如果未从官方 Docker 仓库下载源码，则最终安装的可能是 Docker 的一个复制版本。过去一些 Linux 发行商选择复制了 Docker 的代码，并基于此开发了一些定制化的版本。

需要注意类似的情况，因为运行一个与 Docker 官方版本不同的复制版，可能遇到异常退出的情况。

如果本意就是采用该版本运行，那这不是问题。但是如果本意并非如此，复制版本中发行商提交的一些改动可能导致其版本无法与 Docker 官方版本相兼容。这样就无法从 Docker 公司或者 Docker 公司授权的合作伙伴那里获得商业支持。

