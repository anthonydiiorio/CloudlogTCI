[![](https://img.shields.io/github/v/release/anthonydiiorio/CloudlogTCI)](https://github.com/anthonydiiorio/CloudlogTCI/releases)
![](https://img.shields.io/github/license/anthonydiiorio/CloudlogTCI)

# CloudlogTCI
TCI bridge for [Cloudlog](https://github.com/magicbug/Cloudlog), an excellent web based loggin software for amateur radio.

TCI is a modern alternative to serial based rig control by [Expert Electronics](https://eesdr.com/en/), based on WebSockets. 

![Terminal](/screenshots/term.png)

## Instructions

Download latest release for Windows: [/releases](https://github.com/tanilolli/CloudlogTCI/releases)

Edit `config.yaml` with your Cloudlog server URL and R/W API key

Both VFOs will be available in Cloudlog under the Station/Radio dropdown.

## Build

> go build

Note: For Apple Silicon Macs you must build with go 1.6 or higher.

