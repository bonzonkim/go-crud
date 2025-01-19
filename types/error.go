package types

import "fmt"

const (
	NotFoundUser = iota
	DBError
)

var errMessage = map[int64]string {
	NotFoundUser:	"User Not Found",
	DBError:		"DB Error occurred",
}

func Errorf(code int64, args ...interface{}) error {
	if msg, ok := errMessage[code]; ok {
		return fmt.Errorf("%s	:	%v", msg, args)
	} else {
		return fmt.Errorf("not found error code")
	}
}
