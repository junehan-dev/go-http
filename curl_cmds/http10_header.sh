#!/bin/sh

curl --http1.0 -H "X-Test: Hello" http://$1 --verbose
