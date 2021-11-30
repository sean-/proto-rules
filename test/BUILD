subinclude("//build_defs:proto")

# Tests on the proto rules.


proto_library(
    name = "js_test_proto",
    srcs = ["js_test.proto"],
    languages = [
        "js",
    ],
)

filegroup(
    name = "js_test_proto_filegroup",
    srcs = [":js_test_proto"],
    requires = ["js"],
)

python_test(
    name = "js_out_test",
    srcs = ["js_out_test.py"],
    data = [
        ":js_test_proto_filegroup",
    ],
)

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
