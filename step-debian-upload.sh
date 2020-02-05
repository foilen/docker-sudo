#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Upload to Bintray ]----
DEB_FILE=docker-sudo_${VERSION}_amd64.deb
DEB_PATH=$RUN_PATH/build/debian_out/docker-sudo
curl -T $DEB_PATH/../$DEB_FILE -u$BINTRAY_USER:$BINTRAY_KEY "https://api.bintray.com/content/foilen/debian/docker-sudo/$VERSION/$DEB_FILE;deb_distribution=stable;deb_component=main;deb_architecture=amd64;publish=1"
