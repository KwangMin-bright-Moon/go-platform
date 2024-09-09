package util

import (
	"fmt"
)

func AnyToString1(value any) string {
	if value == nil {
		return "0"
	}

	return fmt.Sprintf("%v", value)
}
