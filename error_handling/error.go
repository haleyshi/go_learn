package main

import "fmt"

// DivideError implement error interface
type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	devidee: %d
	devider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 =", result)
	}

	if _, errorMsg := divide(100, 0); errorMsg != "" {
		fmt.Println("100/0 get error", errorMsg)
	}
}
