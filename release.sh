#!/bin/sh

release="$1"

git checkout -b "release-$release"
cd api/services
./mock.sh || exit 1
cd ../..
git add api/services/mock
git commit -m "update mock"
git push
gh pr create --title "Release $release"
