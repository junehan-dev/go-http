#!/bin/sh

curl --http1.0 -F "file2<test.txt;type=text/html" -F "file1@test.txt;type=text/html" http://$1 --verbose

