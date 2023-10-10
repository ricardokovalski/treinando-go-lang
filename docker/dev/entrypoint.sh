#!/bin/sh

echo "valor debug [${DEBUG}]"
if [ "${DEBUG}" = "true" ]; then
    echo "dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient debug $@"
    dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient debug $@
else    
    echo "go run $@"
    go run $@
fi
