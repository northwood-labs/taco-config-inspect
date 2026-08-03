# taco-config-inspect

> [!TIP]
> “TACOS” refers to _Terraform Automation and COllaboration Software_, and applies to the Terraform and OpenTofu ecosystem.

This repository contains a helper library for extracting high-level metadata about Terraform/OpenTofu modules from their source code. It processes only a subset of the information Terraform/OpenTofu itself would process, and in return it's able to be broadly compatible with modules written for many different versions of Terraform/OpenTofu.

## Current status

Hard-forked from HashiCorp. Currently refactoring, and merging improvements from other HashiCorp forks. The goal is to add support for modern Terraform and OpenTofu features while maintaining backwards compatibility, making it a drop-in replacement for [`terraform-config-inspect`](https://github.com/hashicorp/terraform-config-inspect).

Since this repository is licensed under MPL-2.0, we will continue to merge improvements from HashiCorp as well. Feel free to open an issue and/or poke us if there is an update we've missed.

The intention is for this fork to be _passively maintained_, meaning that we'll come back and make updates periodically on our own, unless someone opens a PR or files an issue. But it's not our focus. If a feature is important to you, implementations are welcome. If you have a passion for this tooling and want to get involved more deeply, reach out about becoming a collaborator.

## Usage as a library

```bash
go install github.com/northwood-labs/taco-config-inspect@latest
```

```go
import "github.com/northwood-labs/taco-config-inspect/tfconfig"

// ...

module, diags := tfconfig.LoadModule(dir)

// ...
```

Due to the [Terraform v1.0 Compatibility Promises](https://www.terraform.io/docs/language/v1-compatibility-promises.html), this library should be able to parse Terraform configurations written in the language as defined with Terraform v1.0, although it may not immediately expose _new_ additions to the language added during the v1.x series.

This library can also interpret valid Terraform configurations targeting Terraform v0.10 through v0.15, although the level of detail returned may be lower in older language versions.

## Command Line Tool

The primary way to use this repository is as a Go library, but as a convenience it also contains a CLI tool called `terraform-config-inspect`, installed automatically by the `go get` command above, that allows viewing module information in either a Markdown-like format or in JSON format. It also supports parsing Terraform stacks using the `--stack` flag.

```bash
taco-config-inspect path/to/module
```

For stacks:

```bash
taco-config-inspect --stack path/to/stack
```

```markdown
# Module `path/to/module`

Provider Requirements:

- **null:** (any version)

## Input Variables

- `a` (default `"a default"`)
- `b` (required): The b variable

## Output Values

- `a`
- `b`: I am B

## Managed Resources

- `null_resource.a` from `null`
- `null_resource.b` from `null`
```

```bash
taco-config-inspect --json path/to/module
```

```json
{
  "path": "path/to/module",
  "variables": {
    "A": {
      "name": "A",
      "default": "A default",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 1
      }
    },
    "B": {
      "name": "B",
      "description": "The B variable",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 5
      }
    }
  },
  "outputs": {
    "A": {
      "name": "A",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 9
      }
    },
    "B": {
      "name": "B",
      "description": "I am B",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 13
      }
    }
  },
  "required_providers": {
    "null": []
  },
  "managed_resources": {
    "null_resource.A": {
      "mode": "managed",
      "type": "null_resource",
      "name": "A",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 18
      }
    },
    "null_resource.B": {
      "mode": "managed",
      "type": "null_resource",
      "name": "B",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 19
      }
    }
  },
  "data_resources": {},
  "module_calls": {}
}
```

For stacks, you can use the `--stack` flag:

```bash
taco-config-inspect --stack path/to/stack
```

```markdown
# Terraform Stack: path/to/stack

## Variables

- **regions** (set(string))
- **identity_token** (string)
- **default_tags** (map(string)): A map of default tags to apply to all AWS resources

## Outputs

- **lambda_urls**: URLs to invoke lambda functions

## Components

- **s3** (source: `./s3`)
- **lambda** (source: `./lambda`)
- **api_gateway** (source: `./api-gateway`)

## Required Providers

- **aws** (source: `hashicorp/aws`)
- **archive** (source: `hashicorp/archive`)
```

```bash
taco-config-inspect --stack --json path/to/stack
```

```json
{
  "path": "path/to/stack",
  "variables": {
    "regions": {
      "name": "regions",
      "type": "set(string)",
      "default": null,
      "required": true,
      "pos": {
        "filename": "path/to/stack/variables.tfstack.hcl",
        "line": 4
      }
    }
  },
  "outputs": {
    "lambda_urls": {
      "name": "lambda_urls",
      "description": "URLs to invoke lambda functions",
      "pos": {
        "filename": "path/to/stack/outputs.tfcomponent.hcl",
        "line": 1
      },
      "type": "list(string)"
    }
  },
  "required_providers": {
    "aws": {
      "source": "hashicorp/aws"
    }
  },
  "components": {
    "s3": {
      "name": "s3",
      "source": "./s3",
      "pos": {
        "filename": "path/to/stack/components.tfcomponent.hcl",
        "line": 4
      }
    }
  }
}
```

## Contributing

This library and tool are intentionally focused on only extracting simple top-level metadata about a single Terraform/OpenTofu module. This is to reduce the maintenance burden of keeping this codebase synchronized with changes to Terraform/OpenTofu themselves: the features extracted by this package are unlikely to change significantly in future versions.

> [!IMPORTANT]
> External PRs for this codebase which add support for additional Terraform language features are **welcome**, provided they include test coverage and documentation updates.

If there is a feature you wish to see added, please open a GitHub issue first so we can discuss the feasability and design before submitting a pull request.

We would be happy to review PRs to fix bugs in existing functionality or to improve the usability of the Go package API, however. We will be hesitant about any breaking changes to the API, since this library is used by a number of existing tools and systems.

To work on this codebase you will need a recent version of Go installed. Please ensure all files match the formatting rules applied by `go fmt` and that all unit tests are passing.
