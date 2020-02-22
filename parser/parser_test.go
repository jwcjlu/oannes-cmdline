package main

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T){
	fd,err:=parseProtoFile("../example/helloworld.proto")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(fd)




}

