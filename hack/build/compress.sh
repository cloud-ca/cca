#!/usr/bin/env bash

set -e

NAME=$1
VERSION=$2

if [ -z "${NAME}" ]; then
    echo "Error: NAME is missing. e.g. ./compress.sh <name> <version>"
    exit 1
fi

if [ -z "${VERSION}" ]; then
    echo "Error: VERSION is missing. e.g. ./compress.sh <name> <version>"
    exit 1
fi

PWD=$(cd $(dirname "$0") && pwd -P)
BUILD_DIR="${PWD}/../../bin"

printf "\033[36m==> Compress binary\033[0m\n"

for platform in $(find ${BUILD_DIR} -mindepth 1 -maxdepth 1 -type d); do
    OSARCH=$(basename ${platform})
    FULLNAME="${NAME}-${VERSION}-${OSARCH}"

    case "${OSARCH}" in
    "windows"*)
        if ! command -v zip >/dev/null; then
            echo "Error: cannot compress, 'zip' not found"
            exit 1
        fi

        zip -q -j ${BUILD_DIR}/${FULLNAME}.zip ${platform}/${NAME}.exe
        printf -- "--> %15s: bin/%s\n" "${OSARCH}" "${FULLNAME}.zip"

        ;;
    *)
        if ! command -v tar >/dev/null; then
            echo "Error: cannot compress, 'tar' not found"
            exit 1
        fi

        tar -czf ${BUILD_DIR}/${FULLNAME}.tar.gz --directory ${platform}/ ${NAME}
        printf -- "--> %15s: bin/%s\n" "${OSARCH}" "${FULLNAME}.tar.gz"

        ;;
    esac
done

cd ${BUILD_DIR}
touch ${NAME}-${VERSION}.sha256sum

for binary in $(find . -mindepth 1 -maxdepth 1 -type f | grep -v "${NAME}-${VERSION}.sha256sum" | sort); do
    binary=$(basename ${binary})

    if command -v sha256sum >/dev/null; then
        sha256sum ${binary} >>${NAME}-${VERSION}.sha256sum
    elif command -v shasum >/dev/null; then
        shasum -a256 ${binary} >>${NAME}-${VERSION}.sha256sum
    fi
done

cd - >/dev/null 2>&1
printf -- "\n--> %15s: bin/%s\n" "sha256sum" "${NAME}-${VERSION}.sha256sum"
