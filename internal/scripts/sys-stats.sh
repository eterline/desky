#!/bin/sh

CPU=$(top -n 1 -b | awk '/^%Cpu/{print int($2)}')

print_stats () {
	echo -n '{'
    # cpu as "cpu": { "load": 40 } where load is in %
	printf '"cpu":{"load": %d }' $CPU
    echo -n ','
    # memory as "mem": { "current": 800, "total": 1024, "load", 82 } where amount is in MB and load in %
	free -m | awk 'NR==2{printf "\"mem\": { \"current\":%d, \"total\":%d, \"load\": %d }", $3,$2,$3*100/$2 }'
	echo -n ','
	# diska as "disk": { "current": 6, "total": 40, "used": 19 } where amount is in GB and used in %
	df -h | awk '$NF=="/"{printf "\"disk\": { \"current\":%d, \"total\":%d, \"used\": %d }", $3,$2,$5}'
	echo -n '}'
}

