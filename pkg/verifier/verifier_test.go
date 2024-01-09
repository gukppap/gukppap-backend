package verifier_test

import (
	"fmt"
	"testing"

	"github.com/guckppap/gukppap-backend/pkg/verifier"
)

type TestStruct struct {
	Url      string `validate:"required,url"`
	Shortcut string `validate:"custom_shortcut"`
}

func TestVerifier(t *testing.T) {
	sh := TestStruct{
		Url:      "https://localhost/api",
		Shortcut: "7777777",
	}

	err := verifier.Run(sh)
	if err != nil {
		fmt.Println("error happen: ", err)
		t.Fail()
	}

}
