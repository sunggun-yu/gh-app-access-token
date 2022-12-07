#!/bin/sh
set -e

if [ "$(uname)" = "Darwin" ] ; then
  OS="darwin"
else
  OS="linux"
fi

LOCAL_ARCH=$(uname -m)

case "${LOCAL_ARCH}" in
  x86_64|amd64)
    ARCH=amd64
    ;;
  armv8*|aarch64*|arm64)
    ARCH=arm64
    ;;
  *)
    echo "This system's architecture, ${LOCAL_ARCH}, isn't supported"
    exit 1
    ;;
esac

RELEASES_URL="https://github.com/sunggun-yu/gh-app-access-token/releases"
FILE_BASENAME="gh-app-access-token"

test -z "$VERSION" && VERSION="$(curl -sfL -o /dev/null -w %{url_effective} "$RELEASES_URL/latest" |
  grep -o 'releases/tag/v[0-9]*.[0-9]*.[0-9]*' |
  awk -F '/' '{ print $3}')"

test -z "$VERSION" && {
  echo "Unable to get gh-app-access-token version." >&2
  exit 1
}

test -z "$INSTALL_PATH" && INSTALL_PATH=/tmp

if ! test -d "$INSTALL_PATH" ; then
  echo "Creating INSTALL_PATH directory $INSTALL_PATH..." >&2
  mkdir -p $INSTALL_PATH
fi

TEMP_DIR=$(mktemp -d)
TAR_FILE="$TEMP_DIR/${FILE_BASENAME}.tar.gz"

if [ "$(uname -s)" = "Darwin" ] ; then
  DOWNLOAD_URL="$RELEASES_URL/download/$VERSION/${FILE_BASENAME}_${OS}_all.tar.gz"
else
  DOWNLOAD_URL="$RELEASES_URL/download/$VERSION/${FILE_BASENAME}_${OS}_${ARCH}.tar.gz"
fi

echo "Downloading gh-app-access-token $VERSION..."
curl -sfLo "$TAR_FILE" "$DOWNLOAD_URL"

if ! test -f "$TAR_FILE"; then
  echo "Unable to download gh-app-access-token." >&2
  exit 1
fi

tar -xzf "$TAR_FILE" -C "$TEMP_DIR" --strip-components=1

if ! test -f "$TEMP_DIR/gh-app-access-token"; then
  echo "Unable to extract gh-app-access-token." >&2
  exit 1
fi

mv $TEMP_DIR/gh-app-access-token $INSTALL_PATH
rm -r $TEMP_DIR

printf "gh-app-access-token has been successfully downloaded into the %s folder on your system.\n" "$INSTALL_PATH"
