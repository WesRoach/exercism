package bob

import "unicode"

// Bob's responses to people who wont stop barging in and talking to him
func Hey(remark string) string {
	contains_lowercase := false
	contains_uppercase := false
	contains_number := false
	ends_with_question_mark := false
	idx_latest_non_qmark := -1
	idx_latest_qmark := -1

	for idx, r := range remark {
		if r == []rune("?")[0] {
			idx_latest_qmark = idx
		} else if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				contains_uppercase = true
			} else {
				contains_lowercase = true
			}
			idx_latest_non_qmark = idx
		} else if unicode.IsNumber(r) {
			contains_number = true
			idx_latest_non_qmark = idx
		}
	}

	if idx_latest_qmark > idx_latest_non_qmark {
		ends_with_question_mark = true
	}

	if (contains_uppercase && ends_with_question_mark) && !contains_lowercase {
		return "Calm down, I know what I'm doing!"
	} else if contains_uppercase && !contains_lowercase {
		return "Whoa, chill out!"
	} else if ends_with_question_mark {
		return "Sure."
	} else if !(contains_lowercase || contains_uppercase || contains_number) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
