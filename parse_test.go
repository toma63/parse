package parse

import ("testing"
	"regexp"
)

func TestMatchSyntax(t *testing.T) {
	template_patt := []string{`\+`, `PLACED`, `\(`, `-?\d+`, `-?\d+`, `\)`, `FW`}
	template := make([]*regexp.Regexp, len(template_patt))
	for i := 0 ; i < len(template_patt) ; i++ {
		template[i] = regexp.MustCompile(template_patt[i])
	}
	tokens := []string{"+", "PLACED", "(", "-1000", "1000", ")", "FW"}
	bogus_tokens := []string{"+", "PLACED", "(", "-1000", "1000", ")", "FO"}
	short := []string{"(", "-1000", "1000", ")", "FO"}

	res := MatchSyntax(tokens, template)
	if res != true {
		t.Errorf("MatchSyntax(%q, %q) should have matched but did not", tokens, template)
	}

	res = MatchSyntax(bogus_tokens, template)
	if res != false {
		t.Errorf("MatchSyntax(%q, %q) should have matched but did not", bogus_tokens, template)
	}

	res = MatchSyntax(short, template)
	if res != false {
		t.Errorf("MatchSyntax(%q, %q) should have failed on length mismatch", short, template)
	}
}

func TestReadLinesStripped(t *testing.T) {
	line_ch := make(chan string, 10)
	ReadLinesStripped("./testdata/rltest", "//", line_ch)
	lc := 0
	lines := []string{"asd", "d", "dsasda ",
		"dsa", "asd ad asdasdre33r", "sdf", "r89768w6er"}
	for line := range line_ch {
		if line != lines[lc] {
			t.Errorf("ReadLinesStripped: line mismatch %q did not matc %q", line, lines[lc])
		}
		lc++
	}
	if lc != len(lines) {
		t.Errorf("ReadLinesStripped: line count mismatch %q did not matc %q", lc, len(lines))
	}
}
