subinclude("///go//build_defs:go")

go_toolchain(
    name = "toolchain",
    install_std = True,
    version = "1.22.3",
)

go_module(
    name = "protoc-gen-go",
    binary = True,
    download = ":protobuf_download",
    exported_deps = [":protobuf"],
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [":protobuf-go"],
)

go_module(
    name = "protoc-gen-go-grpc",
    binary = True,
    module = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    version = "v1.1.0",
    visibility = ["PUBLIC"],
    deps = [":protobuf-go"],
)

go_mod_download(
    name = "protobuf_download",
    licences = ["BSD-3-Clause"],
    module = "github.com/golang/protobuf",
    version = "v1.5.2",
)

go_module(
    name = "xterm",
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/term",
    version = "v0.0.0-20210615171337-6886f2dfbf5b",
    visibility = ["PUBLIC"],
    deps = [":xsys"],
)

go_module(
    name = "protobuf-go",
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
    deps = [
        ":cmp",
    ],
)

go_module(
    name = "protobuf",
    download = ":protobuf_download",
    exported_deps = [
        ":protobuf-go",
    ],
    visibility = ["PUBLIC"],
    install = ["..."],
    module = "github.com/golang/protobuf",
)


go_module(
    name = "cmp",
    install = [
        "cmp/...",
    ],
    visibility = ["PUBLIC"],
    module = "github.com/google/go-cmp",
    version = "b1c9c4891a6525d98001fea424c8926c6d77bb56",
)

go_module(
    name = "grpc",
    install = [
        "balancer/grpclb/...",
        "binarylog/...",
        "internal/status",
        "backoff",
        "encoding/gzip",
        "encoding/proto",
        "health/grpc_health_v1",
        "internal/resolver",
        "internal/resolver/dns",
        "balancer/roundrobin",
        "internal/buffer",
        "peer",
        "internal/grpcsync",
        "internal/grpcutil",
        "resolver",
        "internal/transport/networktype",
        "metadata",
        "attributes",
        "connectivity",
        "credentials/oauth",
        "grpclog",
        "health",
        "internal/grpcrand",
        "serviceconfig",
        "internal/grpclog",
        "internal/metadata",
        ".",
        "codes",
        "encoding",
        "internal/backoff",
        "internal/balancerload",
        "internal/channelz",
        "internal/syscall",
        "resolver/dns",
        "balancer/base",
        "internal",
        "internal/binarylog",
        "resolver/passthrough",
        "tap",
        "credentials",
        "internal/envconfig",
        "internal/transport",
        "stats",
        "status",
        "balancer",
        "internal/credentials",
        "internal/resolver/passthrough",
        "internal/resolver/unix",
        "internal/serviceconfig",
        "keepalive",
    ],
    licences = ["Apache-2.0"],
    module = "google.golang.org/grpc",
    version = "v1.40.0",
    visibility = ["PUBLIC"],
    deps = [
        ":genproto_rpc",
        ":net",
        ":oauth2",
        ":protobuf",
        ":protobuf-go",
        ":xsys",
    ],
)

go_module(
    name = "xsys",
    install = ["..."],
    module = "golang.org/x/sys",
    version = "v0.0.0-20220315194320-039c03cc5b86",
)

go_mod_download(
    name = "genproto_download",
    licences = ["Apache-2.0"],
    module = "google.golang.org/genproto",
    version = "v0.0.0-20210917145530-b395a37504d4",
)


go_module(
    name = "genproto_rpc",
    download = ":genproto_download",
    install = [
        "googleapis/rpc/...",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":protobuf-go",
    ],
)


go_module(
    name = "net",
    install = ["..."],
    module = "golang.org/x/net",
    version = "853a461950ffd0dac439017081f0f3bd9242bb17",
    deps = [
        ":text",
        ":xcrypto",
    ],
)

go_module(
    name = "text",
    install = [
        "encoding/...",
        "secure/...",
        "unicode/...",
        "transform",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/text",
    version = "v0.3.7",
    visibility = ["PUBLIC"],
)

go_module(
    name = "xcrypto",
    install = [
        "cast5",
        "curve25519",
        "curve25519/internal/field",
        "ed25519",
        "ssh/terminal",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20210920023735-84f357641f63",
    deps = [":xterm"],
)

go_module(
    name = "oauth2",
    install = [
        "authhandler",
        "google",
        "google/internal/externalaccount",
        "internal",
        "jws",
        "jwt",
        ".",
    ],
    licences = ["BSD-3-Clause"],
    module = "golang.org/x/oauth2",
    version = "v0.0.0-20210514164344-f6687ab2804c",
    visibility = ["PUBLIC"],
    deps = [
        ":cloud-metadata",
        ":net",
    ],
)

go_module(
    name = "cloud-metadata",
    install = ["compute/metadata"],
    licences = ["Apache-2.0"],
    module = "cloud.google.com/go",
    version = "v0.65.0",
    visibility = ["PUBLIC"],
)
