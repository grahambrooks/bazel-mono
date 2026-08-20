load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "c3e253237109ab2e2a8d3cb075688b98a6e6fce43d849849648e8e3a84f20d6f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.63.0/rules_go-v0.63.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.63.0/rules_go-v0.63.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "6549bd37cf1b82bac406119aef1b26bfec5d1c02d0d5a5518275e4513f47b3b2",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.52.2/bazel-gazelle-v0.52.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.52.2/bazel-gazelle-v0.52.2.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.18.3")

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "f6dcb97e992f13bc9effd794e9bb300f06b0dadc88061f81ae68d8d5994be964",
    strip_prefix = "rules_docker-0.26.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.26.0/rules_docker-v0.26.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
