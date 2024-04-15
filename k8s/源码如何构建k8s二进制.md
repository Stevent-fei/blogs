# 源码构建k8s二进制文件

* 前提条件，一定要有go环境

### 方式一：

使用`make`构建

```shell
make
make all
make all WHAT=cmd/kubelet GOFLAGS=-v
```

### 方式二：

直接使用脚本构建：

```shell
下载k8s源码
git clone https://github.com/kubernetes/kubernetes.git
进入到k8s项目中
cd kubernetes
执行脚本构建所有的二进制
./hack/make-rules/build.sh
```

### 方式三：

单个组件构建，这里使用go build即可

```shell
设置开发环境： 安装 Go 编程语言，并确保版本与 Kubernetes 当前支持的版本一致。查看官方文档了解当前要求的版本。

克隆 Kubernetes 源码仓库： 使用 Git 克隆 Kubernetes 的 GitHub 仓库：

git clone https://github.com/kubernetes/kubernetes.git
cd kubernetes
获取依赖： 在 Kubernetes 项目根目录中，运行以下命令以获取所有依赖项：

make bootstrap
选择要构建的组件： 确定你要构建的组件，例如 kubelet、apiserver 等。每个组件位于 cmd 文件夹下的不同子目录。

构建特定组件： 要为特定组件构建二进制文件，可以在相应的 cmd 子目录下运行 go build。比如，如果你想构建 kubelet：

cd ./cmd/kubelet
go build
这将在当前目录下生成一个名为 kubelet 的二进制文件。

移动二进制文件： 将二进制文件移动到期望的位置，或者使用 go install 直接将其安装到 $GOPATH/bin 目录。

验证二进制文件： 可以通过运行生成的二进制文件来检查是否成功：

./kubelet --help
```

*落日无边江不尽，此身此日更须忙。*