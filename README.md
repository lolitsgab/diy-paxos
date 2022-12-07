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

### Start

```shell
$ bazel run //diypaxos/k8:server.apply
```

### Stop

```shell
$ bazel run //diypaxos/k8:server.delete
```

### Update

```shell
$ bazel run //diypaxos/k8:server.update
```