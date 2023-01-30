# go-syscall-with-strace

This is a sample of an application built in Go using the strace command to see what system-calls are being call.

![Go Version](https://img.shields.io/badge/go-1.19-blue.svg)

## Go version

```shell script
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.5 LTS
Release:        20.04
Codename:       focal

$ go version
go version go1.19.4 linux/amd64
```

## How to Run

### Using [go-task](https://taskfile.dev/)

#### run

```sh
$ task run
task: [run] go build -o ./app
task: [run] strace -y -f -e openat,read,write,clone,close ./app 2>&1 1>/dev/null | tee syscall.log
openat(AT_FDCWD, "/sys/kernel/mm/transparent_hugepage/hpage_pmd_size", O_RDONLY) = 3</sys/kernel/mm/transparent_hugepage/hpage_pmd_size>
read(3</sys/kernel/mm/transparent_hugepage/hpage_pmd_size>, "2097152\n", 20) = 8
close(3</sys/kernel/mm/transparent_hugepage/hpage_pmd_size>) = 0
clone(child_stack=0xc000074000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLSstrace: Process 7511 attached
, tls=0xc000064090) = 7511
[pid  7510] clone(child_stack=0xc000076000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLSstrace: Process 7512 attached
, tls=0xc000064490) = 7512
[pid  7510] clone(child_stack=0xc000070000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLSstrace: Process 7513 attached
, tls=0xc000064890) = 7513
[pid  7512] clone(child_stack=0xc000112000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLSstrace: Process 7514 attached
, tls=0xc000100090) = 7514
[pid  7510] openat(AT_FDCWD, "out.txt", O_RDWR|O_CREAT|O_TRUNC|O_CLOEXEC, 0666) = 3</workspace/go-syscall-with-strace/out.txt>
[pid  7510] write(3</workspace/go-syscall-with-strace/out.txt>, "1\n", 2) = 2
[pid  7510] write(3</workspace/go-syscall-with-strace/out.txt>, "2\n", 2) = 2
[pid  7510] write(3</workspace/go-syscall-with-strace/out.txt>, "3\n", 2) = 2
[pid  7512] write(3</workspace/go-syscall-with-strace/out.txt>, "4\n", 2) = 2
[pid  7512] write(3</workspace/go-syscall-with-strace/out.txt>, "5\n", 2) = 2
[pid  7514] close(3</workspace/go-syscall-with-strace/out.txt>) = 0
[pid  7514] +++ exited with 0 +++
[pid  7513] +++ exited with 0 +++
[pid  7512] +++ exited with 0 +++
[pid  7511] +++ exited with 0 +++
+++ exited with 0 +++
```

#### clean

```sh
$ task clean
task: [clean] rm -f ./app out.txt syscall.log
```
