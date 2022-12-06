load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_library(
    name = "diypaxos_lib",
    srcs = ["main.go"],
    importpath = "diy-paxos/diypaxos",
    visibility = ["//visibility:private"],
    deps = [
        "//diypaxos/proto",
        "//diypaxos/server",
        "//diypaxos/storage",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection",
    ],
)

go_binary(
    name = "diypaxos",
    embed = [":diypaxos_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "diypaxos_image",
    # srcs = ["main.go"],
    embed = [":diypaxos_lib"],
    importpath = "diy-paxos/diypaxos",
)