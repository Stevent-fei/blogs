# docker常用命令总结

## docker容器操作

1、查看docker容器版本

```dockerfile
docker version
```

2、查看docker容器信息

```dockerfile
docker info
```

3、docker的帮助命令,如果我们遇到不会的命令可以使用以下命令来查找

```dockerfile
docker --help
```

4、在后台以交互分方式启动docker容器

```dockerfile
docker run -d -i -t <imageID> /bin/bash 
```

5、查看全部程序在后台运行

```dockerfile
docker ps -a
```

6、进入docker容器有两个命令

1. 使用docker attach命令

```dockerfile
docker attach 容器id
```

缺点：只要连接容器终止，或者使用了exit这个命令 容器就会退出后台运行

2. 使用docker exec命令

```dockerfile
docker exec -it 容器id /bin/bash
```

这个命令使用exit命令后，不会退出后台，一般使用这个命令

7、关闭启动重启

```dockerfile
docker stop 容器id   #停止容器
 
docker start 容器id  #启动容器
 
docker restart 容器id  #重启容器
```

## 容器进程

## docker镜像操作

## 镜像查看

列出本地镜像

```dockerfile
docker images
```

列出所有的镜像

```dockerfile
docker images -a
```

只显示镜像ID

```dockerfile
docker images -q
```

显示所有的镜像id

```dockerfile
docker images -qa
```

显示镜像摘要信息(DIGEST列)

```dockerfile
docker images --digests
```

显示镜像完整信息

```dockerfile
docker images --no-trunc
```

## 镜像搜索

搜索仓库MySQL镜像

```docker
docker search mysql
```

--filter=stars=600：只显示 starts>=600 的镜像

```dockerfile
docker search --filter=stars=600 mysql
```

--no-trunc 显示镜像完整 DESCRIPTION 描述

```dockerfile
docker search --no-trunc mysql
```

--automated ：只列出 AUTOMATED=OK 的镜像

```dockerfile
docker search  --automated mysql
```

## 镜像下载

下载最新镜像,相当于：docker pull imageName:latest

```dockerfile
docker pull imageName
```

下载仓库所有镜像

```dockerfile
docker pull -a imageName
```

## 镜像删除

单个镜像删除，相当于：docker rmi redis:latest

```dockerfile
docker rmi redis
```

强制删除(针对基于镜像有运行的容器进程)

```dockerfile
docker rmi -f redis
```

多个镜像删除，不同镜像间以空格间隔

```dockerfile
docker rmi -f redis tomcat nginx
```

删除本地全部镜像

```dockerfile
docker rmi -f $(docker images -q)
```

## 镜像构建

编写dockerfile

```dockerfile
cd /docker/dockerfile
vim mycentos
```

构建docker镜像

```dockerfile
docker build -f /docker/dockerfile/mycentos -t mycentos:1.1
```

*落日无边江不尽，此身此日更须忙。*