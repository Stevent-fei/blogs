# 放弃本地修改，强制和之前的某次提交同步

有三种情况：

> 1.没有执行 git add的：

可以用命令 ,
```shell
git checkout – filepathname（eg: git checkout – test.md）
```

如果是放弃所有，直接执行

```shell
git checkout .
```

* 此命令用来放弃掉所有还没有加入到缓存区（就是 git add 命令）的修改：内容修改与整个文件删除。
* 但是此命令不会删除掉刚新建的文件。因为刚新建的文件还没已有加入到 git 的管理系统中。所以对于git是未知的。自己手动删除就好了。

> 2.已经执行git add缓存了的：

可以用命令

```shell
git reset HEAD filepathname （比如： git reset HEAD readme.md）
```

同样放弃所有就是

```shell
git reset HEAD .
```

执行完此命令后，文件状态就回归到第一种情况了，此时再按照情况1处理。

> 3.已经用 git commit 提交了的：

可以用命令回退到上一次commit的状态

```shell
git reset --hard HEAD^
```

可以用命令回退到任意版本：

```shell
git reset --hard commitId
```

## 放弃本地修改，强制和远程同步

在使用Git的过程中，有些时候我们只想要git服务器中的最新版本的项目，对于本地的项目中修改不做任何理会，就需要用到Git pull的强制覆盖，具体代码如下：

```shell
$ git fetch --all
$ git reset --hard origin/master 
$ git pull
```