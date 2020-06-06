[![](https://img.shields.io/github/v/release/tanilolli/CloudlogTCI)](https://github.com/tanilolli/CloudlogTCI/releases)
![](https://img.shields.io/github/license/tanilolli/CloudlogTCI)

# CloudlogTCI
TCI bridge for [Cloudlog](https://github.com/magicbug/Cloudlog), an excellent web based loggin software for amateur radio.

TCI is a modern alternative to serial based rig control by Expert Electronics, based on WebSockets. 

Hopefully more manufactures adopt this excellent protocol that supports multiple clients over TCI/IP.

![Terminal](/screenshots/term.png)

## Instructions

Download latest release for Windows: [/releases](https://github.com/tanilolli/CloudlogTCI/releases)

Edit `config.yaml` with your Cloudlog server URL and R/W API key

## Build

Go get dependencies

> go get ./...

Build

> go build cloudlogtci.go

