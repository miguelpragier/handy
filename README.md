# Handy

Handy is a Go library for reusable string validation and normalization policies.
It requires Go 1.27.

## Installation

```sh
go get github.com/miguelpragier/handy/v2
```

## Username

```go
package main

import (
	"fmt"

	"github.com/miguelpragier/handy/v2"
)

func main() {
	username := handy.New(
		[]handy.Transform{handy.Trim, handy.Lower},
		handy.Required,
		handy.RuneLen(3, 24),
		handy.ASCIIAlphaStart,
		handy.Allow("abcdefghijklmnopqrstuvwxyz0123456789._"),
	)
	result := username.Check("  Ana.Silva  ")
	fmt.Println(result.Value, result.Valid()) // ana.silva true
}
```

New creates a policy with transformations and normalized rules. Use Policy directly
when original-input rules or optional values are needed. The earlier names remain
available for existing v2 callers.

Rules run in declaration order and return all violations. OriginalRules run
before transforms and NormalizedRules run afterwards. Set Optional to skip
normalized rules when the transformed value is empty. Original rules still run.

Violations contain stable codes and do not contain the input. Result.Original
and Result.Value contain the original and normalized strings, respectively.
Invalid UTF-8 stops processing with the code `invalid_utf8`. Applications
should decide how to handle and display results containing sensitive input.

## More policies

```go
displayName := handy.Policy{
	Transforms: []handy.Transform{handy.Compact},
	NormalizedRules: []handy.Rule{
		handy.Required,
		handy.RuneLen(2, 80),
		handy.MatchRunes(handy.LettersAndSpaces),
	},
}

reference := handy.Policy{
	Transforms: []handy.Transform{handy.Trim, handy.Upper},
	NormalizedRules: []handy.Rule{
		handy.Required,
		handy.ByteLen(3, 32),
		handy.Allow("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"),
	},
}

_, _ = displayName, reference
```

Custom rules are functions of type `func(string) []handy.Violation`.
Custom transforms are functions of type `func(string) string`. Policies should
be configured before concurrent use and their slices should not be modified
while checks are running.

## Migration from v1

Change the import path to `github.com/miguelpragier/handy/v2` and the Go
version to 1.27 or later. The v1 functions have been removed from v2. For validation, replace `CheckStr` flags with named rules,
and replace implicit transformations with ordered policy transforms. Length
rules explicitly count bytes or Unicode code points. Code points do not
necessarily correspond to user-perceived characters.

Run `go test ./...` to validate your integration.

## Character and separator rules

`Deny` rejects any listed character. `Separators("-_.")`
allows the listed separators between nonempty segments and rejects leading,
trailing, or adjacent separators. Combine these rules with `Required` when
an empty value is invalid.

## Development

Install [Task](https://taskfile.dev/) and run `task --list` to see the
available commands. Use `task check` for formatting, build, vet and tests, or
`task race` for the race detector.
