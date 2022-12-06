load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "storage",
    srcs = ["inmemorystorage.go"],
    importpath = "diy-paxos/diypaxos/storage",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
    ],
)

go_test(
    name = "storage_test",
    srcs = ["inmemorystorage_test.go"],
    embed = [":storage"],
)
