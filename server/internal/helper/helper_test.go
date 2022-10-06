package helper

import "testing"

func TestHttpStatusToString(t *testing.T) {
	if "234" != HttpStatusToString(234) {
		t.Error("expected: 234, got:", HttpStatusToString(200))
	}
}

func TestStringTohttpStatus(t *testing.T) {
	if 234 != StringToHttpStatus("234") {
		t.Error("expected: 234, got:", StringToHttpStatus("234"))
	}
}
