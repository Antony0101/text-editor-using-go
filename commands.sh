#!/bin/bash

echo "Running... $1"

case $1 in
    build)
        go build -o bin/main ./cmd/main
        ;;
    test)
        echo "Testing..."
        ;;
    run)
        go run ./cmd/main
        ;;
    exec)
        ./bin/main
        ;;
    *)
        echo "Unknown command: $1"
        ;;
esac