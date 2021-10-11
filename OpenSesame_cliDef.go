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
	Port  string // listening port
	Path  string // path to serve files from
	Fixed int    // fixed web path # to serve files with
	Help  bool   // show usage help
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func initVars() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Port, "port", ":18888",
		"listening port")
	flag.StringVar(&Opts.Path, "path", "./",
		"path to serve files from")
	flag.IntVar(&Opts.Fixed, "fx", 0,
		"fixed web path # to serve files with")
	flag.BoolVar(&Opts.Help, "help", false,
		"show usage help")
}

func initVals() {
	exists := false
	// Now override those default values from environment variables
	if len(Opts.Port) == 0 ||
		len(os.Getenv("OPENSESAME_PORT")) != 0 {
		Opts.Port = os.Getenv("OPENSESAME_PORT")
	}
	if len(Opts.Path) == 0 ||
		len(os.Getenv("OPENSESAME_PATH")) != 0 {
		Opts.Path = os.Getenv("OPENSESAME_PATH")
	}
	if _, exists = os.LookupEnv("OPENSESAME_HELP"); Opts.Help || exists {
		Opts.Help = true
	}

}

const usageSummary = "  -port\tlistening port (OPENSESAME_PORT)\n  -path\tpath to serve files from (OPENSESAME_PATH)\n  -fx\tfixed web path # to serve files with (OPENSESAME_FX)\n  -help\tshow usage help (OPENSESAME_HELP)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags..]\n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, usageSummary)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nWill serve the files from the given path via web server\nof the given port using a one-time random path.\n\nExit and restart will serve from another random path.\n\nThe `-port` / `-path` can be overridden by environment variable(s)\n `OPENSESAME_PORT` / `OPENSESAME_PATH`\n")
	os.Exit(0)
}
