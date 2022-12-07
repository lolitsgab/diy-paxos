load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

k8s_deploy(
    name = "server",
    images = {
        "img": "//diypaxos:diypaxos_image",
    },
    resolver = "//diypaxos/k8/resolver:resolver",
    template = ":simple-kv-store.yaml",
)
