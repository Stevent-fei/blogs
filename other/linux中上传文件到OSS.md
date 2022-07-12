# linux中下载安装ossutil

> 官方链接
> [https://help.aliyun.com/document_detail/120075.html#concept-303829](https://help.aliyun.com/document_detail/120075.html#concept-303829)

```SHELL
wget http://gosspublic.alicdn.com/ossutil/1.6.9/ossutil64
chmod 755 ossutil64
./ossutil64 config
输入配置文件路径/home/自己的user/.ossutilconfig，CH或EN语言，endpoint，accessKeyID，accessKeySecret，accessKeySecre
```

# 传输文件

cp命令可参考
[https://help.aliyun.com/document_detail/120057.html#concept-303810](https://help.aliyun.com/document_detail/120057.html#concept-303810)

```
./ossutil64 cp ./res.tgz oss://bucket名称/
```

*落日无边江不尽，此身此日更须忙。*