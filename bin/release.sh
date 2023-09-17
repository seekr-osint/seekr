#!/bin/sh

release="$1"
if [ "$release" == "" ];then
  echo "args"
  exit 1
fi

git checkout -b "release-$release"
cd api/services
./mock.sh || exit 1
cd ../..
git add api/services/mock
git commit -m "update mock"
git push --set-upstream origin "release-$release"

gh pr create --title "Release $release"
gh pr review --approve
