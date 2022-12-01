load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "simplekvstore_proto",
    srcs = ["simplekvstore.proto"],
    visibility = ["//visibility:public"],
    # deps=["@com_google_googleapis//:rpc_status_protos_src"],
    deps = ["@go_googleapis//google/rpc:status_proto"],
)

go_proto_library(
    name = "simplekvstore_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "diy-paxos/diypaxos/proto",
    proto = ":simplekvstore_proto",
    visibility = ["//visibility:public"],
    deps = ["@go_googleapis//google/rpc:status_go_proto"],
)

go_library(
    name = "proto",
    embed = [":simplekvstore_go_proto"],
    importpath = "diy-paxos/diypaxos/proto",
    visibility = ["//visibility:public"],
)
