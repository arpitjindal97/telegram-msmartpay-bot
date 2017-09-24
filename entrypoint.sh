#!/bin/sh

export DISPLAY=:0
kill -9 $(ps aux | awk '/Xvfb/ { print $2}')

Xvfb :0 &
cd /arpit

./scrapper
