#!/bin/sh
go build -o  tea_linux    /Users/kongdeyu/darwin_golang/cryptocoin/src/remind
cd /Users/kongdeyu/darwin_golang/cryptocoin/src/remind
nohup ./tea_linux > tea.file 2>&1 &
ps -ef | grep tea_linux


