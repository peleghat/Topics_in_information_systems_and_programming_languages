package ErrorsFolder

import (
	"errors"
	_ "errors"
	_ "fmt"
)

// Set of common errors, used in our project, in the database functions

var ErrDbConnection = errors.New("failed to connect to the db")
var ErrAlreadyExist = errors.New("the value already exist in the table")
var ErrNotExist = errors.New("the value not exist in the table")
var ErrIllegalValues = errors.New("contains illegal values")
var ErrDbQuery = errors.New("invalid query")
