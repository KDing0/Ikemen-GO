@echo off
protoc --proto_path=..\src  --go_out=..\src ..\src\*.proto
echo.
pause