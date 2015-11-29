package parse

import "testing"

func TestMatchSyntax(t *testing.T) {
	template := []string{`\+`, `PLACED`, `\(`, `-?\d+`, `-?\d+`, `\)`, `FW`}
	tokens := []string{"+", "PLACED", "(", "-1000", "1000", ")", "FW"}
	bogus_tokens := []string{"+", "PLACED", "(", "-1000", "1000", ")", "FO"}
	short := []string{"(", "-1000", "1000", ")", "FO"}

	res, _ := MatchSyntax(tokens, template)
	if res != true {
		t.Errorf("MatchSyntax(%q, %q) should have matched but did not", tokens, template)
	}

	res, _ = MatchSyntax(bogus_tokens, template)
	if res != false {
		t.Errorf("MatchSyntax(%q, %q) should have matched but did not", bogus_tokens, template)
	}

	_, err := MatchSyntax(short, template)
	if err == nil {
		t.Errorf("MatchSyntax(%q, %q) should have errored on length mismatch", short, template)
	}
	if err.Error() != "MatchSyntax: lengths don't match" {
		t.Error("Error should have been: 'MatchSyntax: lengths don't match")
	}
}
