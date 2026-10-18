package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

type fakeGenerator struct{ calls []domain.LevelParams }

func (f *fakeGenerator) Generate(p domain.LevelParams) (*domain.TileMap, error) {
	f.calls = append(f.calls, p)
	return &domain.TileMap{Width: 2, Height: 2, Grid: [][]domain.TileType{{domain.Start, domain.Exit}, {domain.Ground, domain.Ground}}}, nil
}

func TestOneShotAndErrors(t *testing.T) {
	for _, tc := range []struct {
		args             []string
		code, calls      int
		wantOut, wantErr string
	}{
		{nil, 0, 1, "PE\n##\n", ""},
		{[]string{"-width", "10"}, 1, 0, "", "width"},
		{[]string{"-unknown"}, 2, 0, "", "flag"},
		{[]string{"-help"}, 0, 0, "", "Usage"},
		{[]string{"-interactive", "-height", "2"}, 1, 0, "", "height"},
	} {
		f := &fakeGenerator{}
		var out, errOut bytes.Buffer
		code := Run(tc.args, strings.NewReader(""), &out, &errOut, f)
		if code != tc.code || len(f.calls) != tc.calls || out.String() != tc.wantOut || !strings.Contains(errOut.String(), tc.wantErr) {
			t.Fatalf("args %v: code %d, calls %d, out %q, err %q", tc.args, code, len(f.calls), out.String(), errOut.String())
		}
	}
}

func TestInteractiveHistoryAndSettings(t *testing.T) {
	f := &fakeGenerator{}
	input := "1\n2\n1\n10\n1\n35\n6\n9\n1\n3\n1\n3\n2\n4\n"
	var out, errOut bytes.Buffer
	if code := Run([]string{"-interactive"}, strings.NewReader(input), &out, &errOut, f); code != 0 {
		t.Fatalf("exit %d: %s", code, errOut.String())
	}
	if len(f.calls) != 2 || f.calls[0].Width != 30 || f.calls[1].Width != 35 {
		t.Fatalf("settings not applied: %+v", f.calls)
	}
	for _, phrase := range []string{"LevelGen dev", "width must", "Choose 1, 2, 3, or 4", "1: 30x30", "2: 35x30", "+--+"} {
		if !strings.Contains(out.String(), phrase) {
			t.Fatalf("missing %q in %q", phrase, out.String())
		}
	}
}

func TestPreview(t *testing.T) {
	m := &domain.TileMap{Width: 2, Height: 2, Grid: [][]domain.TileType{{domain.Start, domain.Exit}, {domain.Ground, domain.Ground}}}
	got, err := Preview(m, domain.LevelParams{Seed: 7})
	if err != nil {
		t.Fatal(err)
	}
	want := "Map 2x2  seed 7\n+--+\n|PE|\n|##|\n+--+\nP Start  E Exit  # Ground  = Platform  ^ Spike\n"
	if got != want {
		t.Fatalf("preview %q, want %q", got, want)
	}
	if m.String() != "PE\n##\n" {
		t.Fatal("plain map changed")
	}
}
