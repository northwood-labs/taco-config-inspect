
# Module `testdata/basics`

Provider Requirements:

* **null:** (any version)

## Input variables

* `A` (default `"A default"`)
* `B` (required): The B variable
* `C` (required): The C variable

## Output values

* `A`
* `B`: I am B
* `C`: C is sensitive

## Managed resources

* `null_resource.A` from `null`
* `null_resource.B` from `null`
* `null_resource.C` from `null`
