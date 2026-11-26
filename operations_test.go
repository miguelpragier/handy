package handy_test

import (
	"reflect"
	"regexp"
	"testing"
	"unicode"

	"github.com/miguelpragier/handy/v2"
)

func TestTransforms(t *testing.T) {
	tests := []struct {
		name        string
		fn          handy.StringTransform
		input, want string
	}{
		{"trim unicode space", handy.TrimSpace, "\u2003 Olá \t", "Olá"},
		{"trim empty", handy.TrimSpace, "", ""},
		{"compact spaces", handy.CompactSpace, " \tOlá\n\u2003mundo  ", "Olá mundo"},
		{"compact empty", handy.CompactSpace, "\t ", ""},
		{"lower", handy.Lower, "ÁBC Ε", "ábc ε"},
		{"upper", handy.Upper, "ábc ε", "ÁBC Ε"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(tt.input); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRules(t *testing.T) {
	tests := []struct {
		name        string
		rule        handy.Rule
		input, code string
	}{
		{"required empty", handy.Required, "", "required"},
		{"required spaces", handy.Required, " ", ""},
		{"byte lower", handy.ByteLength(2, 3), "a", "byte_length"},
		{"byte min", handy.ByteLength(2, 3), "é", ""},
		{"byte max", handy.ByteLength(2, 3), "abc", ""},
		{"byte above", handy.ByteLength(2, 3), "éé", "byte_length"},
		{"rune lower", handy.RuneLength(2, 3), "é", "rune_length"},
		{"rune min", handy.RuneLength(2, 3), "éé", ""},
		{"rune max", handy.RuneLength(2, 3), "ééé", ""},
		{"rune above", handy.RuneLength(2, 3), "éééé", "rune_length"},
		{"allowed empty", handy.AllowedRunes("éa"), "", ""},
		{"allowed unicode", handy.AllowedRunes("éa"), "aé", ""},
		{"allowed invalid", handy.AllowedRunes("éa"), "a!", "characters"},
		{"predicate accepted", handy.RunesMatching(unicode.IsLetter), "éβ", ""},
		{"predicate rejected", handy.RunesMatching(unicode.IsLetter), "é2", "characters"},
		{"ascii uppercase", handy.StartsWithASCIIAlpha, "Z9", ""},
		{"ascii lowercase", handy.StartsWithASCIIAlpha, "a", ""},
		{"ascii empty", handy.StartsWithASCIIAlpha, "", "prefix"},
		{"ascii unicode", handy.StartsWithASCIIAlpha, "É", "prefix"},
		{"ascii digit", handy.StartsWithASCIIAlpha, "1a", "prefix"},
		{"prefix success", handy.Prefix("pre"), "prefix", ""},
		{"prefix failure", handy.Prefix("pre"), "other", "prefix"},
		{"suffix success", handy.Suffix(".go"), "file.go", ""},
		{"suffix failure", handy.Suffix(".go"), "file.txt", "suffix"},
		{"pattern match", handy.Pattern(regexp.MustCompile("^[0-9]+$")), "123", ""},
		{"pattern mismatch", handy.Pattern(regexp.MustCompile("^[0-9]+$")), "123a", "pattern"},
		{"forbidden empty", handy.ForbiddenRunes("!"), "", ""},
		{"forbidden match", handy.ForbiddenRunes("!é"), "café", "forbidden_characters"},
		{"forbidden allowed", handy.ForbiddenRunes("!"), "café", ""},
		{"separators valid", handy.Separators("-_"), "a-b_c", ""},
		{"separators first", handy.Separators("-_"), "-abc", "separators"},
		{"separators last", handy.Separators("-_"), "abc_", "separators"},
		{"separators adjacent", handy.Separators("-_"), "a-_b", "separators"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rule(tt.input)
			var want []handy.Violation
			if tt.code != "" {
				want = []handy.Violation{{Code: tt.code}}
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %#v, want %#v", got, want)
			}
		})
	}
}

func TestLettersAndSpaces(t *testing.T) {
	for _, tt := range []struct {
		r    rune
		want bool
	}{
		{'é', true}, {'中', true}, {'\n', true}, {'5', false}, {'!', false}, {'\u0301', false},
	} {
		if got := handy.LettersAndSpaces(tt.r); got != tt.want {
			t.Errorf("LettersAndSpaces(%q) = %v, want %v", tt.r, got, tt.want)
		}
	}
}

func TestInvalidRuleConfiguration(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{"negative byte length", func() { handy.ByteLength(-1, 2) }},
		{"reversed byte length", func() { handy.ByteLength(3, 2) }},
		{"negative rune length", func() { handy.RuneLength(-1, 2) }},
		{"reversed rune length", func() { handy.RuneLength(3, 2) }},
		{"nil predicate", func() { handy.RunesMatching(nil) }},
		{"nil pattern", func() { handy.Pattern(nil) }},
		{"empty separators", func() { handy.Separators("") }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("expected panic")
				}
			}()
			tt.fn()
		})
	}
}

