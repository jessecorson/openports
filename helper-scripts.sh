#!/usr/bin/env bash

if [ $1 == "kill" ] ; then
  if [[ $2 =~ ^[0-9]*$ ]] && [[ $2 != '' ]]; then
    echo "Killing process for TCP $2"
    kill $(netstat -vanp tcp | grep $2 | grep LISTEN | awk '{print$9}')
  else
    kill $(netstat -vanp tcp | grep 8080 | grep LISTEN | awk '{print$9}')
  fi
fi
