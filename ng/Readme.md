# next-gen (ng) server 
## Description
ng-server is based on ubuntu:20.04. Its primary purpose is to serve incoming requests. It has auto-reconnect loop that tries to connect in every 40 seconds. Even at failure it will keep on retrying.

## Changelog
Reduced retry time to avoid race conditions - CON-900

## Commands
> cd ngserver
> chmod +x ngserver.sh
> ./ngserver,sh