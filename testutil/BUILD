load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "testutil",
    testonly = 1,
    srcs = ["testutil.go"],
    importpath = "example.com/owner/repo/testutil",
    visibility = ["//visibility:public"],
    deps = ["//lib"],
)
