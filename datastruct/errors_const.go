package datastruct

// Error code definition
const (
	//App Error

	// Header status
	HeaderStatusOk   int = 899
	StatusBadRequest int = 997

	//App Error
	//Status OK value <=500
	ErrSuccess          int = 00
	ErrFailedQuery      int = 100
	ErrInvalidParameter int = 101
	ErrReadDBValue      int = 102
	ErrGetToken         int = 103

	ErrStringConvert         int = 305
	ErrInvalidFormat         int = 306
	ErrGetCategoryListFailed int = 307
	ErrGetRequiredInfoFailed int = 308
	ErrInvalidToken          int = 330

	//value 901-998 Unauthorized
	ErrUnauthorized      int = 901
	ErrWrongUserPassword int = 902

	//value >999 //Internal Server Error
	ErrOthers int = 999
)

// Error message definition
const (
	DescSuccess               string = "Success"
	DescGetToken              string = "Failed to get token"
	DescInvalidToken          string = "Invalid Token"
	DescFailedRequest         string = "Failed Request API"
	DescGetCategoryListFailed string = "Get Category List Failed"
	DescGetRequiredInfoFailed string = "Get Required Info Failed"
)
