package main


import (
	"github.com/oannes-cmdline/cmds"
	"log"
	"os"
	"os/exec"
)

// 支持的所有subcmd，各个命令字需要主动注册
var subcmds = cmds.RegisteredSubCmds()

func main() {

	if _, err := exec.LookPath("protoc"); err != nil {
		log.Fatal("Please install protoc ... error:\n\t==> %v", err)
		return
	}

	if _, err := exec.LookPath("protoc-gen-go"); err != nil {
		log.Fatal("Please install protoc-gen-go ... error:\n\t==> %v", err)
		return
	}

	if l := len(os.Args); l == 1 {
		subcmds["help"].Run()
		return
	}

	c, ok := subcmds[os.Args[1]]
	if !ok || c == nil {
		subcmds["help"].Run()
		return
	}

	if err := c.Run(os.Args[2:]...); err != nil {
		log.Fatal("Run command:%v error:\n\t==> %v", os.Args, err)
	}
}