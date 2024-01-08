package verifier_test

import (
	"fmt"
	"testing"

	"github.com/guckppap/gukppap-backend/pkg/verifier"
)

type TestStruct struct {
	Url string `validate:"required,url"`
}

func TestVerifier(t *testing.T) {
	u := new(TestStruct)
	err := verifier.Run(u)
	if err == nil {
		t.Fatalf("nil != 이어야 함!")
	}
	fmt.Println(err.Error())

	u.Url = "안녕하세요 반가습니다"
	err = verifier.Run(u)
	if err == nil {
		t.Fatalf("nil != 이어야 함!")
	}
	fmt.Println(err.Error())

	u.Url = "https://localhost:9190/api"

	err = verifier.Run(u)
	if err != nil {
		t.Fatalf("nil == 이어야 함!")
	}
}
