#!/bin/sh

curl --http1.0 --basic --user june:pwd1234 http://$1 --verbose

