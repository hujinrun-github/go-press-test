package constant

type ErrorCode int

const (
	ERR_CODE_SUCCESS        ErrorCode = 0
	ERR_CODE_REQUEST_FAILED ErrorCode = -100
	ERR_TYPE_CONVERT        ErrorCode = -101
	ERR_INVALID_PARAM       ErrorCode = -102
	ERR_OPEN_FILE           ErrorCode = -103
	ERR_READ_FILE           ErrorCode = -104
	ERR_READ_FILE_LINE      ErrorCode = -105

	ERR_UNKNOW ErrorCode = -1000000000
)
