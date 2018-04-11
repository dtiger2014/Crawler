#!/bin/sh

function StopServer()
{
	SERVERPIDS=`ps aux | grep $1 | grep $2  |grep -v "grep" |awk '{print $2}'`
	for pid in $SERVERPIDS; do
		kill $pid
		echo kill server ------ $1  pid $pid
	done 
	
	while true
	do
		CHECKPIDS=`ps aux | grep $1 | grep $2  |grep -v "grep" |awk '{print $2}'`
		if [[ -n $CHECKPIDS ]]
		then
			continue
		fi
	break
	done
	echo kill server ------ $1 ok 
}

StopServer crawler /home/vagrant/labs

