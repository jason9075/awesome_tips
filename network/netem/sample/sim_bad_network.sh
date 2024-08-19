#!/bin/bash

DEFAULT_DELAY="100ms"
DEFAULT_LOSS="10%"

DELAY=${1:-$DEFAULT_DELAY}
LOSS=${2:-$DEFAULT_LOSS}

INTERFACE="wlan0"

echo "Bad network with: delay=$DELAY, loss=$LOSS."  

sudo tc qdisc add dev $INTERFACE root netem delay $DELAY loss $LOSS


