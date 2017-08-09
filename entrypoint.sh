#!/bin/sh

export DISPLAY=:0
kill -9 $(ps aux | awk '/Xvfb/ { print $2}')
kill -9 $(ps aux | awk '/java/ { print $2}')
Xvfb :0 &
cd /arpit
java -jar selenium-server-standalone.jar &
./scrapper
