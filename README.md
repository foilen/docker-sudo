# About

This is an application to let users use some parts of Docker, but only on their owned instances.

# Local Usage

## Compile

`./create-local-release.sh`

The file is then in `build/bin/docker-sudo`

## Install

```
sudo chown root ./build/bin/docker-sudo
sudo chmod u+s ./build/bin/docker-sudo
```

## Configure

```
Files with the list of containers per user:
  /etc/docker-sudo/containers-{userName}.conf
Custom images folders:
  /etc/docker-sudo/images/{imageName}/Dockerfile

(Optional) Config files with some parameters
cat > /etc/docker-sudo/config.json << _EOF
{
	"network" : "myNetwork"
}
_EOF
```

## Execute

To see the help:
`./build/bin/docker-sudo`

To see ps:
`./build/bin/docker-sudo ps`

# Create release

`./create-public-release.sh`

That will show the latest created version. Then, you can choose one and execute:
`./create-public-release.sh X.X.X`

# Use with debian

Get the version you want from https://deploy.foilen.com/docker-sudo/ .

```bash
wget https://deploy.foilen.com/docker-sudo/docker-sudo_X.X.X_amd64.deb
sudo dpkg -i docker-sudo_X.X.X_amd64.deb
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
