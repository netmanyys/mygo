#!/bin/bash
# Clean up previous container
containerID=$(docker ps | grep my-golang-server | awk '{print $1}');
docker stop $containerID;
docker rm $containerID

docker build -t my-golang-server .
docker run -d -p 9999:9999 my-golang-server