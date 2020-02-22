package cmds

import (
	"flag"
	"sync"
)

var (
	all  map[string]commander = make(map[string]commander)
	lock sync.RWMutex
)

func init() {
	ReginsterSubCmd("help", newHelpCmd())
	ReginsterSubCmd("create", newCreateCmd())
}
func ReginsterSubCmd(name string, command commander) {
	lock.Lock()
	all[name] = command
	lock.Unlock()
}
func RegisteredSubCmds() map[string]commander {
	return all
}

type commander interface {
	Run(agr ...string) error
	UsageLine() string
	DescShort() string
	DescLong() string
	FlagSet() *flag.FlagSet
}

func newHelpCmd() *HelpCmd {

	fs := flag.NewFlagSet("helpcmd", flag.ContinueOnError)
	fs.Bool("v", false, "verbose help info")

	u := baseCmd{
		usageLine: "oann help",
		descShort: `how to display help:oann help`,
		descLong: `oann <cmd> <options>: 
                   oann help:
                  -h display this help
                  -v display verbose info`,
		flagSet: fs,
	}

	return &HelpCmd{u}
}

type baseCmd struct {
	usageLine string
	descShort string
	descLong  string
	flagSet   *flag.FlagSet
}
