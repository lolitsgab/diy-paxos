load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "diy-paxos/diypaxos/server",
    visibility = ["//visibility:public"],
    deps = [
        "//diypaxos/proto",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)
