#!/bin/sh

curl --http1.0 --data-urlencode title="Head First PHP & MySQL" \
--data-urlencode author="Lynn Beighley, Michael Morrison" \
http://$1 --verbose
