package we

import (
	"strings"
)

// EscapeForDotEnv escapes val to use it in .env file read by react-scripts.
//
// The value is expected to place between "" like this VARNAME="%s"
// I'm not sure the implemeted process follows any specs but it seems ok for me...
// https://create-react-app.dev/docs/adding-custom-environment-variables/
func EscapeForDotEnv(val string) string {
	ret := val

	// Add leading \ to each special charactors $, \
	for _, c := range []string{"$", "\\"} {
		ret = strings.ReplaceAll(ret, c, "\\"+c)
	}

	// Keep " as it is
	//
	// for _, c := range []string{"\""} {
	//   	ret = strings.ReplaceAll(ret, c, "\""+c)
	// }

	// Replace carriage return with \n
	ret = strings.ReplaceAll(ret, "\n", "\\n")

	return ret
}

// \ -> \\
// $ -> \$
// " -> "   (as it is)
// carriage return -> \n
