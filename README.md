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

# Commands

```
ps : Show the status of all the containers you can see
bash <containerName> : Go into a container you can see
logs <containerName> : View the logs of a container you can see
restart <containerName> : Restart a container you own
start <containerName> : Start a container you own
stop <containerName> : Stop a container you own
tails <containerName> : View the logs of a container you can see and tail it
run <customImageName or hubImageName> : Start an image with your home directory mounted and start bash
exec <customImageName or hubImageName> <command> [arg1] [argN]: Start an image with your home directory mounted and execute the command
```
