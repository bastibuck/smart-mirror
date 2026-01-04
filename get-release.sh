#!/bin/bash

git fetch
if [ $? -ne 0 ]; then
    echo "Error: git fetch failed."
    exit 1
fi

git rebase
if [ $? -ne 0 ]; then
    echo "Error: git rebase failed."
    exit 1
fi

sudo make rebuild
if [ $? -ne 0 ]; then
    echo "Error: make rebuild failed."
    exit 1
fi
