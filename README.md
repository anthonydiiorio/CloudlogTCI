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

![TCI](/screenshots/tci.png)

Both VFOs will be now available under the Station/Radio dropdown and the Hardware Interfaces page.

### Bandmap 🗺️

New! If you want to use the Bandmap feature in Wavelog set `bandmap: true` in the config file. To control RX2 from Wavelog, set the port to 54322 in the Hardware Interfaces settings. Click on the callsign in the bandmap to tune the radio.

![Hardware Interfaces](/screenshots/bandmap-settings.png)

If you have a port conflict, change the defaults for `rx1port` and `rx2port` in the config file. Then update the CAT URL in Wavelog.

## Build from source

Install Go https://go.dev/

```bash
go build
```
