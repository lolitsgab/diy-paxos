load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = [
        "consensushandlers.go",
        "interceptors.go",
        "server.go",
    ],
    importpath = "diy-paxos/diypaxos/server",
    visibility = ["//visibility:public"],
    deps = [
        "//diypaxos/proto",
        "//diypaxos/storage",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
    ],
)
