package strip_json_comments

import (
	"regexp"
)

// Package is an exact port of `Sindresorhus` strip-json-comments nodejs package.
// Ported by `trapcodeio`
// File index.js in repository: https://github.com/sindresorhus/strip-json-comments

const notInsideComment = 0
const singleComment = 1
const multiComment = 2

type Options struct {
	Whitespace     bool
	TrailingCommas bool
}

func stripWithoutWhitespace() string {
	return ""
}

func stripWithWhitespace(string string, start *int, end *int) string {
	if start != nil && end != nil {
		string = string[*start:*end]
	} else if start != nil {
		string = string[*start:]
	}

	m := regexp.MustCompile(`\S`)
	c := m.ReplaceAllString(string, " ")

	return c
}

func isEscaped(jsonString string, quotePosition int) bool {
	index := quotePosition - 1
	backslashCount := 0

	for string(jsonString[index]) == "\\" {
		index -= 1
		backslashCount += 1
	}

	return backslashCount%2 == 1
}

// StripJsonCommentsWithOptions - strips comments from JSON string with options
func StripJsonCommentsWithOptions(jsonString string, options *Options) string {
	if options == nil {
		options = &Options{Whitespace: true, TrailingCommas: false}
	}

	isInsideString := false
	isInsideComment := notInsideComment
	offset := 0
	buffer := ""
	result := ""
	commaIndex := -1

	strip := func(index int) string {
		if options.Whitespace {
			return stripWithWhitespace(jsonString, &offset, &index)
		} else {
			return stripWithoutWhitespace()
		}
	}

	for index := 0; index < len(jsonString); index++ {
		currentCharacter := string(jsonString[index])
		nextCharacter := ""

		if index+1 < len(jsonString) {
			nextCharacter = string(jsonString[index+1])
		}

		if isInsideComment == notInsideComment && currentCharacter == `"` {
			// Enter or exit string
			escaped := isEscaped(jsonString, index)
			if !escaped {
				isInsideString = !isInsideString
			}
		}

		if isInsideString {
			continue
		}

		if isInsideComment == notInsideComment && currentCharacter+nextCharacter == "//" {
			// Enter single-line comment
			buffer += jsonString[offset:index]
			offset = index
			isInsideComment = singleComment
			index++
		} else if isInsideComment == singleComment && currentCharacter+nextCharacter == "\r\n" {
			// Exit single-line comment via \r\n
			index++
			isInsideComment = singleComment
			buffer += strip(index)
			offset = index
			continue
		} else if isInsideComment == singleComment && currentCharacter == "\n" {
			// Exit single-line comment via \n
			isInsideComment = notInsideComment
			buffer += strip(index)
			offset = index
		} else if isInsideComment == notInsideComment && currentCharacter+nextCharacter == "/*" {
			// Enter multiline comment
			buffer += jsonString[offset:index]
			offset = index
			isInsideComment = multiComment
			index++
			continue
		} else if isInsideComment == multiComment && currentCharacter+nextCharacter == "*/" {
			// Exit multiline comment
			index++
			isInsideComment = notInsideComment
			buffer += strip(index + 1)
			offset = index + 1
			continue
		} else if options.TrailingCommas && isInsideComment == notInsideComment {
			if commaIndex != -1 {
				if currentCharacter == "}" || currentCharacter == "]" {
					// Strip trailing comma
					buffer += jsonString[offset:index]
					result += strip(1) + buffer[1:]
					buffer = ""
					offset = index
					commaIndex = -1
				} else if currentCharacter != " " && currentCharacter != "\t" && currentCharacter != "\r" && currentCharacter != "\n" {
					// Hit non-whitespace following a comma; comma is not trailing
					buffer += jsonString[offset:index]
					offset = index
					commaIndex = -1
				}
			} else if currentCharacter == "," {
				// Flush buffer prior to this point, and save new comma index
				result += buffer + jsonString[offset:index]
				buffer = ""
				offset = index
				commaIndex = index
			}
		}
	}

	var end string
	if isInsideComment > notInsideComment {
		if options.Whitespace {
			end = stripWithWhitespace(jsonString[offset:], nil, nil)
		} else {
			end = stripWithoutWhitespace()
		}

	} else {
		end = jsonString[offset:]
	}

	return result + buffer + end
}

// StripJsonComments - strips comments from JSON string
func StripJsonComments(jsonString string) string {
	return StripJsonCommentsWithOptions(jsonString, nil)
}
