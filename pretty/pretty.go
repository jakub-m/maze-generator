package pretty

import (
	"encoding/json"
	"fmt"
	"maze/cell"
)

func Format(c cell.Cell) string {
	b, err := json.MarshalIndent(c, "", " ")
	if err == nil {
		return string(b)
	} else {
		return fmt.Sprintf("ERROR: %v", err)
	}
}
