package apperrors_test

import (
	"testing"

	"github.com/guckppap/gukppap-backend/pkg/apperrors"
)

func TestAppErrors(t *testing.T) {
	err := apperrors.New(apperrors.Internal, "hi")
	err1 := apperrors.New(apperrors.Internal, "by")
	if apperrors.IsEqual(apperrors.Internal, err, err1) == false {
		t.Fatalf("에러 발생!, 두개의 error의 타입은 intneral이므로 true가 나와야 합니다.")
	}

	err = apperrors.New(apperrors.Internal, "nofoudn!")
	if apperrors.IsEqual(apperrors.Internal, err) == true {
		t.Fatalf("에러 발생!, Internal 에러가 아니므로 false가 나와야 함!")
	}
}
