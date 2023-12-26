// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "OpenSesame" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Host        string // host name preferred for external access
	Port        string // listening port
	Path        string // path to serve files from
	Fixed       int    // fixed web path # to serve files with
	MediaGalley bool   // as media galley backend server
	Help        bool   // show usage help
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func initVars() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Host, "host", "localhost",
		"host name preferred for external access")
	flag.StringVar(&Opts.Port, "port", ":18888",
		"listening port")
	flag.StringVar(&Opts.Path, "path", "./",
		"path to serve files from")
	flag.IntVar(&Opts.Fixed, "fx", 0,
		"fixed web path # to serve files with")
	flag.BoolVar(&Opts.MediaGalley, "mg", false,
		"as media galley backend server")
	flag.BoolVar(&Opts.Help, "help", false,
		"show usage help")
}

func initVals() {
	exists := false
	// Now override those default values from environment variables
	if len(Opts.Host) == 0 &&
		len(os.Getenv("OPENSESAME_HOST")) != 0 {
		Opts.Host = os.Getenv("OPENSESAME_HOST")
	}
	if len(Opts.Port) == 0 &&
		len(os.Getenv("OPENSESAME_PORT")) != 0 {
		Opts.Port = os.Getenv("OPENSESAME_PORT")
	}
	if len(Opts.Path) == 0 &&
		len(os.Getenv("OPENSESAME_PATH")) != 0 {
		Opts.Path = os.Getenv("OPENSESAME_PATH")
	}
	if Opts.Fixed == 0 &&
		len(os.Getenv("OPENSESAME_FX")) != 0 {
		if i, err := strconv.Atoi(os.Getenv("OPENSESAME_FX")); err == nil {
			Opts.Fixed = i
		}
	}
	if _, exists = os.LookupEnv("OPENSESAME_MG"); Opts.MediaGalley || exists {
		Opts.MediaGalley = true
	}
	if _, exists = os.LookupEnv("OPENSESAME_HELP"); Opts.Help || exists {
		Opts.Help = true
	}

}

const usageSummary = "  -host\thost name preferred for external access (OPENSESAME_HOST)\n  -port\tlistening port (OPENSESAME_PORT)\n  -path\tpath to serve files from (OPENSESAME_PATH)\n  -fx\tfixed web path # to serve files with (OPENSESAME_FX)\n  -mg\tas media galley backend server (OPENSESAME_MG)\n  -help\tshow usage help (OPENSESAME_HELP)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags..]\n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, usageSummary)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		`
Will serve the files from the given path via web server
of the given port using a one-time random path.
- downloads are served from such path + 'd'
- upload is also possible from such path + 'u'

Exit and restart will serve from another random path.

The '-host', '-port' / '-path' can be overridden by environment variable(s)
'OPENSESAME_HOST', 'OPENSESAME_PORT' / 'OPENSESAME_PATH'
`)
	os.Exit(0)
}
