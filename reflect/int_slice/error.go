package main

// ErrorMessage implement error interfacer with a msg string
type ErrorMessage struct {
	msg string
}

func (e *ErrorMessage) Error() string {
	return e.msg
}

// ERR_ELEMENT_EXIST
var ERR_ELEMENT_EXIST = &ErrorMessage{"ERROR: ELEMENT EXIST"}

// ERR_ELEMENT_NOT_EXIST
var ERR_ELEMENT_NOT_EXIST = &ErrorMessage{"ERROR: ELEMENT NOT EXIST"}
