package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func handleFizzBuzz(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if v, ok := vars["val"]; ok {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			fizzBuzz(i, w)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Error converting string '%s' to int64 for fizzbuzz value", v)))
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error finding fizzbuzz value"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{val}", handleFizzBuzz)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
