#!/bin/sh
FLAKE="./flake.nix"
old_line_with_hash="$(grep 'vendorSha256 = "sha256' "$FLAKE")"
old_hash="$(printf '%s' "$old_line_with_hash" | cut -d'"' -f2)"
new() {
  rm -rf ./vendor
  go mod vendor
  nix hash path ./vendor/
  rm -rf ./vendor
}
new_hash="$(new)"
printf 'old: %s\nnew: %s\n' "$old_hash" "$new_hash"
if [ "$old_hash" != "$new_hash" ];then
  sed -i "s|$old_hash|$new_hash|g" flake.nix
  printf 'Hash updated.\n'
fi
