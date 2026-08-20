load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "685052b498b6ddfe562ca7a97736741d87916fe536623afb7da2824c0211c369",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.33.0/rules_go-v0.33.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "fec7a24acb027359eb7d2b31d116b17c81d2fab35842f6bfb2dd5945e1071346",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.53.0/bazel-gazelle-v0.53.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.53.0/bazel-gazelle-v0.53.0.tar.gz",
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
