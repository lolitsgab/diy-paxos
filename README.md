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

Adapted from this [Medium article](https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d).

1. `cd ~/code/diy-paxos/diypaxos`

2. `eval $(minikube -p minikube docker-env)`

3. `bazel run //diypaxos:diypaxos_image -- --no-run` (this will error out but create a tagged image)

4. `cd diypaxos/k8; kubectl apply -f simple-kv-store.yaml`

5. `kubectl get pods,svc,ep`

6. see logs with `kubectl logs`

7. Tear down with `kubectl delete -f simple-kv-store.yaml`

```shell
kubectl exec -it kvstore-service-0 -- /bin/bash
```