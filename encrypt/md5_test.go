package encrypt_test

import (
	"testing"

	"github.com/ak5w/gkit/encrypt"
)

type testCase struct {
	desc  string
	input string
	wannt string
}

var testGroup = []testCase{
	testCase{
		desc:  "normal simple case",
		input: "123456",
		wannt: "e10adc3949ba59abbe56e057f20f883e",
	},
}

func TestMD5(t *testing.T) {

	for _, test := range testGroup {
		got := encrypt.Md5(test.input)
		if got != test.wannt {
			t.Errorf("for the case %#v, failed details as: the values of %#v is not %#v\n", test.desc, got, test.wannt)
		}
	}
}
