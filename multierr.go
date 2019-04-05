package errors

import (
	"fmt"
	"strings"
)

type E []error

func (e E) Error() string {
	if len(e) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d error(s) occurred:", len(e)))
	for _, err := range e {
		sb.WriteString(fmt.Sprintf("\n* %s", err))
	}
	return sb.String()
}

// Append appends the provided error if it is not nil.
func (e *E) Append(err error) {
	if err != nil {
		*e = append(*e, err)
	}
}
