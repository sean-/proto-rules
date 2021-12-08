# Proto rules
Protobuf rules for the [Please](https://please.build) build system.

# Basic usage 
First add the base proto plugin to your project:
```python
# BUILD
plugin_repo(
    name = "proto",
    revision = "<Some git tag, commit, or other reference>",
)
```

Then add the proto language plugin of your choice. Please provides the following:
* [Go](https://github.com/please-build/go-proto-rules)
* [Python](https://github.com/please-build/python-proto-rules)
* [C/C++](https://github.com/please-build/cc-proto-rules)
* [Java](https://github.com/please-build/java-proto-rules)
* [Javascript](https://github.com/please-build/java-proto-rules)

Follow the setup instruction for each language you wish to generate proto code for.

If your language is not listed above, see the SDK section bellow.

## Generating code
There are two modes of operation for the built in protobuf rules. You can either generate just the 
protobuf messages, or all the messages, as well as any services defined.

The `grpc_library()` does the latter: generating all the messages and services. If can be used as
such:
```python
grpc_library(
    name = "service",
    srcs = ["service.proto"],
    # Optionally restrict to a subset of the configured languages. 
    # This defaults to all the languages you configured above. 
    languages = ["go"],  
)
```

If you just need the message types, use `proto_library()`:
```python
proto_library(
    name = "service",
    srcs = ["service.proto"],
    languages = ...,
)
```

You can then depend on these with your standard language rules:
```python
python_binary(
    name = "my_service",
    main = "my_service.py",
    deps = [":service"],
)
```

If you want to generate different code from a `.proto` file, see the SDK section below. 

## Downloading protoc automatically 
By default, this plugin expects `protoc` to be on the path. To download protoc
automatically with Please, add the following to `third_paty/proto/BUILD`:

```python
protoc_binary(
    name = "protoc",
    version = "<protoc version>",
    visibility = ["PUBLIC"],
)
```

Remember to change the protoc version to the desired version. The plugin can then be configured to 
use this instead of the path as such:
```
[Plugin "proto"]
ProtocTool = //third_party/proto:protoc
```

# Configuration
This plugin can be configured via the plugins section as follows: 
```
[Plugin "proto"]
SomeConfig = some-value
```

## ProtocTool (str, target)
As described above, this sets the protoc tool to use. Can be set to a `protoc_binary()` target, a tool on the path, 
or an absolute path to the tool.

```
[Plugin "proto"]
ProtocTool = //third_party/proto:protoc
```

## Definitions (repeatable target)
The build definitions for each proto or grpc language. 

```
[Plugin "proto"]
Definition = ///python-proto//build_defs:py
Definition = ///java-proto//build_defs:java
Definition = ///go-proto//build_defs:go
```

## ProtocFlag (repeatable str)
Any additional protoc flags to apply universally 

```
[Plugin "proto"]
ProtocFlag = --some-flag
ProtocFlag = --some-other-flag=value
```

# SDK
The proto rules can be extended in two different ways. Additional languages can be added to the `grpc_library()` 
and `proto_library()` rules, but entirely new types of rules can also be added. For example, perhaps you wish to 
generate a `proto_test()` that validates something about your `.proto` file, or you wish to geneerate a gRPC to 
json gateway. 

For either, there's an SDK which can be added to a `.build_defs` file as such:
```python
subinclude("///proto//build_defs/sdk:sdk")
```

## Adding new languages
To add a new language, you must define a function that returns a `proto_language()`, and expose this via `proto_build_defs()`. 
For an imaginary `foo` language, this might look like: 

```python
def foo_proto_language():
    return proto_language(
        language = "foo",
        build_def = foo_proto_library,
    )
```

The `build_def` parameter is the actual build definition that takes the `.proto` file, runs `protoc` on it and does 
what it needs to with the output. For example, it might compile the `.foo` generated files with a `foo_library()`. 
This build definition must have the following function signature:

```python
def(
    name:str, # The name of this rule
    parent:str, # The name of the parent rule (i.e. the proto_library(), or the grpc_library())
    srcs:list, # The .proto src files 
    deps:list=[], # Any deps of the rule
    visibility:list=None, # The visibility of the rule
    labels:list&features&tags=[], # Any additional labels to apply to this rule
    test_only:bool&testonly=False, # Whether this rule is test only 
    root_dir:str='', # The root director that proto import are relvative to. This should be apssed to protoc_rule() 
    protoc_flags:list=[], # Any additional proto flags to apply. This should just be passed to protoc_rule() 
    additional_context:dict=None # language specific context that was passed to the `proto_library()` or `grpc_library()`. 
                                 # Can be used by your rule for nefarious purposes. 
) -> None
```

For our imaginary language this might look like: 

```python
def foo_proto_library(name:str, language:str, srcs:list, deps:list=[], visibility:list=None, labels:list&features&tags=[],
                      test_only:bool&testonly=False, root_dir:str='', protoc_flags:list=[], additional_context:dict=None):
    # Read from the foo-proto plugin config. If you're not using the Plugin api, [BuildConfig] can be used instead. 
    deps += [CONFIG.FOO_PROTO.GO_DEP]

    # If we have any language specific stuff we want to apply to this rule, we can use the additional context
    if additional_context["PACKAGE"]:
        pkg = additional_context["PACKAGE"]
    else:
        pkg = "pkg"
    protoc = protoc_rule(
        name = name,
        srcs = srcs,
        language = "foo",
        tools= {"foo": [CONFIG.FOO_PROTO.PLUGIN]},
        protoc_flags = protoc_flags,
        plugin_flags = [
            '--foo_out="$OUT_DIR"',
            '--plugin=protoc-gen-go="$(which $TOOLS_FOO)"',
            f'--foo_opt=package={pkg}',
        ],
        labels = labels,
        test_only = test_only,
        root_dir = root_dir,
        deps = deps,
        visibility = visibility,
    )

    return foo_library(
        name = name,
        srcs = [protoc],
        deps = deps,
        test_only = test_only,
        labels = labels,
        visibility = visibility,
        package = pkg,
    )
```

Depending on how the foo protoc plugin works, we may need to define a similar `foo_grpc_library()`. We can then expose this
in our build file as such:

```python
proto_build_defs(
    name = "foo",
    srcs = ["foo.build_defs"],
    visibility = ["PUBLIC"],
    proto_languages = {
        # We're using the same definition for both gRPC and protobuf but these could well be defferent depending on how your 
        # languge's protoc plugin works. 
        # 
        # The keys in this dictionary are the proto language types that we provide from this rule. The values are the functions 
        # we defined above that return our proto_language() for each type.
        "grpc_language": ["foo_proto_language"], 
        "proto_language": ["foo_proto_language"],
    }
)
```

The final step is to configure this in the repo that will be using it:
```
[Plugin "proto"]
Definitions = ///foo-proto//build_defs:foo
```

The `proto_library()` and `grpc_library()` rules will then provide your `foo_library()` to any rules that depends on them 
with `requires = ["foo"]`. See [Require/Provides](https://please.build/require_provide.html) for more information on this 
mechanism. 

## Adding new proto target types
We currently support protobuf and gRPC plugins, however there are other things you might want to generate from `.proto` files. 
Luckily adding new types is relatively straight forward. For example, to add a `grpc_gateway` type, you can use the following 
build rule:

```python
protoc_plugins(
    name = "grpc_gateway_languages",
    build_defs = ["///foo-proto//build_defs:foo"],
    type = "grpc_language",
)
```

This will create a target that can be subincluded and exposes a function based on the rule name that returns all the configured proto
languages of that type. In this case `grpc_gateway_language()`.

That build target must expose a proto language for this type, for example:
```python
proto_build_defs(
    name = "foo",
    srcs = ["foo.build_defs"],
    visibility = ["PUBLIC"],
    proto_languages = {
        # The key here must match the type on the `protoc_plugins` rule above
        "grpc_gateway": ["foo_gateway_language"], 
        ...
    }
)
```

You can then subinclude this target, which will expose a `grpc_gateway_languages()` as described above: 
```python
subinclude("//build_defs:grpc_gateway_languages", "///proto//build_defs:proto")

def grpc_gateway_library(name:str, srcs:list, deps:list=None, visibility:list=None, l
                 labels:list&features&tags=[], test_only:bool&testonly=False, root_dir:str='', protoc_flags:list=None):
    """Defines a rule for a grpc library.

    Args:
      name (str): Name of the rule
      srcs (list): Input .proto files.
      deps (list): Dependencies (other grpc_library or proto_library rules)
      visibility (list): Visibility specification for the rule.
      languages (list | dict): List of languages to generate rules for, chosen from the set {cc, py, go, java, etc.}.
                               Alternatively, a dict mapping the language name to a definition of how to build it
                               (see proto_language for more details of the values).
      labels (list): List of labels to apply to this rule.
      test_only (bool): If True, this rule can only be used by test rules.
      root_dir (str): The directory that the protos are compiled relative to. Useful if your
                      proto files have import statements that are not relative to the repo root.
      protoc_flags (list): Additional flags to pass to protoc.
    """
    return proto_library(
        name = name,
        srcs = srcs,
        deps = deps,
        languages = merge_languages(languages, grpc_gateway_languages()),
        visibility = visibility,
        labels = labels,
        test_only = test_only,
        root_dir = root_dir,
        protoc_flags = protoc_flags,
    )
```