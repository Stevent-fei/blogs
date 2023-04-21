# kubernetes常见面试题

# 1. 什么是Kubernetes？

* Kubernetes是一种开源容器编排和管理平台，可以自动化应用程序的部署、扩展和管理。

# 2. Kubernetes的优点是什么？

* 自动化应用程序的部署、扩展和管理
* 支持多租户和多环境部署
* 容错性高，支持Rolling Update、Blue/Green Deployment等
* 支持水平扩展和弹性伸缩
* 灵活的网络配置和服务发现
* 支持多种存储后端

# 3. Kubernetes中的Pod是什么？

* Kubernetes中的Pod是最小的可部署单元，由一个或多个容器组成。

# 4. Kubernetes中的Service是什么？

* Kubernetes中的Service是一种虚拟网络抽象，用于暴露Pod的网络服务。

# 5. Kubernetes中的Deployment是什么？

* Kubernetes中的Deployment是一种控制器，用于管理Pod的副本和更新。

# 6. Kubernetes中的ConfigMap和Secret是什么？

* Kubernetes中的ConfigMap用于存储应用程序的配置信息。
* Kubernetes中的Secret用于存储敏感数据，如密码和证书。

# 7. Kubernetes中的Node是什么？

* Kubernetes中的Node是运行容器的主机。

# 8. Kubernetes中的Namespace是什么？

* Kubernetes中的Namespace是用于组织和隔离Kubernetes对象的虚拟集合。

# 9. Kubernetes中的Volume是什么？

* Kubernetes中的Volume是一种抽象，用于存储容器中的数据。

# 10. Kubernetes中的Ingress是什么？

* Kubernetes中的Ingress是一种控制器，用于管理HTTP和HTTPS流量的路由。

# 11. 常用组件：

* etcd：分布式键值存储系统，用于存储集群中所有节点的配置信息和状态信息。

* kube-apiserver：Kubernetes API Server，是整个Kubernetes的核心组件之一，它提供了集群的中央控制器，用于维护整个系统的状态并控制整个集群的生命周期。

* kube-scheduler：Kubernetes Scheduler，是用来进行Pod的调度的组件，通过Kubernetes API Server接收到新的Pod创建请求后，将选择最合适的节点来创建Pod。

* kube-controller-manager：Kubernetes Controller Manager，是Kubernetes中的控制器管理器，主要是负责对集群中的各种资源进行自动化管理，如Replication
  Controller、Node Controller、Endpoints Controller等。

* kubelet：Kubernetes Node Agent，是Kubernetes中的节点代理组件，主要负责监控节点的状态以及管理容器的生命周期。

* kube-proxy：Kubernetes Network Proxy，是Kubernetes中的网络代理组件，主要用于提供一种简单的方式来实现Pod之间的网络通信。

# 12. Kubernetes原理：

Kubernetes是一个容器编排平台，它的核心思想是将应用程序打包成一个或多个容器，然后将这些容器部署到集群中的各个节点上。Kubernetes主要由以下几个部分组成：

* Master节点：Master节点是Kubernetes的控制中心，它负责管理整个集群的生命周期，包括部署、管理、监控和维护等。

* Node节点：Node节点是Kubernetes集群中的工作节点，它用于运行应用程序和服务。

* Pod：Pod是Kubernetes集群中的最小单元，它是由一个或多个容器组成的，用于运行应用程序和服务。

* Service：Service是Kubernetes中的网络服务，它用于管理Pod的访问方式和访问策略，实现Pod之间的通信。

Kubernetes的工作流程如下：

    1. 用户通过Kubernetes API Server向Master节点提交一个Pod的创建请求。

    2. Master节点接收到请求后，通过kube-scheduler进行调度，选择最合适的Node节点。

    3. Node节点接收到Pod创建请求后，通过kubelet运行Pod中的容器。

    4. Pod中的容器运行完毕后，kubelet将容器的状态更新到API Server中。

    5. 用户可以通过Kubernetes API Server查询Pod和容器的状态信息，或者通过Kubernetes Dashboard进行可视化管理和监控。

# 13. kubernetes中的镜像拉取策略分为哪些：

# 14. k8s四大组件：

