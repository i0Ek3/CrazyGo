#!/bin/bash

echo "Test all func: "
sleep 3
go test -v hello_test.go
echo " "
echo "Test all TestA* func: "
sleep 3
go test -v -run TestA hello_test.go
echo " "
echo "Test TestA func:"
sleep 3
go test -v -run TestA$ hello_test.go
