subinclude("//build_defs/sdk")

protoc_plugins(
    name = "proto_languages",
    build_defs = CONFIG.PROTO.DEFINITIONS,
    label_prefix = "proto_language:",
)

protoc_plugins(
    name = "grpc_languages",
    build_defs = CONFIG.PROTO.DEFINITIONS,
    label_prefix = "grpc_language:",
)

filegroup(
    name = "proto",
    srcs = ["proto.build_defs"],
    visibility = ["PUBLIC"],
)
