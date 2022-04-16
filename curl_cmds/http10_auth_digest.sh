#!/bin/sh

curl --http1.0 -d message="hi" -d to="june" --digest -u june:pwd1234 http://127.0.0.1:$1/digest --verbose
