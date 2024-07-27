#!/bin/bash

ps aux | grep "/tmp/node-x" |grep -v "wget" | grep -v grep| awk '{print $2}'  | xargs kill -9 ;
rm /tmp/node-x; 
wget -c "https://github.com/Qingluan/node-x/releases/download/v1.0.0/node-x" -O /tmp/node-x && chmod +x /tmp/node-x  && ufw allow 31111 ; /tmp/node-x -d ;
