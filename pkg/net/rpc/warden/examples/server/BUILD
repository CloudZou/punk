package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/CloudZou/punk/pkg/net/rpc/warden/examples/server",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/ecode:go_default_library",
        "//pkg/ecode/pb:go_default_library",
        "//pkg/log:go_default_library",
        "//pkg/net/rpc/warden:go_default_library",
        "//pkg/net/rpc/warden/proto/testproto:go_default_library",
        "//pkg/time:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
