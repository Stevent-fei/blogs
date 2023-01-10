# git 推送分支操作

```shell
git branch release-v0.1.0
git checkout release-v0.1.0
git push -f (推送本地新的分支)
git tag v0.1.0
git push --tags
```

# git 删除本地tag以及远程tag

```shell
删除本地tag: git tag -d xxx

删除远程tag: git push origin :xxx(注意空格)

删除本地分支: git branch -d xxx

删除远程分支: git push origin -d xxx
```

*落日无边江不尽，此身此日更须忙。*