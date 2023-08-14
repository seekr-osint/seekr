#!/bin/sh
rm mock/*.json
tor() {
  firejail --net=tornet --noprofile go test
}
net() {
  go test
}
vpn() {
  proton start
  go test
  proton stop
  proton status
}
if [ "$1" == "" ]; then
  vpn
else
  "$1"
fi

