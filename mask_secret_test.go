package handy_test

import (
	"slices"
	"testing"
	"unicode"

	"github.com/miguelpragier/handy/v2"
)

func TestMask(t *testing.T) {
	for _, tt := range []struct {
		name, pattern, input string
		valid                bool
	}{
		{"uppercase", "##@0", "12A0", true},
		{"lowercase", "##@0", "12z0", true},
		{"wrong order", "##@0", "1A20", false},
		{"literal zero", "##@0", "12A1", false},
		{"short", "##@0", "12A", false},
		{"long", "##@0", "12A00", false},
		{"unicode digit", "##@0", "١2A0", false},
		{"unicode letter", "##@0", "12É0", false},
		{"unicode literal", "é.#", "é.7", true},
		{"dot literal", "é.#", "éX7", false},
		{"escapes", "\\#\\@\\\\#", "#@\\7", true},
		{"escaped placeholder is literal", "\\#", "7", false},
		{"empty", "", "", true},
		{"empty mask with input", "", "A", false},
		{"replacement rune literal", "�", "�", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := handy.New(nil, handy.Mask(tt.pattern)).Check(tt.input)
			if result.Valid() != tt.valid || (!tt.valid && !slices.Equal(result.Violations, []handy.Violation{{Code: "mask"}})) {
				t.Fatalf("unexpected result: %#v", result)
			}
		})
	}
	if got := handy.Mask("�")("\xff"); !slices.Equal(got, []handy.Violation{{Code: "mask"}}) {
		t.Fatalf("direct rule accepted invalid UTF-8: %v", got)
	}
}

func TestCompileMaskErrors(t *testing.T) {
	for _, pattern := range []string{"#", "\xff", "\\", "\\x"} {
		rule, err := handy.CompileMask(pattern)
		if pattern == "#" {
			if err != nil || rule == nil || len(rule("5")) != 0 {
				t.Fatalf("CompileMask(%q) failed: %v", pattern, err)
			}
		} else if err == nil || rule != nil {
			t.Errorf("CompileMask(%q) unexpectedly succeeded", pattern)
		}
	}
	assertPanics(t, func() { handy.Mask("\\") })
}

func TestCheckSecret(t *testing.T) {
	blocked := []string{"Password12"}
	rules := []handy.Rule{handy.RuneLen(10, 20), handy.Blocklist(blocked), handy.MinRunes(2, unicode.IsDigit)}
	blocked[0] = "changed"
	for _, tt := range []struct {
		input string
		want  []handy.Violation
	}{
		{"Password12", []handy.Violation{{Code: "blocklist"}}},
		{"password12", nil},
		{" Password12 ", nil},
		{"Passphrase١٢", nil},
		{"A1", []handy.Violation{{Code: "rune_length"}, {Code: "min_runes"}}},
	} {
		if got := handy.CheckSecret(tt.input, rules...); !slices.Equal(got, tt.want) {
			t.Errorf("CheckSecret(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
	got := handy.CheckSecret("\xff", func(string) []handy.Violation {
		t.Fatal("rule called with invalid UTF-8")
		return nil
	})
	if !slices.Equal(got, []handy.Violation{{Code: "invalid_utf8"}}) {
		t.Fatalf("invalid UTF-8: %v", got)
	}
	if got := handy.MinRunes(0, unicode.IsDigit)(""); len(got) != 0 {
		t.Errorf("zero minimum: %v", got)
	}
	assertPanics(t, func() { handy.MinRunes(-1, unicode.IsDigit) })
	assertPanics(t, func() { handy.MinRunes(1, nil) })
}

func TestConfirm(t *testing.T) {
	if got := handy.Confirm(" Café ", " Café "); len(got) != 0 {
		t.Errorf("equal values: %v", got)
	}
	for _, other := range []string{" café ", " Café", " Cafe\u0301 "} {
		if got := handy.Confirm(" Café ", other); !slices.Equal(got, []handy.Violation{{Code: "confirmation"}}) {
			t.Errorf("different values: %v", got)
		}
	}
}

func assertPanics(t *testing.T, run func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Error("expected panic")
		}
	}()
	run()
}
