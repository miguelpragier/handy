# Handy v2

[![Go Reference](https://pkg.go.dev/badge/github.com/miguelpragier/handy/v2.svg)](https://pkg.go.dev/github.com/miguelpragier/handy/v2)
[![Latest release](https://img.shields.io/github/v/release/miguelpragier/handy)](https://github.com/miguelpragier/handy/releases/latest)
[![Go 1.27+](https://img.shields.io/badge/go-1.27%2B-00ADD8?logo=go&logoColor=white)](https://github.com/miguelpragier/handy/blob/main/go.mod)
[![License](https://img.shields.io/github/license/miguelpragier/handy)](https://github.com/miguelpragier/handy/blob/main/LICENSE)

**Turn raw strings into values your application can use.**

Normalize usernames, validate fixed-format codes, and check secrets with
reusable rules. Handy requires Go 1.27 and uses only the standard library.

[Quick start](#quick-start) · [Code masks](#code-masks) · [Passwords and secrets](#passwords-and-secrets) · [More policies](#more-policies)

## Installation

```sh
go get github.com/miguelpragier/handy/v2
```

## Quick start

Trim a username, lowercase it, and return every rule violation in one pass.

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
	fmt.Println(username.Check("  9 Ana!  ").Violations) // [{prefix} {characters}]
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

## Code masks

Use `#` for an ASCII digit and `@` for an ASCII letter. Every other character
matches literally, including `0`. Handy checks the entire string.

```go
package main

import (
	"fmt"

	"github.com/miguelpragier/handy/v2"
)

func main() {
	code := handy.New(nil, handy.Mask("##@0"))
	fmt.Println(code.Check("12A0").Valid())   // true
	fmt.Println(code.Check("12A1").Violations) // [{mask}]

	ticket := handy.New(nil, handy.Mask(`\#-@@-##`))
	fmt.Println(ticket.Check("#-AB-27").Valid()) // true
}
```

Use `\#`, `\@` and `\\` for literal placeholder characters and
backslashes. Use `CompileMask(pattern)` for patterns loaded from configuration
so you can handle syntax errors. `Mask(pattern)` panics on invalid syntax.

## Passwords and secrets

`CheckSecret` applies rules to the exact value received, preserving spaces
and case. It returns only violations. Supply your application's length
limits and blocklist. Use `Confirm` to check a separate confirmation field.

```go
package main

import (
	"fmt"

	"github.com/miguelpragier/handy/v2"
)

func main() {
	rules := []handy.Rule{
		handy.RuneLen(15, 128),
		handy.Blocklist([]string{"passwordpassword"}),
	}

	password := "violet rivers under cloud"
	fmt.Println(len(handy.CheckSecret(password, rules...)) == 0) // true
	fmt.Println(handy.CheckSecret("passwordpassword", rules...)) // [{blocklist}]
	fmt.Println(handy.Confirm(password, "violet rivers"))         // [{confirmation}]
}
```

When an existing system requires character classes, add a rule such as
`handy.MinRunes(2, unicode.IsDigit)` using Go's `unicode` package. This rule
counts digits across Unicode. The blocklist compares exact values. Built-in
violations contain only codes. Custom rules should also avoid including
secrets in their violation codes.

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

## Character and separator rules

`Deny` rejects any listed character. `Separators("-_.")`
allows the listed separators between nonempty segments and rejects leading,
trailing, or adjacent separators. Combine these rules with `Required` when
an empty value is invalid.

## Development

Install [Task](https://taskfile.dev/) and run `task --list` to see the
available commands. Use `task check` for formatting, build, vet and tests, or
`task race` for the race detector.
