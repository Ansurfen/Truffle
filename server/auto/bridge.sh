#!/bin/bash

sudo docker network create --driver bridge --subnet=10.2.36.0/16 --gateway=10.2.1.1 truffle