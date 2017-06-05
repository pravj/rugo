// Package utils implements the utility functions for the interpreter.
package utils

// CheckError panics if there is an error.
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// IsDigitCharacter returns true if the input character represents a digit (0-9).
func IsDigitCharacter(c string) bool {
	return (len(c) == 1 && (c >= "0" && c <= "9"))
}

// IsAlphaCharacter returns true if the input character represents an alphabet.
// The following set of characters are allowed, (a-z, A-Z, _).
func IsAlphaCharacter(c string) bool {
	return ((len(c) == 1) && ((c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c == "_")))
}

// IsAlphaNumericCharacter returns true if the input character is either digit or alphabet.
func IsAlphaNumericCharacter(character string) bool {
	return (len(character) == 1 && (IsDigitCharacter(character) || IsAlphaCharacter(character)))
}
