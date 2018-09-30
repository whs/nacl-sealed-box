load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/whs/nacl-sealed-box
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = [
        "nonce.go",
        "sealed_box.go",
    ],
    importpath = "github.com/whs/nacl-sealed-box",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_x_crypto//blake2b:go_default_library",
        "@org_golang_x_crypto//nacl/box:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["sealed_box_test.go"],
    embed = [":go_default_library"],
    deps = ["@org_golang_x_crypto//nacl/box:go_default_library"],
)
