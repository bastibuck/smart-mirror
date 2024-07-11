# Smart Mirror

## Setup, starting and stopping

### Initial setup

This repo includes a script to download and run the newest "release" (it's more like a **nightly** build of the `main` branch currently) automatically.

You would need

- 64bit system e.g. Raspberry Pi 4
- `ssh`, `git` and `docker` installed and setup
- a GitHub account with your ssh key connected

---

Than you simply

1. clone this repo
1. make `get-release.sh` executable (`chmod +x get-release.sh`)
1. run `get-release.sh`
1. smart-mirror container will be created and expose `http://[HOSTNAME/IP]`

### Updates

Should be as simple as running `get-release.sh` again. It will download a new version, remove old docker images and rebuild a new container from the new image.

### Some helpful commands (for me mostly 😂)

| Task               | Desc                                                                                          | Notes                   |
| ------------------ | --------------------------------------------------------------------------------------------- | ----------------------- |
| Start              | `sudo docker start smart-mirror`                                                              |                         |
| Stop               | `sudo docker stop smart-mirror`                                                               |                         |
| Follow logs        | `docker logs -f smart-mirror`                                                                 |
| Run app in browser | `DISPLAY=:0 chromium-browser http://localhost:80 --start-fullscreen --kiosk --no-first-run &` | Hit enter after opening |
| Stop browser       | `sudo killall chromium-browser`                                                               |                         |