func TestPolicyOrderAndResults(t *testing.T) {
	var calls []string
	rule := func(name string) handy.Rule {
		return func(s string) []handy.Violation {
			calls = append(calls, name+":"+s)
			return []handy.Violation{{Code: name}}
		}
	}
	transform := func(name string) handy.StringTransform {
		return func(s string) string {
			calls = append(calls, name+":"+s)
			return s + name
		}
	}
	p := handy.Policy{
		OriginalRules:   []handy.Rule{rule("first"), rule("second")},
		Transforms:      []handy.StringTransform{transform("A"), transform("B")},
		NormalizedRules: []handy.Rule{rule("third"), rule("fourth")},
	}
	result := p.Check("x")
	wantCalls := []string{"first:x", "second:x", "A:x", "B:xA", "third:xAB", "fourth:xAB"}
	wantViolations := []handy.Violation{{Code: "first"}, {Code: "second"}, {Code: "third"}, {Code: "fourth"}}
	if !reflect.DeepEqual(calls, wantCalls) || !reflect.DeepEqual(result.Violations, wantViolations) || result.Original != "x" || result.Value != "xAB" || result.Valid() {
		t.Fatalf("calls = %v, result = %#v", calls, result)
	}
	if !(handy.Result{}).Valid() {
		t.Fatal("empty result should be valid")
	}
}

func TestPolicyUTF8AndOptional(t *testing.T) {
	called := false
	mark := func(string) []handy.Violation { called = true; return nil }
	p := handy.Policy{OriginalRules: []handy.Rule{mark}, Transforms: []handy.StringTransform{handy.TrimSpace}, NormalizedRules: []handy.Rule{mark}}
	invalid := string([]byte{0xff})
	r := p.Check(invalid)
	if called || r.Original != invalid || r.Value != invalid || !reflect.DeepEqual(r.Violations, []handy.Violation{{Code: "invalid_utf8"}}) {
		t.Fatalf("invalid input: called=%v result=%#v", called, r)
	}
	called = false
	p = handy.Policy{OriginalRules: []handy.Rule{handy.Prefix("x")}, Transforms: []handy.StringTransform{func(string) string { return invalid }}, NormalizedRules: []handy.Rule{mark}}
	r = p.Check("a")
	if called || !reflect.DeepEqual(r.Violations, []handy.Violation{{Code: "prefix"}, {Code: "invalid_utf8"}}) {
		t.Fatalf("invalid transform result: called=%v result=%#v", called, r)
	}
	p = handy.Policy{OriginalRules: []handy.Rule{handy.Prefix("x")}, Transforms: []handy.StringTransform{handy.TrimSpace}, Optional: true, NormalizedRules: []handy.Rule{handy.Required}}
	r = p.Check(" ")
	if r.Value != "" || !reflect.DeepEqual(r.Violations, []handy.Violation{{Code: "prefix"}}) {
		t.Fatalf("optional result = %#v", r)
	}
}

func TestShortAPI(t *testing.T) {
	p := handy.New(
		[]handy.Transform{handy.Trim, handy.Lower},
		handy.Required,
		handy.RuneLen(3, 24),
		handy.ASCIIAlphaStart,
		handy.Allow("abcdefghijklmnopqrstuvwxyz0123456789._"),
	)
	if got := p.Check("  Ana.Silva "); got.Value != "ana.silva" || !got.Valid() {
		t.Fatalf("unexpected valid result: %#v", got)
	}
	if got := p.Check("  9 Ana! "); !reflect.DeepEqual(got.Violations, []handy.Violation{{Code: "prefix"}, {Code: "characters"}}) {
		t.Fatalf("unexpected invalid result: %#v", got)
	}
	if got := handy.Compact(" a \t b "); got != "a b" {
		t.Errorf("Compact = %q", got)
	}
	for name, rule := range map[string]handy.Rule{
		"byte length": handy.ByteLen(1, 2),
		"forbidden":   handy.Deny("!"),
		"predicate":   handy.MatchRunes(unicode.IsLetter),
	} {
		t.Run(name, func(t *testing.T) {
			if len(rule("a")) != 0 || len(rule("!!!")) == 0 {
				t.Fatal("unexpected rule result")
			}
		})
	}
}
