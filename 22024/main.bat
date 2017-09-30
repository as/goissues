@echo off

set batdir=%GOPATH%\src\github.com\as\goissues\22024

dir %batdir%\stage\cygwinok >nul 2>&1 || (echo has22024: fatal: cygwin script isnt running && goto:eof)

set CGO_ENABLED=0
set GOROOT_BOOTSTRAP=C:\Go
set GOROOT=%batdir%\stage\go
set GOBIN=%GOROOT%\bin

mkdir %batdir%\stage
pushd %batdir%\stage

git clone https://go.googlesource.com/go

pushd go

git checkout c8aec4095e089ff6ac50d18e97c3f46561f14f48
git reset --hard
git bisect start
git bisect bad c8aec4095e089ff6ac50d18e97c3f46561f14f48
git bisect good 2b7a7b710f096b1b7e6f2ab5e9e3ec003ad7cd12

:: Git will destroy your terminals console with overengineered conio.h here
git bisect run %batdir%\has22024.bat


