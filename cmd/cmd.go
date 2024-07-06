package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	XjreOption  string
	Class       string
	Args        []string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}
func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "Classpath", "", "Classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "Classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	Args := flag.Args()
	if len(Args) > 0 {
		cmd.Class = Args[0]
		cmd.Args = Args[1:]
	}
	return cmd
}
