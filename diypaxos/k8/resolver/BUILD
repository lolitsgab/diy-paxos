# gazelle:ignore
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["resolver.go"],
    importpath = "github.com/bazelbuild/rules_k8s/k8s/go/cmd/resolver",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_google_go_containerregistry//pkg/authn:go_default_library",
        "@com_github_google_go_containerregistry//pkg/name:go_default_library",
        "@com_github_google_go_containerregistry//pkg/v1:go_default_library",
        "@com_github_google_go_containerregistry//pkg/v1/daemon:go_default_library",
        "@com_github_google_go_containerregistry//pkg/v1/remote:go_default_library",
        "@in_gopkg_yaml_v2//:go_default_library",
        "@io_bazel_rules_docker//container/go/pkg/compat:go_default_library",
        "@io_bazel_rules_docker//container/go/pkg/utils:go_default_library",
    ],
)

go_binary(
    name = "resolver",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["resolver_test.go"],
    embed = [":go_default_library"],
    deps = ["@com_github_google_go_cmp//cmp:go_default_library"],
)