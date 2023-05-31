# git代理设置 
[不建议设置全局代理, 多环境下可能混乱 注意使用协议，能使用ssh协议尽量使用ssh协议方式 具体的代理ip和端口视情况调整，以下用端口1080示例 HTTP协议代理设置 
1、使用命令直接设置代理 --global 表示全局，不需要可以不加

```shell
git config --global https.proxy ***
```

1、例子：

# socks

```shell
git config --global http.proxy 'socks5://127.0.0.1:1080' 
git config --global https.proxy 'socks5://127.0.0.1:1080'
```

# http

```shell
git config --global http.proxy http://127.0.0.1:1080
git config --global https.proxy https://127.0.0.1:1080
```

# 只对github.com使用代理，其他仓库不走代理

```shell
git config --global http.https://github.com.proxy socks5://127.0.0.1:1080 
git config --global https.https://github.com.proxy socks5://127.0.0.1:1080
```

# 取消github代理

```shell
git config --global --unset http.https://github.com.proxy
git config --global --unset https.https://github.com.proxy
```

2、直接修改~/.gitconfig文件

```shell
[http]
proxy = socks5://127.0.0.1:1080
[https]
proxy = socks5://127.0.0.1:1080
```

3、取消代理

```shell
git config --global --unset http.proxy 
git config --global --unset https.proxy 
```

# SSH协议代理设置 修改ssh配置文件~

```shell
~/.ssh/config 没有的话新建一个文件 
Windows ssh配置文件路径:C:\Users\你的用户名\.ssh\config 
Linux ssh配置文件路径:/home/你的用户名/.ssh/config
```

```shell
ProxyCommand connect -S 代理地址:端口 %h %p 
```

例子

```shell
# 全局

# ProxyCommand connect -S 127.0.0.1:1080 %h %p

# 只为特定域名设定

Host github.com gitlab.com 
ProxyCommand connect -S 127.0.0.1:1080 %h %p 
```

-S 代表走socks代理。（ -H 实现http和https的仓库的克隆） 多个地址设置：Host后面使用空格隔开，而不是,

多账号设置 生成公私钥，并在对应账号配置公钥

```shell
// 生成两个邮箱对应的ssh公私钥 
ssh-keygen -t ed25519 -C "1@email"
ssh-keygen -t ed25519 -C "2@email"
```

配置例子：

```shell
# Host：仓库网站的别名，随意取

# HostName：仓库网站的域名（PS：IP 地址应该也可以）

# User：仓库网站上的用户名

# IdentityFile：私钥的绝对路径
```