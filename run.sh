#!/bin/bash
docker rm bottoms-up
docker run -d --name bottoms-up --net=host bottoms-up
echo "Tailing logs.."
docker logs -f bottoms-up
