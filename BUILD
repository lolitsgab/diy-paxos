load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix diy-paxos
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)