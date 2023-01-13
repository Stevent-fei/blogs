# 使用 UPX 压缩可执行文件

> [UPX](https://upx.github.io/) 可以有效地对可执行文件进行压缩，并且压缩后的文件可以直接由系统执行，支持多系统和平台。 使用 UPX 来压缩可执行文件是一种减少发布包大小的有效方式。

## 安装

* 从 [github release page](https://github.com/upx/upx/releases) 下载预编译的二进制文件

* macOS 可以使用 brew 安装：

```shell
brew install upx
```

### 使用 、压缩

```shell
upx [options] yourfile
```

upx 对文件的默认操作即为压缩，使用上述命令会使用默认参数压缩并替换文件 yourfile。 upx 支持如下可选参数：

`-1[23456789]`：不同的压缩级别，数值越高压缩率越高，但耗时更长。对于小于 512 KiB 的文件默认使用 `-8`，其他的默认为 `-7`。
`--best`：最高压缩级别
`--brute`：尝试使用各种压缩方式来获取最高压缩比
`--ultra-brute`：尝试使用更多的参数来获取更高的压缩比
`-o [file]`：将压缩文件保存为 [file]

### 解压

```shell
upx -d [yourfile]
```

## 优劣

压缩的程序占用更少的硬盘空间，但会在打开时消耗更多的 CPU 资源，在运行时占用更多的内存（或 swap 空间、/tmp 存储等）。

### 优点

* UPX 可以压缩各种类型的可执行文件
* 压缩后的文件可以直接由操作系统执行
* 压缩过程不会修改源文件，也就意味着解压后直接可以得到原始文件
* 不会产生额外的动态库调用

### 缺点

* 运行的程序不会共享数据段（汇编），所以多实例运行的程序不适合压缩
* 使用 `ldd` 和 `size` 命令无法获取到程序的有效信息

## 原理

为什么压缩后的文件可由系统直接执行？ UPX 将程序压缩，并在头部加入解压的程序，具体的原理可以参看参考[2]。 在 Linux 系统中可以使用 `strings` 命令查看可执行文件的内容，通过查看 UPX 压缩后的程序可以看到，UPX
在文件中写入了自己的特征码。

## 参考 
1: [UPX manual](https://github.com/upx/upx/blob/master/doc/upx.pod)
2: [Packers, How They Work, Featuring UPX](https://dzone.com/articles/packers-how-they-work-featuring-upx)