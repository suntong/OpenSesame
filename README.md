# OpenSesame
Open Sesame -- One-time based file sharing

## Usage

```sh

$ OpenSesame -h
Usage of OpenSesame:
  -help
        show usage help
  -path string
        path to serve files from (default "./")
  -port string
        listening port (default ":18888")

$ OpenSesame -help

Usage:
 OpenSesame [flags..]

Flags:

  -port listening port (OPENSESAME_PORT)
  -path path to serve files from (OPENSESAME_PATH)
  -help show usage help

Details:

  -help
        show usage help
  -path string
        path to serve files from (default "./")
  -port string
        listening port (default ":18888")

Will serve the files from the given path via web server
of the given port using a one-time random path.

Exit and restart will serve from another random path.

The `-port` / `-path` can be overridden by environment variable(s)
 `OPENSESAME_PORT` / `OPENSESAME_PATH`

```
