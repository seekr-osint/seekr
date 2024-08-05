up() {
  git fetch
  git checkout "$1"
  go mod tidy
  ./update-hash.sh
  git add flake.nix go.mod go.sum
  git commit -m "updated hash"
  git push
  gh pr checks --watch
  gh pr review --approve
  gh pr merge --squash
}