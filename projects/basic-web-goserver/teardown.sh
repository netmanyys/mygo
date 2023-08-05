#!/bin/bash

containerID=$(docker ps | grep my-golang-server | awk '{print $1}');
docker stop $containerID;
docker rm $containerID