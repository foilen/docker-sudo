# About

This is an application to let users use some parts of Docker, but only on their owned instances.

# Local Usage


## Compile

`./gradlew clean build` 

## Install

```
sudo chown root ./build/gopath/build/out/docker-sudo
sudo chmod u+s ./build/gopath/build/out/docker-sudo
```

## Configure

```
Files with the list of containers per user:
  /etc/docker-sudo/containers-{userName}.conf
Custom images folders:
  /etc/docker-sudo/images/{imageName}/Dockerfile
```

## Execute

To see the help:
`./build/gopath/build/out/docker-sudo`

To see ps:
`./build/gopath/build/out/docker-sudo ps`

# Create release

`./create-public-release.sh`

# Use with debian

```bash
echo "deb https://dl.bintray.com/foilen/debian stable main" | sudo tee /etc/apt/sources.list.d/foilen.list
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
sudo apt update
sudo apt install docker-sudo
```
