#!/bin/bash

set -e

# Check params
if [ $# -ne 1 ]
	then
		echo Usage: $0 version;
    echo E.g: $0 0.1.0
		echo Version is MAJOR.MINOR.BUGFIX
		echo Latest versions:
		git tag | tail -n 5
		exit 1;
fi

# Set environment
export LANG="C.UTF-8"
export VERSION=$1

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Compile ]----
./gradlew build

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
cp -rv build/gopath/build/out/docker-sudo $DEB_PATH/usr/local/bin/

cd $DEB_PATH/..
dpkg-deb --no-uniform-compression --build docker-sudo
mv docker-sudo.deb $DEB_FILE

echo ----[ Upload to Bintray ]----
cd $RUN_PATH
curl -T $DEB_PATH/../$DEB_FILE -u$BINTRAY_USER:$BINTRAY_KEY "https://api.bintray.com/content/foilen/debian/docker-sudo/$VERSION/$DEB_FILE;deb_distribution=stable;deb_component=main;deb_architecture=amd64;publish=1"

echo ----[ Git Tag ]==----
git tag -a -m $VERSION $VERSION

echo ----[ Operation completed successfully ]==----

echo
echo You can send the tag: git push --tags
