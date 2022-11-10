package stripjsoncomments

import "testing"

// Package is an exact port of `Sindresorhus` strip-json-comments nodejs package.
// Ported by `trapcodeio`
// File test.js in repository: https://github.com/sindresorhus/strip-json-comments

func Test_Replace_Comments_With_WhiteSpace(t *testing.T) {
	test("//comment\n{\"a\":\"b\"}", "         \n{\"a\":\"b\"}", t)
	test("/*//comment*/{\"a\":\"b\"}", "             {\"a\":\"b\"}", t)
	test("{\"a\":\"b\"//comment\n}", "{\"a\":\"b\"         \n}", t)
	test(`{"a":"b"/*comment*/}`, `{"a":"b"           }`, t)
	test("{\"a\"/*\n\n\ncomment\r\n*/:\"b\"}", "{\"a\"  \n\n\n       \r\n  :\"b\"}", t)
	test("/*!\n * comment\n */\n{\"a\":\"b\"}", "   \n          \n   \n{\"a\":\"b\"}", t)
}

func Test_Remove_Comments(t *testing.T) {
	options := &Options{Whitespace: false}

	testWithOptions("//comment\n{\"a\":\"b\"}", "\n{\"a\":\"b\"}", options, t)
	testWithOptions(`/*//comment*/{"a":"b"}`, `{"a":"b"}`, options, t)
	testWithOptions("{\"a\":\"b\"//comment\n}", "{\"a\":\"b\"\n}", options, t)
	testWithOptions(`{"a":"b"/*comment*/}`, `{"a":"b"}`, options, t)
}

func Test_Doesnt_Strip_Comments_Inside_Strings(t *testing.T) {
	test(`{"a":"b//c"}`, `{"a":"b//c"}`, t)
	test(`{"a":"b/*c*/"}`, `{"a":"b/*c*/"}`, t)
	test(`{"/*a/":"b"}`, `{"/*a/":"b"}`, t)
	test("{\"\\\"/*a\":\"b\"}", "{\"\\\"/*a\":\"b\"}", t)
}

func Test_Consider_Escaped_Slashes_When_Checking_For_Escaped_String_Quote(t *testing.T) {
	test("{\"\\\\\":\"https://foobar.com\"}", "{\"\\\\\":\"https://foobar.com\"}", t)
	test("{\"foo\\\"\":\"https://foobar.com\"}", "{\"foo\\\"\":\"https://foobar.com\"}", t)
}

func Test_Line_Endings_No_Comments(t *testing.T) {
	test("{\"a\":\"b\"\n}", "{\"a\":\"b\"\n}", t)
	test("{\"a\":\"b\"\r\n}", "{\"a\":\"b\"\r\n}", t)
}

func Test_Line_Endings_With_Single_Line_Comment(t *testing.T) {

	test("{\"a\":\"b\"//c\n}", "{\"a\":\"b\"   \n}", t)
	test("{\"a\":\"b\"//c\r\n}", "{\"a\":\"b\"   \r\n}", t)
}

func Test_Line_Endings_With_Line_Block_Comment(t *testing.T) {
	test("{\"a\":\"b\"/*c*/\n}", "{\"a\":\"b\"     \n}", t)
	test("{\"a\":\"b\"/*c*/\r\n}", "{\"a\":\"b\"     \r\n}", t)
}

func Test_Line_Endings_With_Multi_Line_Block_Comment(t *testing.T) {
	test("{\"a\":\"b\",/*c\nc2*/\"x\":\"y\"\n}", "{\"a\":\"b\",   \n    \"x\":\"y\"\n}", t)
	test("{\"a\":\"b\",/*c\r\nc2*/\"x\":\"y\"\r\n}", "{\"a\":\"b\",   \r\n    \"x\":\"y\"\r\n}", t)
}

func Test_Line_Endings_Works_At_EOF(t *testing.T) {
	options := &Options{Whitespace: false}

	test("{\"a\":\"b\"\r\n} //EOF", "{\"a\":\"b\"\r\n}      ", t)
	testWithOptions("{\"a\":\"b\"\r\n} //EOF", "{\"a\":\"b\"\r\n} ", options, t)
}

func Test_Handles_Wierd_Escaping(t *testing.T) {
	test("{\"x\":\"x \\\"sed -e \\\"s/^.\\\\{46\\\\}T//\\\" -e \\\"s/#033/\\\\x1b/g\\\"\\\"\"}", "{\"x\":\"x \\\"sed -e \\\"s/^.\\\\{46\\\\}T//\\\" -e \\\"s/#033/\\\\x1b/g\\\"\\\"\"}", t)
}

func Test_Strips_Trailing_Commas(t *testing.T) {
	testWithOptions(`{"x":true,}`, `{"x":true }`, &Options{true, true}, t)
	testWithOptions(`{"x":true,}`, `{"x":true}`, &Options{false, true}, t)
	testWithOptions("{\"x\":true,\n  }", "{\"x\":true \n  }", &Options{true, true}, t)
	testWithOptions("[true, false,]", "[true, false ]", &Options{true, true}, t)
	testWithOptions("[true, false,]", "[true, false]", &Options{false, true}, t)
	testWithOptions("{\n  \"array\": [\n    true,\n    false,\n  ],\n}", "{\n  \"array\": [\n    true,\n    false\n  ]\n}", &Options{false, true}, t)
	testWithOptions("{\n  \"array\": [\n    true,\n    false /* comment */ ,\n /*comment*/ ],\n}", "{\n  \"array\": [\n    true,\n    false  \n  ]\n}", &Options{false, true}, t)
}

func Test_Handles_Malformed_Block_Comments(t *testing.T) {
	// pass
	test("[] */", "[] */", t)
	// fail
	testShouldFail("[] /*", "[] /*", t)
}

// ============================================================================
// ============================================================================
// ===== Test Helper Functions ================================================
// ============================================================================
// ============================================================================

// testWithOptions is a shorthand function for testing with options
func testWithOptions(json string, expected string, options *Options, t *testing.T) {
	result := StripWithOptions(json, options)

	if result != expected {
		t.Errorf("Expected %v of lenght %v \n | Got %v of lenght %v", result, len(result), expected, len(expected))
	}
}

// test is a shorthand function for testing
func test(json string, expected string, t *testing.T) {
	testWithOptions(json, expected, nil, t)
}

// testShouldFail - test that should fail
func testShouldFail(json string, expected string, t *testing.T) {
	result := Strip(json)

	if result == expected {
		t.Errorf("Expected %v of lenght %v \n | Got %v of lenght %v", result, len(result), expected, len(expected))
	}
}
