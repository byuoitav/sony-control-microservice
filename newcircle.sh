#!/bin/bash

rm circle.yml
rm Dockerfile*
rm .dockerignore

cp ../kramer-microservice/circle.yml ./
cp ../kramer-microservice/docker* ./
cp ../kramer-microservice/.dockerignore ./
cp ../kramer-microservice/makefile ./
