#!/usr/bin/env bash

CMD="${1}"

for t in ./*/; do
    (cd "${t}"
        eval "${CMD}")
done
