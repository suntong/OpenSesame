
{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}
[![PoweredBy Go Easy Wireframe](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

# {{toc 5}}

# {{.Name}} - One-time based file sharing

The `{{.Name}}` is for sharing files temporarily. If you have some big files that you don't want to host anywhere else, you can spin up `{{.Name}}` to serve those shared files directly from your box using a random path/url, then close it down after the other party have grabbed them. The url is only good for one-time, and will become useless even if it leaked later.


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

## Author(s)

The `{{.Name}}` is brought you by

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

and its [contributor](https://github.com/suntong/{{.Name}}/graphs/contributors).

_Powered by_ [**Go Easy WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
