#!/bin/sh
docker run --rm --name mariadb -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mariadb:latest

