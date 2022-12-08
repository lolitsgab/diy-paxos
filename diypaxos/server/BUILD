load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = [
        "consensushandlers.go",
        "server.go",
    ],
    importpath = "diy-paxos/diypaxos/server",
    visibility = ["//visibility:public"],
    deps = [
        "//diypaxos/proto",
        "//diypaxos/storage",
        "@com_github_avast_retry_go//:retry-go",
    ],
)
