package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", r)

	quotedJSON := strconv.Quote(json)

	return []byte(quotedJSON), nil
}
