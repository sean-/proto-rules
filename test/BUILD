# Tests on the legacy proto rules.

subinclude("///proto//build_defs:legacy")

grpc_library(
    name = "flags_proto",
    srcs = ["flags.proto"],
    languages = ["go"],
    protoc_flags = ["--include_source_info"],
)

go_test(
    name = "flags_test",
    srcs = ["flags_test.go"],
    deps = [":flags_proto"],
)
