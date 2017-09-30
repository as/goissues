@echo off

del %batdir%\stage\fixed
del %batdir%\stage\pv.exe

pushd %GOROOT%\src
call make.bat
popd

%GOBIN%\go tool compile %batdir%\pv.go
%GOBIN%\go tool link pv.o
copy /Y a.out ..\pv.exe

taskkill -f -im ping.exe

:: wait for cywgin to execute pv, or let it hang
:: if it executes and terminates, %batdir%\fixed
:: is created

ping -n 6 localhost >nul
taskkill -f -im pv.exe

dir %batdir%\stage\fixed && exit /B 0 || exit /B 1
