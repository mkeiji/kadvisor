#!/bin/bash

if [ $(pgrep dlv) ]; then
    echo
    echo "stop debugging"
    if [ $(pgrep make) ]; then
        kill -INT -$(pgrep make)
    fi
    echo
else
    echo
    echo "dlv process not found..."
    echo
fi
