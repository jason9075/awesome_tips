#!/bin/bash

INTERFACE="wlan0"

sudo tc qdisc del dev $INTERFACE root netem

echo "Network restore."


