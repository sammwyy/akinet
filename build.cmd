@echo off

:: Initial variables
set BIN_NAME=akinet

:: Compile for linux
set GOOS=linux
echo Compiling binary for linux...
go build -o ./out/akinet ./src/main.go

:: Compile for windows
set GOOS=windows
echo Compiling binary for windows...
go build -o ./out/akinet_win32.exe ./src/main.go

:: Show a message when the process end
echo Task completed.