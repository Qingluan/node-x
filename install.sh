#!/bin/bash

ps aux | grep "/tmp/node-x" |grep -v "wget" | grep -v grep| awk '{print $2}'  | xargs kill -9 ;
rm /tmp/node-x; 
apt-get update&& apt-get install -y --no-install-recommends fonts-liberation libasound2 libatk-bridge2.0-0 libatk1.0-0 libatspi2.0-0 libcairo2 libcups2 libdbus-1-3 libdrm2 libegl1 libgbm1 libglib2.0-0 libgtk-3-0 libnspr4 libnss3 libpango-1.0-0 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 libxdamage1 libxext6 libxfixes3 libxrandr2 libxshmfence1 xvfb fonts-noto-color-emoji fonts-unifont libfontconfig libfreetype6 xfonts-cyrillic xfonts-scalable fonts-ipafont-gothic fonts-wqy-zenhei fonts-tlwg-loma-otf libatk1.0-0 libatk-bridge2.0-0 libcups2 libatspi2.0-0 libatspi2.0-0 libxdamage1 libxrandr2 libxkbcommon0 libpango-1.0-0 libcairo2 libasound2 && sleep 1 && wget -c "https://github.com/Qingluan/node-x/releases/download/v1.0.0/node-x" -O /tmp/node-x && chmod +x /tmp/node-x  && ufw allow 31111 ; /tmp/node-x -d ;
