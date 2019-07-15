package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func fizzBuzz(val int64, w http.ResponseWriter) {
	buf := bytes.NewBuffer(nil)

	for i := int64(0); i <= val; i++ {
		s := ""
		if i%3 == 0 {
			s += "fizz"
		}
		if i%5 == 0 {
			s += "buzz"
		}

		buf.WriteString(fmt.Sprintf("%d: %s\n", i, s))
	}

	w.Write(buf.Bytes())
}
