#!/bin/sh

curl --http1.0 -H "X-Test: Hello" -H "X-Test: World" http://$1 --verbose
