#!/bin/sh

nohup $PWD/bin/crawler >logs/crawler.out 2>&1 &

sleep 1

ps -ef | grep $PWD
