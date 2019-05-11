package bob

import "strings"
import "unicode"

// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
// * Bob answers 'Sure.' if you ask him a question.
//   (ends with a question mark)
// * He answers 'Whoa, chill out!' if you yell at him.
//   (all caps)
// * He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//   (all caps + question mark)
// * He says 'Fine. Be that way!' if you address him without actually saying anything.
//   (whitespace)
// * He answers 'Whatever.' to anything else.
//
// From the tests, it's assuming that we'll be chomping the string (trimming whitespace from ends)

const (
	calmDown   = "Calm down, I know what I'm doing!"
	chillOut   = "Whoa, chill out!"
	sure       = "Sure."
	thisIsFine = "Fine. Be that way!"
	whatevs    = "Whatever."
)
// Hey func take remark (sentence) as an string
func Hey(remark string) string {

	isEmpty := false
	isVerbiage := false
	isQuestion := false
	isYell := false

	// chomp
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		isEmpty = true
	}

	if strings.IndexFunc(remark, unicode.IsLetter) p>= 0 {
		isVerbiage = true
	}

	if strings.HasSuffix(remark, "?") {
		isQuestion = true
	}

	upper := strings.ToUpper(remark)
	if isVerbiage && upper == remark {
		isYell = true
	}

	// Print out the truth table For caveman debugging.
	// fmt.Printf("question: %v, yell %v, verbiage %v, empty %v\n",
	//		isQuestion, isYell, isVerbiage, isEmpty)

	if isVerbiage && isYell && isQuestion {
		return calmDown
	} else if isYell {
		return chillOut
	} else if isQuestion {
		return sure
	} else if isEmpty {
		return thisIsFine
	} else {
		return whatevs
	}
}
