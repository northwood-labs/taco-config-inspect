
# Module `testdata/overrides`

Provider Requirements:

* **null:** (any version)

## Input variables

* `A` (required): The A variable OVERRIDDEN
* `B` (required): The B variable
* `C` (required): An entirely new variable C

## Output values

* `A`: I am an overridden output!
* `B`: I am B

## Managed resources

* `null_resource.A` from `null`
* `null_resource.B` from `null`

## Child modules

* `foo` from `foo/bar/baz` (`1.0.2_override`)
