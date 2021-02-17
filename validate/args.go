package validate

import "fmt"

// Args checks to see if the required no. of command line flags were passed
func Args(length int, args []string, validationFn func(args []string) (err error)) (err error) {
	if len(args) < length {
		return fmt.Errorf("Missing Arguments")
	}

	err = validationFn(args)

	return
}
