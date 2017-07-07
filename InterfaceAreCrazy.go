package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/**
 * Interace is hard in GO
 * Check this out again when I am ready
 * http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
 *
 * @param  url  an absolute URL giving the base location of the image
 * @param  name the location of the image, relative to the url argument
 * @return      the image at the specified URL
 * @see         Image
 * @7/7/2017 3:59 PM
 * Author : cyrsis@github 
 */

// start with a string representation of our JSON data
var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func main() {
	// our target will be of type map[string]interface{}, which is a pretty generic type
	// that will give us a hashtable whose keys are strings, and whose values are of
	// type interface{}
	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}

