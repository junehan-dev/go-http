#!/bin/sh

curl --http1.0 --compressed -H "Accept-Encoding: gzip, br" http://$1 --verbose
