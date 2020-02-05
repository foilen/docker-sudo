#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Create .deb ]----
DEB_FILE=docker-sudo_${VERSION}_amd64.deb
DEB_PATH=$RUN_PATH/build/debian_out/docker-sudo
rm -rf $DEB_PATH
mkdir -p $DEB_PATH $DEB_PATH/DEBIAN/ $DEB_PATH/usr/local/bin/

cat > $DEB_PATH/DEBIAN/control << _EOF
Package: docker-sudo
Version: $VERSION
Maintainer: Foilen
Architecture: amd64
Description: This is an application to let users use some parts of Docker, but only on their owned instances.
_EOF

cp -rv DEBIAN $DEB_PATH/
cp -rv build/bin/* $DEB_PATH/usr/local/bin/

cd $DEB_PATH/..
dpkg-deb --no-uniform-compression --build docker-sudo
mv docker-sudo.deb $DEB_FILE
