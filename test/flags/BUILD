subinclude("//build_defs:proto")

grpc_library(
    name = "flags_proto",
    srcs = ["flags.proto"],
    languages = ["go"],
)

go_test(
    name = "flags_test",
    srcs = ["flags_test.go"],
    deps = [":flags_proto"],
)