package main

import "testing"

func TestSanitizeSpellEndpoint(t *testing.T) {
	tests := []struct {
		name  string
		spell string
		want  string
	}{
		{
			name:  "simple lowercase spell",
			spell: "light",
			want:  "light",
		},
		{
			name:  "spell with spaces",
			spell: "Light Healing",
			want:  "lighthealing",
		},
		{
			name:  "spell with apostrophe",
			spell: "Apprentice's Strike",
			want:  "apprenticesstrike",
		},
		{
			name:  "spell with mixed case and spaces",
			spell: "Great Fireball",
			want:  "greatfireball",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeSpellEndpoint(tt.spell); got != tt.want {
				t.Errorf("SanitizeSpellEndpoint(%q) = %q, want %q", tt.spell, got, tt.want)
			}
		})
	}
}

func TestSpaceMap(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "no spaces",
			str:  "nospaces",
			want: "nospaces",
		},
		{
			name: "spaces between words",
			str:  "two words",
			want: "twowords",
		},
		{
			name: "multiple spaces and tabs",
			str:  "a \t b  c",
			want: "abc",
		},
		{
			name: "empty string",
			str:  "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpaceMap(tt.str); got != tt.want {
				t.Errorf("SpaceMap(%q) = %q, want %q", tt.str, got, tt.want)
			}
		})
	}
}
