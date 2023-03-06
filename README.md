
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/OpenSesame?status.svg)](http://godoc.org/github.com/suntong/OpenSesame)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/OpenSesame)](https://goreportcard.com/report/github.com/suntong/OpenSesame)
[![PoweredBy Go Easy Wireframe](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

# TOC
- [OpenSesame - One-time based file sharing](#opensesame---one-time-based-file-sharing)
  - [Usage](#usage)
  - [Author(s)](#author(s))

# OpenSesame - One-time based file sharing

The `OpenSesame` is for sharing files temporarily. If you have some big files that you don't want to host anywhere else, you can spin up `OpenSesame` to serve those shared files directly from your box using a random path/url, then close it down after the other party have grabbed them. The url is only good for one-time, and will become useless even if it leaked later.

From v1.0.8 and up, uploading is possible too.



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
- downloads are served from such path + 'd'
- upload is also possible from such path + 'u'

Exit and restart will serve from another random path.

The `-port` / `-path` can be overridden by environment variable(s)
 `OPENSESAME_PORT` / `OPENSESAME_PATH`

$ OpenSesame
2023/03/06 16:31:06 Serving at http://localhost:18888, with
                 download path /3531258666d (http://localhost:18888/3531258666d)
                 upload path /3531258666u

```

## Author(s)

The `OpenSesame` is brought you by

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

and its [contributor](https://github.com/suntong/OpenSesame/graphs/contributors).

_Powered by_ [**Go Easy WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
