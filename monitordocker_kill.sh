#!/bin/bash

# This script uses the logspout and http stream tools to let you watch the docker containers
# in action.
#
# More information at https://github.com/gliderlabs/logspout/tree/master/httpstream

if [ -z "$1" ]; then
   DOCKER_NETWORK=network_custom
else
   DOCKER_NETWORK="$1"
fi

if [ -z "$2" ]; then
   PORT=8000
else
   PORT="$2"
fi

echo Starting monitoring on all containers on the network ${DOCKER_NETWORK}

docker kill logspout 2> /dev/null 1>&2 || true
docker rm logspout 2> /dev/null 1>&2 || true