package cmds

import (
	"errors"
	"flag"
	"path/filepath"
)

type CreateCmd struct {
	baseCmd
	ProtoFile string
	ProtoDir  string
	Language  string
}

func newCreateCmd() *CreateCmd {

	fs := flag.NewFlagSet("helpcmd", flag.ContinueOnError)
	fs.Bool("v", false, "verbose help info")

	u := baseCmd{
		usageLine: "oann create",
		descShort: `how to create project：
      oann create -protofile=？
      oann create -protodir=？ -protofile=? -protocol=?
      oann create  -protofile=? -protocol=? -alias`,
		descLong: `how to create project：
      oann create -protofile=？
      oann create -protodir=？ -protofile=? -protocol=?
      oann create  -protofile=? -protocol=? -alias`,
		flagSet: fs,
	}
	return &CreateCmd{baseCmd: u}
}
func (c *CreateCmd) Run(args ...string) error {
	var protofile string
	if len(protofile) == 0 {
		return errors.New("no protofile specified")
	}
	p, err := filepath.Abs(protofile)
	if err != nil {
		panic(err)
	}
	c.ProtoFile = filepath.Base(p)

	return nil
}
func (c *CreateCmd) UsageLine() string {
	return c.usageLine
}
func (c *CreateCmd) DescShort() string {
	return c.descShort
}
func (c *CreateCmd) DescLong() string {
	return c.descLong
}
func (c *CreateCmd) FlagSet() *flag.FlagSet {
	return c.flagSet
}
