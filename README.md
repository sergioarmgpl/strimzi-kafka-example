# STRIMZI KAFKA EXAMPLE

## Strimzi Installation
```
kubectl create namespace kafka
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
kubectl get pod -n kafka --watch
kubectl logs deployment/strimzi-cluster-operator -n kafka -f
```
## Kafka Cluster provisioning using Strimzi
```
kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n kafka 
kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka 
or
watch kubectl get pod -n kafka
```

## Build and deploy Consumer Daemon
```
cd src/consumer
/bin/bash build.sh <YOUR_DOCKER_USER>
kubectl apply -f golang-consumer.yaml
```

## Build and deploy Producer API
```
cd src/producer
/bin/bash build.sh <YOUR_DOCKER_USER>
kubectl apply -f golang-producer-api.yaml
```

# ENJOY IT!

# REFERENCES
- https://strimzi.io/quickstarts/