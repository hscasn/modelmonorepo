#!/usr/bin/env bash

ENV="${1}"
VERSION="${2}"

VALID_ENVIRONMENTS=("prod")

function not_in_list {
    local needle="${1}"
    local haystack="${2}"
    for hay in ${haystack[@]}; do
        if [[ "${hay}" == "${needle}" ]]; then
            return 1
        fi
    done
    return 0
}

function print_usage {
    echo "Usage:"
    echo "  deploy.sh <environment> <version>"
    echo "  example: deploy.sh prod latest"
}

if [[ "${#}" != 2 ]]; then
    (>&2 echo "Invalid number of arguments")
    print_usage
    exit 1
fi

if not_in_list "${ENV}" "${VALID_ENVIRONMENTS}"; then
    (>&2 echo "Invalid environment ${ENV}. Valid choices: ${VALID_ENVIRONMENTS}")
    print_usage
    exit 1
fi

if [[ ! -f "versionsets/${VERSION}.yaml" ]]; then
    (>&2 echo "Invalid version ${VERSION}. Check the versionsets directory - a valid version is the name of a file")
    print_usage
    exit 1
fi

(cd src; go run main.go "${ENV}" "${VERSION}")
