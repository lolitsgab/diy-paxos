load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "utils",
    srcs = ["iputils.go"],
    importpath = "diy-paxos/diypaxos/utils",
    visibility = ["//visibility:public"],
    deps = [
        "//diypaxos/server",
        "@com_github_avast_retry_go//:retry-go",
    ],
)