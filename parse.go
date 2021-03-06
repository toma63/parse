package parse

import ("os"
        "bufio"
        "regexp"
)

// line reader for concurrent use
// filepath, line comment string, and an output chan containing lines
// of text.  Comments and blank lines are stripped
// Closes the channel on EOF
func ReadLinesStripped(filepath string, line_comment string, lines chan<- string) {

        // match comment string to end of line
        com_re := regexp.MustCompile(line_comment + ".*$")

        fd, err := os.Open(filepath)

        if err != nil {
                panic(err)
        }

        defer fd.Close()

        scanln := bufio.NewScanner(fd)

        for scanln.Scan() {
                line := scanln.Text()
                strip_line := com_re.ReplaceAllLiteralString(line, "")
                if strip_line == "" { continue }
                lines <- strip_line
        }
        if err := scanln.Err(); err != nil {
                panic(err)
        }

        close(lines)
}


// given a channel of lines, split into tokens given an re object
// results go to an output chan of strings
func SplitTokenizer(split_re *regexp.Regexp, lines <-chan string, tokens chan<- string) {

        for line := range lines {
                for _, token := range split_re.Split(line, -1) {
                        if token == "" { continue }
                        tokens <- token
                }
        }
        close(tokens)
}

// take n tokens from a string channel
func TakeN(n int, ch <-chan string) []string {
	res := make([]string, n)
	for i := 0 ; i < n ; i++ {
		tok, ok := <- ch
		if !ok { panic("parse error, premature end of file") } 
		res = append(res, tok)
	}
	return(res)
}

// remove tokens from a string channel until stop string is seen
func TakeUntil(stop string, ch <-chan string) []string {
	res := []string{}
	for {
		tok, ok := <- ch
		if ! ok { panic("parse error, premature end of file") } 
		res = append(res, tok)
		if tok == stop { return(res) }
	}
}

// remove tokens from a string channel until stop regexp is seen
func TakeUntilRE(stop_re *regexp.Regexp, ch <-chan string) []string {
	res := []string{}
	for {
		tok, ok := <- ch
		if ! ok { panic("parse error, premature end of file") } 
		res = append(res, tok)
		if stop_re.MatchString(tok) { return(res) }
	}
}

// map a slice of regexps against a slice of strings
// returning a boolean indicating whether all match
func MatchSyntax(tokens []string, template []*regexp.Regexp) bool {

	//check for length match
	size := len(template)
	if size != len(tokens) {
		return false
	}

	match := true

	for i := 0 ; i < size ; i++ {
		if ! template[i].MatchString(tokens[i]) {
			match = false
		}
	}
	return match
}
