#!/usr/bin/env bash

for f in *.ejson; do
    ejson decrypt "${f}" > "${f%.*}.json"
done
