# Understand Docker

## OverlayFS

```shell
$ mkdir upper lower merged work
$ echo "from lower" > lower/in_lower.txt
$ echo "from upper" > upper/in_upper.txt
$ echo "from lower" > lower/in_both.txt
$ echo "from upper" > upper/in_both.txt
$ sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
$ cat merged/in_both.txt
```

```shell
$ echo 'new file' > merged/new_file
$ ls -l */new_file
```

```shell
$ rm merged/in_both.txt
$ ls -l upper/in_both.txt  lower/in_both.txt  merged/in_both.txt
```

```shell
$ mount -t overlay overlay -o lowerdir:/dir1:/dir2:/dir3:...:/dir25,upperdir=...
```

## Namespace

```shell
$ lsns -t net
$ cd /proc/25452/ns/
$ nsenter -t <pid> -n ip addr
```

## cgroups

```shell
$ cat /proc/25452/cgroup
11:pids:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
10:freezer:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
9:hugetlb:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
8:perf_event:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
7:blkio:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
6:cpuset:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
5:memory:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
4:devices:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
3:cpu,cpuacct:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
2:net_cls,net_prio:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
1:name=systemd:/kubepods/besteffort/pod8d80a5f8-cb1e-4b28-ba54-393e6b363e20/a99d384f32fc7aeb8a06934e387ed9ea30992676257a61af37d705805f1dffb7
```

```shell
$ docker ps
```

```shell
$ docker inspect <containerid>| grep -i cgroup
"CgroupParent": "kubepods-burstable-podfc9d9da9_7d7a_4970_b306_8ee27f121de1.slice",
```

```shell
$ cd /sys/fs/cgroup/memory/kubepods.slice/kubepods-burstable.slice
```

```shell
$ cd kubepods-burstable-podfc9d9da9_7d7a_4970_b306_8ee27f121de1.slice
```

```shell
$ ls
$ cat memory.limit_in_bytes
1073741824
```