* Master节点：负责管理整个集群的资源和状态，包括API服务器，调度器和控制器
* Node节点：运行在物理或虚拟机上的服务器，负责执行容器化应用
* Etcd：分布式键值存储系统，用于存储集群状态的配置信息
* Container Runtime：负责管理Node节点上运行容器的组件，常用的有Docker和containerd

# 15. k8s集群节点主要分为控制平面节点(Master)和工作节点

一、控制平面节点主要包括以下组件： 1、kube-apiserver：集群中所有资源的统一访问入口；

    2、kube-scheduler：将新创建的pod调度到合适的节点上

    3、kube-controller-manager：集群中所有资源对象的自动化控制中心；

    4、etcd：保存集群中的所有资源对象的数据、

二、工作节点主要包括以下组件： 1、kubelet：负责pod对应的容器的创建、启停等任务，同时与Master节点密切协作，实现集群管理的基本功能；

2、kube-proxy：将对service的访问转发到后端的一组pod上；

3、容器运行时（Container Runtime）：容器运行时是负责运行容器的软件，k8s支持许多容器运行时，常用的是docker。

k8s架构原理：

```text
Kubernetes（k8s）是一个高度可扩展的容器编排平台，可以自动化应用程序的部署、扩展和管理。它的架构原理可以分为以下几个方面：

控制平面与数据平面：Kubernetes架构中有两个主要组件：控制平面和数据平面。控制平面负责管理和协调集群中所有的节点和容器，而数据平面则处理实际的容器通信和网络数据流。

Master节点和Worker节点：Kubernetes集群由Master节点和Worker节点组成。Master节点管理整个集群，包括控制平面和数据平面。Worker节点则是负责运行应用程序和服务的节点。

节点组件：Kubernetes节点上的组件包括kubelet、kube-proxy和容器运行时。kubelet是Kubernetes节点上的代理程序，负责与Master节点通信并执行Pod的生命周期。kube-proxy是负责网络代理的组件，它可以将流量转发到正确的Pod。容器运行时则负责运行Pod中的容器。

Pod：Pod是Kubernetes中最小的可部署单元，它包含一个或多个容器。Pod中的容器共享网络和存储空间，并可以通过本地主机上的IPC、PID和网络命名空间进行通信。

控制器：Kubernetes控制器是一种控制Pod和ReplicaSet（副本集）、Deployment、StatefulSet等对象的机制。它们确保Pod数量和状态与期望的相符，并可以根据需要自动创建、更新和删除Pod。

总之，Kubernetes的架构原理是通过Master节点和Worker节点之间的通信，以及控制平面和数据平面之间的协调来实现高度自动化的容器编排和管理。
```

# k8s中pod node container之间的区别

在Kubernetes中，Pod、Node和Container是三个不同的概念：

Pod：是Kubernetes中最小的部署单元，包含一个或多个紧密关联的容器和一个共享的存储卷。Pod提供了一个抽象层，使得容器在Kubernetes中可以像一个单独的逻辑单元进行部署、调度和管理。

Node：是Kubernetes中的工作节点，它是物理机器或虚拟机，用于运行Pod中的容器。Node包括物理资源（例如CPU、内存、磁盘等），以及运行容器所需的Kubernetes服务、代理和其他必要的组件。

Container：是Kubernetes中的应用程序容器，它可以在Pod中运行。容器是一个独立的可执行软件包，包含应用程序、运行时环境和所有依赖项。

总结来说，Pod是Kubernetes中最小的部署单元，Node是运行Pod的物理节点，而Container则是运行在Pod中的应用程序容器。在Kubernetes中，Pod、Node和Container都是非常重要的概念，理解它们之间的关系和作用可以帮助我们更好地使用Kubernetes进行容器编排和管理。

# kubectl apply 和 kubectl create有什么区别

```text
kubectl apply 和 create 都可以用来创建 Kubernetes 对象，但是它们之间有一些区别：

apply 可以用于更新对象，而 create 只能用于创建新对象。如果对象已经存在，使用 create 会导致错误。

apply 可以使用 YAML 或 JSON 文件来定义对象，而 create 只能使用命令行参数。

apply 会在对象已经存在时进行更新，而 create 会在对象已经存在时抛出错误。

apply 可以进行部分更新，而 create 只能进行完整的创建。如果对象已经存在，apply 可以只更新部分字段而不影响其他字段。

总之，apply 更加灵活和强大，而 create 更加简单和直接。选择哪一个取决于你的需求和使用场景。
```