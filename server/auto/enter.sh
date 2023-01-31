#!/bin/bash

if [ $# -lt 1 ];
then
    echo "The argument is less than 1"
else
    sudo docker exec -it ${1} /bin/bash
fi