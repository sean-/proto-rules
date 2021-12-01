subinclude("@python//build_defs:python")

python_wheel(
    name = "protobuf",
    outs = ["google"],
    version = "3.12.2",
    deps = [":six"],
    visibility = ["PUBLIC"],
)

python_wheel(
    name = "six",
    outs = ["six.py"],
    hashes = ["43a5af1176750c6100480a370863422642afdad3f2f3191298af951c4f4f6080"],
    version = "1.14.0",
)
