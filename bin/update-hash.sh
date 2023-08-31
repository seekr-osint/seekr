#!/bin/sh
# should be run as git hook like this:
# ./update-hash.sh | xargs git add
FLAKE="./flake.nix"
old_line_with_hash="$(grep 'vendorSha256 = "sha256' "$FLAKE")"
old_hash="$(printf '%s' "$old_line_with_hash" | cut -d'"' -f2)"
new() {
  rm -rf ./vendor
  go mod vendor
  nix hash path ./vendor/
  rm -rf ./vendor
}
mod() {
  tidyf="$(cat < go.sum)"
  go mod tidy
  new_tidyf="$(cat < go.sum)"
  if [ "$tidyf" != "$new_tidyf" ]; then
    printf 'go.mod\ngo.sum\n'
  fi
}
mod || exit 1
new_hash="$(new)" || exit 1
if [ "$old_hash" == "" ]; then
  printf 'error: got empty old hash\n'
  exit 1
fi
if [ "$new_hash" == "" ]; then
  printf 'error: got empty new hash\n'
  exit 1
fi
if [ "$old_hash" != "$new_hash" ];then
  sed -i "s|$old_hash|$new_hash|g" flake.nix
  printf 'flake.nix\n'
fi
