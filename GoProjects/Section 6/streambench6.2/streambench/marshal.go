package streambench

import (
	"encoding/json"
	"io"
)

// Marshal does some marshaling.
func Marshal(w io.Writer) {
	j, _ := json.Marshal(&Data{
		ID:       "abcde",
		Name:     "Ben",
		Email:    "ben@ben.website",
		Whatever: 100,
	})
	w.Write(j)
}
