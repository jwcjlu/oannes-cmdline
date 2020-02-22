package main

import (
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/oannes-cmdline/tpl"
	"os"
	"text/template"
)

func parseProtoFile(fname string, protodirs ...string) ([]*desc.FileDescriptor, error) {
	parser := protoparse.Parser{
		ImportPaths:           protodirs,
		IncludeSourceCodeInfo: true,
	}
	return parser.ParseFiles(fname)
}
func main() {
	fd,err:=parseProtoFile("/Users/jinwei/gopath/src/github.com/oannes-cmdline/example/helloworld.proto")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(fd)
	t,err1:=template.New("/Users/jinwei/gopath/src/github.com/oannes-cmdline/install/asset_go/rpc").
		Funcs(tpl.FuncMap).ParseFiles("oann.go.tpl")
	if err1!=nil{
		fmt.Println(err1)
		return
	}
   t.Execute(os.Stdout,fd)

}