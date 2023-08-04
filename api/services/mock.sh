#!/bin/sh
rm mock/*.json
firejail --net=tornet --noprofile go test
