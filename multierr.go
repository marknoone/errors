package multierr

import (
	"fmt"
	"strings"
)

type MultiError []error

func (me MultiError) Error() string {
	if len(me) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d error(s) occurred:", len(me)))
	for _, err := range me {
		sb.WriteString(fmt.Sprintf("\n* %s", err))
	}
	return sb.String()
}

// Append appends the provided error if it is not nil.
func (me *MultiError) Append(err error) {
	if err != nil {
		*me = append(*me, err)
	}
}
