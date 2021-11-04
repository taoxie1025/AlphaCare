#!/bin/bash

DOCKER_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')
echo "$DOCKER_HOST   host.docker.internal" | sudo tee -a /etc/hosts