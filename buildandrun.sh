#!/bin/sh
cd cmd/web
go build -o ../build/thinkingincodeapp
cd ../
./build/thinkingincodeapp
$SHELL 