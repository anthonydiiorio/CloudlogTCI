[![](https://img.shields.io/github/v/release/anthonydiiorio/CloudlogTCI)](https://github.com/anthonydiiorio/CloudlogTCI/releases)
![](https://img.shields.io/github/license/anthonydiiorio/CloudlogTCI)

# CloudlogTCI 📻
TCI bridge for [Cloudlog](https://github.com/magicbug/Cloudlog) and [Wavelog](https://github.com/wavelog/wavelog).

TCI is a modern alternative to serial based rig control by [Expert Electronics](https://eesdr.com/en/), based on WebSockets. 

![Terminal](/screenshots/term.png)

## Instructions

1. Download the latest Windows release from: [Releases](https://github.com/tanilolli/CloudlogTCI/releases), or build from source
3. Create a Read/Write API key in Cloudlog: [Cloudlog/wiki/API](https://github.com/magicbug/Cloudlog/wiki/API)
4. Edit `config.yaml` with your Cloudlog server URL and Read/Write API key
5. Enable TCI in ExpertSDR

Both VFOs will be now available under the Station/Radio dropdown and the Hardware Interfaces page.

## Build from source

Install Go https://go.dev/

```bash
go build
```
