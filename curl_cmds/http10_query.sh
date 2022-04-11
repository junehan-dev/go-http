#!/bin/sh

curl --http1.0 --get --data-urlencode "search word" http://$1 --verbose
