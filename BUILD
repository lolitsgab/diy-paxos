load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix diy-paxos
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-prune",
    ],
    command = "update-repos",
)
