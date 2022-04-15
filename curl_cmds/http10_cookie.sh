#!/bin/sh

echo "-----------------FIRST: NO COOKIE"
curl --http1.0 --cookie-jar cookie.txt http://$1 --verbose
echo "\n---------------SECOND: YEP COOKIE"
curl --http1.0 --cookie cookie.txt http://$1 --verbose

