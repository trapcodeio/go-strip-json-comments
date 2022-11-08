package strip_json_comments

import "testing"

// Package is an exact port of `Sindresorhus` strip-json-comments nodejs package.
// Ported by `trapcodeio`
// File test.js in repository: https://github.com/sindresorhus/strip-json-comments

func Test_Replace_Comments_With_WhiteSpace(t *testing.T) {
	json, expected := "//comment\n{\"a\":\"b\"}", "         \n{\"a\":\"b\"}"
	striped := StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `/*//comment*/{"a":"b"}`, `             {"a":"b"}`
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = "{\"a\":\"b\"//comment\n}", "{\"a\":\"b\"         \n}"
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `{"a":"b"/*comment*/}`, `{"a":"b"           }`
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = "{\"a\"/*\n\n\ncomment\n*/:\"b\"}", "{\"a\"  \n\n\n       \n  :\"b\"}"
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = "/*!\n * comment\n */\n{\"a\":\"b\"}", "   \n          \n   \n{\"a\":\"b\"}"
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `{/*comment*/"a":"b"}`, `{           "a":"b"}`
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}
}

func Test_Remove_Comments(t *testing.T) {
	options := &Options{Whitespace: false}

	json, expected := "//comment\n{\"a\":\"b\"}", "\n{\"a\":\"b\"}"
	striped := StripJsonCommentsWithOptions(json, options)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `/*//comment*/{"a":"b"}`, `{"a":"b"}`
	striped = StripJsonCommentsWithOptions(json, options)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = "{\"a\":\"b\"//comment\n}", "{\"a\":\"b\"\n}"
	striped = StripJsonCommentsWithOptions(json, options)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `{"a":"b"/*comment*/}`, `{"a":"b"}`
	striped = StripJsonCommentsWithOptions(json, options)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}
}

func Test_Doesnt_Strip_Comments_Inside_Strings(t *testing.T) {
	json, expected := `{"a":"b//c"}`, `{"a":"b//c"}`
	striped := StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `{"a":"b/*c*/"}`, `{"a":"b/*c*/"}`
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = `{"/*a/":"b"}`, `{"/*a/":"b"}`
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}

	json, expected = "{\"\\\"/*a\":\"b\"}", "{\"\\\"/*a\":\"b\"}"
	striped = StripJsonComments(json)

	if striped != expected {
		t.Errorf("Expected %s, got %s", expected, striped)
		return
	}
}
