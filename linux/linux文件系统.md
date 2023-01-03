# 操作系统

登录需要输入用户名和密码，用户名和密码是区分大小写。

```shell
root权限
su root

修改密码:
[root@iZbp144o88s9eacbvib422Z ~]# passwd
更改用户 root 的密码
新的 密码:xxx

查看当前用户:
[root@iZbp144o88s9eacbvib422Z ~]# whoami
root

查看当前在线用户:
可以使用users、who、w命令
[root@iZbp144o88s9eacbvib422Z ~]# users
root
[root@iZbp144o88s9eacbvib422Z ~]# who
root     pts/0        2022-12-27 14:17 (61.185.225.94)
[root@iZbp144o88s9eacbvib422Z ~]# w
 14:20:56 up 1 day,  2:55,  1 user,  load average: 0.09, 0.03, 0.05
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
root     pts/0    61.185.225.94    14:17    0.00s  0.02s  0.00s w

```