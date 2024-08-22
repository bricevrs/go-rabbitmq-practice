# RabbitMQ HA Cluster

This is a RabbitMQ HA cluster with 3 nodes. We use the official RabbitMQ docker image and the official RabbitMQ HA plugin 

- The cluster is created using docker-compose, and the nodes are connected using the HA plugin. Build_cluster.sh is a script that helps to connect the nodes together.
- The cluster is created using the `rabbitmq-cluster` network, and the nodes are named `rabbitmq1`, `rabbitmq2`, and `rabbitmq3`

## How to run the cluster

- 1. Start the cluster using the following command
```bash
make dev
```

- 2. Run the following to connect the nodes together
```bash
./build_cluster.sh
```

