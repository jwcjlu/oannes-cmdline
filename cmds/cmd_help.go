package cmds

import (
	"flag"
	"fmt"
)

type HelpCmd struct {
	baseCmd
}

func (c *HelpCmd) Run(args ...string) error {
	c.flagSet.Parse(args)
	verbose := c.flagSet.Lookup("v").Value.(flag.Getter).Get().(bool)
	var tip string
	if verbose {
		tip = c.DescLong()
	} else {
		tip = c.DescShort()
	}
	fmt.Println(tip)
	return nil
}
func (c *HelpCmd) UsageLine() string {
	return c.usageLine
}
func (c *HelpCmd) DescShort() string {
	return c.descShort
}
func (c *HelpCmd) DescLong() string {
	return c.descLong
}
func (c *HelpCmd) FlagSet() *flag.FlagSet {
	return c.flagSet
}
