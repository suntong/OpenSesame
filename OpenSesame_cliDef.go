// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "OpenSesame" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Port string // listening port
	Path string // path to serve files from
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Port, "port", ":18888",
		"listening port")
	flag.StringVar(&Opts.Path, "path", "./",
		"path to serve files from")

	// Now override those default values from environment variables
	if len(Opts.Port) == 0 ||
		len(os.Getenv("OPENSESAME_PORT")) != 0 {
		Opts.Port = os.Getenv("OPENSESAME_PORT")
	}
	if len(Opts.Path) == 0 ||
		len(os.Getenv("OPENSESAME_PATH")) != 0 {
		Opts.Path = os.Getenv("OPENSESAME_PATH")
	}

}

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags..]\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\n")
	os.Exit(0)
}
