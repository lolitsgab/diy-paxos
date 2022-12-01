load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "diypaxos_lib",
    srcs = ["main.go"],
    importpath = "diy-paxos/diypaxos",
    visibility = ["//visibility:private"],
    deps = [
        "//diypaxos/proto",
        "//diypaxos/server",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection",
    ],
)

go_binary(
    name = "diypaxos",
    embed = [":diypaxos_lib"],
    visibility = ["//visibility:public"],
)
