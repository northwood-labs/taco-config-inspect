
# Module `testdata/legacy-block-labels`

Core Version Constraints:

* `>= 0.11.0`

Provider Requirements:

* **aws:** `1.0.0`
* **external:** (any version)
* **notnull:** (any version)
* **noversion:** (any version)

## Input variables

* `foo` (default `"foo default"`): foo description

## Output values

* `foo`: foo description

## Managed resources

* `null_resource.foo` from `notnull`

## Data resources

* `data.external.foo` from `external`

## Child modules

* `foo` from `foo/bar/baz` (`1.2.3`)
