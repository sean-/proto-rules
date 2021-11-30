subinclude("//build_defs:proto")

# Test that we can use a 'generated' source as well as a real one.
filegroup(
    name = "gen_proto",
    srcs = ["test.proto"],
)

proto_library(
    name = "test_proto",
    srcs = [":gen_proto"],
    languages = ["py"],
)

python_test(
    name = "specific_out_test",
    srcs = ["specific_out_test.py"],
    deps = [
        ":test_proto",
    ],
)
