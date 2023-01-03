# git合并多次commit提交
1. 直接删除上一次提交

```shell
git reset --hard HEAD^
git push -f
```

2. 合并多次提交

```shell
git rebase -i HEAD~3
```

执行返回如下结果：

```shell
pick d1232342cf1f9 fix: 第一次提交

pick 479711234f216 fix: 第二次提交

pick fb282144124c8d fix: 第三次提交
```

这时候我们需要修改commit，按 `i` 进入编辑模式，把需要修改的commit前面的pick换成s或者squash，像这样

命令说明：

> p，pick：保留该commit（缩写:p）
>  r，reword：保留该commit，但我需要修改该commit的注释（缩写:r）
>  e，edit：保留该commit, 但我要停下来修改该提交(不仅仅修改注释)（缩写:e）
>  s，squash：将该commit和前一个commit合并（缩写:s）
>  f，fixup：将该commit和前一个commit合并，但我不要保留该提交的注释信息（缩写:f）
>  x，exec：执行shell命令（缩写:x）
>  d，drop：我要丢弃该commit（缩写:d）

经常用到的是`edit`，`squash`，`fixup`

示例：

```undefined
pick d1232342cf1f9 fix: 第一次提交

s 479711234f216 fix: 第二次提交

pick fb282144124c8d fix: 第三次提交
```

这里修改完以后会直接跳到修改commit 名字的那页，我们把需要的commit名字留下，不需要的可以`#`注释掉

3. 推送到远程仓库

```undefined
git push -f
```

*落日无边江不尽，此身此日更须忙。*