package handy

import (
	"reflect"
	"regexp"
	"testing"
)

func TestPolicy(t *testing.T) {
	p := Policy{
		Transforms: []StringTransform{TrimSpace, Lower},
		NormalizedRules: []Rule{Required, RuneLength(3, 24), StartsWithASCIIAlpha,
			AllowedRunes("abcdefghijklmnopqrstuvwxyz0123456789._")},
	}
	good := p.Check("  Ana.Silva  ")
	if good.Value != "ana.silva" || !good.Valid() {
		t.Fatalf("unexpected result: %#v", good)
	}
	bad := p.Check("  9 Ana!  ")
	if bad.Value != "9 ana!" || !reflect.DeepEqual(bad.Violations, []Violation{{"prefix"}, {"characters"}}) {
		t.Fatalf("unexpected violations: %#v", bad)
	}
	if result := p.Check(string([]byte{0xff})); result.Violations[0].Code != "invalid_utf8" {
		t.Fatalf("invalid UTF-8 accepted: %#v", result)
	}
}

func TestPolicyStagesAndOptional(t *testing.T) {
	p := Policy{OriginalRules: []Rule{Prefix(" ")}, Transforms: []StringTransform{TrimSpace}, Optional: true, NormalizedRules: []Rule{Required}}
	if result := p.Check(" "); !result.Valid() || result.Value != "" {
		t.Fatalf("unexpected optional result: %#v", result)
	}
	if result := p.Check("x"); result.Valid() || result.Violations[0].Code != "prefix" {
		t.Fatalf("original rule did not run: %#v", result)
	}
}

func TestStringRules(t *testing.T) {
	p := Policy{Transforms: []StringTransform{CompactSpace}, NormalizedRules: []Rule{RuneLength(3, 20), RunesMatching(LettersAndSpaces), Pattern(regexp.MustCompile("^\\pL"))}}
	if result := p.Check("  Éva   Silva "); !result.Valid() || result.Value != "Éva Silva" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if len(ByteLength(2, 2)("é")) != 0 || len(RuneLength(1, 1)("é")) != 0 {
		t.Fatal("length units are incorrect")
	}
}

func TestSeparators(t *testing.T) {
	rule := Separators("-_.")
	for _, test := range []struct {
		input string
		valid bool
	}{
		{"", true}, {"abc", true}, {"abc-123.x", true},
		{"-abc", false}, {"abc-", false}, {"abc._x", false},
		{"é-β", true},
	} {
		if got := len(rule(test.input)) == 0; got != test.valid {
			t.Errorf("Separators(%q) valid = %v", test.input, got)
		}
	}
}

func TestForbiddenRunesAndTransformUTF8(t *testing.T) {
	if len(ForbiddenRunes("!é")("abc")) != 0 || len(ForbiddenRunes("!é")("café")) != 1 {
		t.Fatal("forbidden characters were not checked")
	}
	p := Policy{Transforms: []StringTransform{func(string) string { return string([]byte{0xff}) }}, NormalizedRules: []Rule{Required}}
	result := p.Check("valid")
	if !reflect.DeepEqual(result.Violations, []Violation{{"invalid_utf8"}}) {
		t.Fatalf("unexpected result: %#v", result)
	}
}
