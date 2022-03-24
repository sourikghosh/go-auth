package pkg

import (
	"testing"
)

func TestGenerateFromPassword(t *testing.T) {
	password := "dsdas"
	hash, err := GenerateFromPassword(password)
	if err != nil {
		t.Fail()
		t.Errorf("expect no err but got %s", err.Error())
	}

	if hash == "" {
		t.Fail()
		t.Error("expect hash but got empty string")
	}

	password = ""
	hash2, err := GenerateFromPassword(password)
	if err != nil {
		t.Fail()
		t.Errorf("expect no err but got %s", err.Error())
	}

	if hash2 == "" {
		t.Fail()
		t.Error("expect hash but got empty string")
	}
}
