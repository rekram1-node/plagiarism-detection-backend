package handlers

import (
	"regexp"
	"strings"

	"github.com/neurosnap/sentences/english"
)

func citationCheck(str string) (string, []string) {
	var newStr string
	var validMLA []string

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}
	sentences := tokenizer.Tokenize(str)

	for _, sentence := range sentences {
		s := sentence.Text
		if !validCitation(s) {
			newStr += strings.TrimSpace(s) + " "
			continue
		}
		validMLA = append(validMLA, strings.TrimSpace(s))
	}

	return newStr, validMLA
}

func validCitation(s string) bool {
	myValid := regexp.MustCompile(`\("[A-Za-z0-9,./:?'"\s]+"|[A-Za-z0-9,./:?'"\s]+\)[.!?]\z`)
	return myValid.MatchString(s)
}
