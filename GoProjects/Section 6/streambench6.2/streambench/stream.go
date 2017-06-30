package streambench

import (
	"encoding/json"
	"io"
)

// Encode encdoes some garbo.
func Encode(w io.Writer) {
	json.NewEncoder(w).Encode(&Data{
		ID:       "abcde",
		Name:     "Ben",
		Email:    "ben@ben.website",
		Whatever: 100,
	})
}
