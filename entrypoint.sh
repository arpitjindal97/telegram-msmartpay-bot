#!/bin/sh

export DISPLAY=:0
Xvfb :0 &
cd /arpit
java -jar selenium-server-standalone-*.jar &
./scrapper
