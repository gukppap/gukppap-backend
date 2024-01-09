package apperrors

var (
	InternalServerError = newErrIdent("internalServerError")
	NotFoundError       = newErrIdent("NotFoundError")
	BadReqeustError     = newErrIdent("BadRequestError")
	RequestParsingError = newErrIdent("RequestParsingError")
)
