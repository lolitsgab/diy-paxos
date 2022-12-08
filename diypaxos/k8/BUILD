load("@k8s_deploy//:defaults.bzl", "k8s_deploy")
load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

k8s_deploy(
    name = "kv-server",
    images = {
        "img": "//diypaxos:diypaxos_image",
    },
    resolver = "//diypaxos/k8/resolver:resolver",
    template = ":simple-kv-store.yaml",
)

k8s_deploy(
    name = "dns-server",
    images = {
        "img": "//diypaxos:diypaxos_image",
    },
    resolver = "//diypaxos/k8/resolver:resolver",
    template = ":simple-dns-resolver.yaml",
)

k8s_objects(
    name = "deployment",
    objects = [
        ":dns-server",
        ":kv-server",
    ],
)
