#!/bin/ash

/bin/ip $*

if [ "$1" == "link" ] && [ "$5" == "up" ]; then
  sleep 0.5
fi
