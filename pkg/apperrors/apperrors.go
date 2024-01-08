package apperrors

// errIdentType 타입은 식별 타입입니다.
// appError 구조체에 사용되며, error를 식별하기 위해 사용됩니다.
type errIdentType string

// newErrIdent는 errIdentType의 생성자 입니다.
func newErrIdent(errIdent string) errIdentType {
	return errIdentType(errIdent)
}

// appError 구조체는 이 앱의 에러 핸들링에 관해 쉽게 처리하는 주체 입니다.
type appError struct {
	errIdent errIdentType
	cause    string
}

// Error 함수는 error Interface의 구현 함수 입니다.
func (a *appError) Error() string {
	return a.cause
}

// New 함수는 appError의 생성자 입니다. 사용자는 이 함수를 통해 error를 제어 하게 됩니다.
func New(cause string, errIdent errIdentType) error {
	return &appError{
		cause:    cause,
		errIdent: errIdent,
	}
}

// IsEqual 함수는 매겨변수인 ident와 errors의 ident를 비교하여 일치한지 확인하여 참/ 거짓을 반환하는 함수 입니다.
func IsEqual(errIdent errIdentType, errors ...error) bool {
	for _, err := range errors {
		// err를 appError로 변경
		appError, ok := err.(*appError)
		if ok == false || appError.errIdent != errIdent {
			return false
		}
	}

	return true
}
