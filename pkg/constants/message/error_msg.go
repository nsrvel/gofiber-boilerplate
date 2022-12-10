package general

import "fmt"

//* Body parser
const ErrBodyParser = "Please check your input!"
const ErrBodyParserInd = "Tolong periksa input anda!"

//* Integer
func ErrValueIsNotInteger(field string) string {
	return fmt.Sprintf("Make sure '%s' is integer!", field)
}
func ErrValueIsNotIntegerInd(field string) string {
	return fmt.Sprintf("Pastikan value dari '%s' adalah bilangan bulat!", field)
}
