#! /bin/bash

sudo ip addr add 10.1.0.10/24 dev tap0
sudo ip link set dev tap0 up

arping -I tap0 10.1.0.10

