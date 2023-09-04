#!/bin/sh
rm mock/*.json
tor() {
  firejail --net=tornet --noprofile go test
}
net() {
  go test
}
vpn() {
  status="$(proton status)"
  echo $status
  proton start
  sleep 10
	if [ "$(proton status)" == "down" ]; then
    echo "FAILED TO START PROTON"
		exit 1
	fi
  go test
  if [ "$status" == "down" ]; then
    proton stop
  fi
  proton status
}
if [ "$1" == "" ]; then
  vpn
else
  "$1"
fi

