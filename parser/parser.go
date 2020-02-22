package parser

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

func parseProtoFile(fname string, protodirs ...string) ([]*desc.FileDescriptor, error) {
	parser := protoparse.Parser{
		ImportPaths:           protodirs,
		IncludeSourceCodeInfo: true,
	}
	return parser.ParseFiles(fname)
}
