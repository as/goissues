# Repeat Steps

- Install cygwin (any version)
- Open mintty terminal (the standard bash one)
	```
	cd /cygdrive/c/your/gopath/src/github.com/as/goissues/22024
	echo yes > stage/cygwinok
	while true; do ping -n 3600 localhost >/dev/null; . test; done
	
	```
- Open normal cmd.exe
	```
	cd %GOPATH%\src\github.com\as\goissues\22024
	main.bat
	```
	
# Files

```
README.md       This
stage           Go repository target staging area
stage\go        Go repository 
stage\cygwinok  A sanity check that uses the honor system to ensure you installed cygwin
pv.go           A port of pipeviewer affected by issue22024 
test            A while loop that sources the faulty command on each iteration of the bisect
has22024.bat    A pass/fail oracle for the issue, for git bisect
main.bat        Driver that clones the git repo and starts bisecting
```

# Expected Result

```
c05b06a12d005f50e4776095a60d6bd9c2c91fac is the first bad commit
commit c05b06a12d005f50e4776095a60d6bd9c2c91fac
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Feb 10 15:17:38 2017 -0800

    os: use poller for file I/O

    This changes the os package to use the runtime poller for file I/O
    where possible. When a system call blocks on a pollable descriptor,
    the goroutine will be blocked on the poller but the thread will be
    released to run other goroutines. When using a non-pollable
    descriptor, the os package will continue to use thread-blocking system
    calls as before.

    For example, on GNU/Linux, the runtime poller uses epoll. epoll does
    not support ordinary disk files, so they will continue to use blocking
    I/O as before. The poller will be used for pipes.

    Since this means that the poller is used for many more programs, this
    modifies the runtime to only block waiting for the poller if there is
    some goroutine that is waiting on the poller. Otherwise, there is no
    point, as the poller will never make any goroutine ready. This
    preserves the runtime's current simple deadlock detection.

    This seems to crash FreeBSD systems, so it is disabled on FreeBSD.
    This is issue 19093.

    Using the poller on Windows requires opening the file with
    FILE_FLAG_OVERLAPPED. We should only do that if we can remove that
    flag if the program calls the Fd method. This is issue 19098.

    Update #6817.
    Update #7903.
    Update #15021.
    Update #18507.
    Update #19093.
    Update #19098.

    Change-Id: Ia5197dcefa7c6fbcca97d19a6f8621b2abcbb1fe
    Reviewed-on: https://go-review.googlesource.com/36800
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

:040000 040000 4abda0d5912e06198d1679c95d85cb2d6a47d8f3 8ba0a385b69c42debcce51b54cc14ddcf12b9e7a M      src
bisect run success
```