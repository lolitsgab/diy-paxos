# DIY Paxos Implementation

This repo will contain a simple implementation of Paxos. The underlying datastore for which we will attempt to keep in consensus will be a basic KV store.

## KvStore

The server exposes the KvStoreService. This service expose the following endpoints:

* Get
  * Get a value by key.
* Insert
  * Insert a key-value pair
* Remove
  * Remove by key
* Update
  * Update by key.
* Upsert
  * Update value if key already exists, insert value if it does not exist.
  
The key must be of type  `string`, and the value must be of type `int32`. I may extend the value to be a byte blob, but for now, looking to keep things simple/


## Running Docker container via K8s

We are deploying a Golang gRPC server using K8s. The base server is Ubuntu 16.0, but we can deploy a distroless container by removing the
`base` line from the deployment BUILD file. We are keeping it with a base of K8s just to facilitate debugging.

### Start/Stop/Update

#### Start

```shell
eval $(minikube docker-env) # only need to run once
bazel run //diypaxos/k8:deployment.apply

# or
eval $(minikube docker-env) # only need to run once
bazel run //diypaxos/k8:deployment.delete && bazel run //diypaxos/k8:deployment.apply
```

#### Stop

```shell
bazel run //diypaxos/k8:deployment.delete
```

#### Update

```shell
bazel run //diypaxos/k8:deployment.update
```

### View K8s Deployments

```shell
kubectl get pods,svc,ep
```

### View K8s Logs

#### For one host:

```shell
kubectl logs kvstore-service-0 
```

#### For all hosts:

```shell
kubectl logs -l app=kvstore --all-containers --ignore-errors
```

### Spawn Shell on Container

```shell
kubectl exec -it kvstore-service-0 -- /bin/bash
```

Replace `kvstore-service-0` with the replica you want to connect to.

### Resolve all hosts in a cluster

```shell
kubectl apply -f https://k8s.io/examples/admin/dns/dnsutils.yaml
kubectl exec -i -t dnsutils -- nslookup kubernetes.default
```

### Send requests using grpc_cli

You can send gRPC calls via the CLI using [grpc_cli](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).

```shell
grpc_cli call localhost:8080 SimpleKvStore.Get "key: 'hi'"
```

