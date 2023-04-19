#!/bin/sh
# should be run as git hook like this:
# ./update-hash.sh | xargs git add
FLAKE="./flake.nix"
old_line_with_hash="$(grep 'vendorSha256 = "sha256' "$FLAKE")"
old_hash="$(printf '%s' "$old_line_with_hash" | cut -d'"' -f2)"
tidy="$(go mod tidy -v)"
new() {
  rm -rf ./vendor
  go mod vendor
  nix hash path ./vendor/
  rm -rf ./vendor
}
if [ "$tidy" != "\n" ]; then
  printf 'go.mod\ngo.sum\n'
fi
new_hash="$(new)"
if [ "$old_hash" != "$new_hash" ];then
  sed -i "s|$old_hash|$new_hash|g" flake.nix
  printf 'flake.nix\n'
fi
