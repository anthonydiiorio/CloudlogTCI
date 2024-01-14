[![](https://img.shields.io/github/v/release/anthonydiiorio/CloudlogTCI)](https://github.com/anthonydiiorio/CloudlogTCI/releases)
![](https://img.shields.io/github/license/anthonydiiorio/CloudlogTCI)

# CloudlogTCI 📻
TCI bridge for [Cloudlog](https://github.com/magicbug/Cloudlog), an excellent web based logging software for amateur radio.

TCI is a modern alternative to serial based rig control by [Expert Electronics](https://eesdr.com/en/), based on WebSockets. 

CloudlogTCI allows you to connect your SunSDR to Cloudlog with the TCI protocol.

![Terminal](/screenshots/term.png)

## Instructions

1. Download latest release from: [Releases](https://github.com/tanilolli/CloudlogTCI/releases)
2. Create a Read/Write API key in Cloudlog: [Cloudlog/wiki/API](https://github.com/magicbug/Cloudlog/wiki/API)
3. Edit `config.yaml` with your Cloudlog server URL and Read/Write API key
4. Enable TCI in ExpertSDR

Both VFOs will be now available in Cloudlog under the Station/Radio dropdown.

## Build

```bash
go build
```
Note: For Apple Silicon Macs you must build with Go 1.6 or higher.

## Mac Users

**The recommended way to to install on macOS is to build from source.**
