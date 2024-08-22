#!/bin/bash

docker exec go-rabbitmq-practice-rabbitmq2-1 rabbitmqctl stop_app
docker exec go-rabbitmq-practice-rabbitmq2-1 rabbitmqctl join_cluster rabbitmq1@rabbitmq1
docker exec go-rabbitmq-practice-rabbitmq2-1 rabbitmqctl start_app

docker exec go-rabbitmq-practice-rabbitmq3-1 rabbitmqctl stop_app
docker exec go-rabbitmq-practice-rabbitmq3-1 rabbitmqctl join_cluster rabbitmq1@rabbitmq1
docker exec go-rabbitmq-practice-rabbitmq3-1 rabbitmqctl start_app