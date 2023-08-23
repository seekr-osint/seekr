#!/bin/sh

is_integer() {
    case $1 in
        ''|*[!0-9]*) return 1 ;;
        *) return 0 ;;
    esac
}

get_next_patch_version() {
    current_version="$1"
    major=$(echo "$current_version" | cut -d. -f1)
    minor=$(echo "$current_version" | cut -d. -f2)
    patch=$(echo "$current_version" | cut -d. -f3)
    
    if is_integer "$major" && is_integer "$minor" && is_integer "$patch"; then
        next_patch=$((patch + 1))
        echo "$major.$minor.$next_patch"
    else
        echo "Invalid version format"
    fi
}


latest_tag="v$(git describe --tags --abbrev=0)"
release="$latest_tag"
if [ "$(printf %s "$release" | cut -c 1)" = "v" ]; then
    version="${release#v}"

    major=$(echo "$version" | cut -d. -f1)
    minor=$(echo "$version" | cut -d. -f2)
    patch=$(echo "$version" | cut -d. -f3)

    if is_integer "$major" && is_integer "$minor" && is_integer "$patch"; then
        echo "Current release: $release"
        
        next_patch_version=$(get_next_patch_version "$version")
        echo "Next patch version: v$next_patch_version"
    else
        echo "$release is not a valid semantic version number"
    fi
else
    echo "$release does not start with 'v'"
fi
echo "enter release number"
read -n new_release
cd api/services
./mock.sh # commit stuff
