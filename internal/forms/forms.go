package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form)Required(fields ...string){
	for _, field := range fields{
		value := f.Get(field)
		if strings.TrimSpace(value) == ""{
			f.Errors.Add(field, fmt.Sprintf("Field [%s] cannot be blanc", field))
		}
	}
}

// Valid returns true when there are no errors
func (f *Form)Valid() bool{
	return len(f.Errors) == 0
}

// Has checks is field is in request and is not empty
func (f *Form)Has(field string, r *http.Request) bool{
	fd := r.Form.Get(field)
	if fd == ""{
		f.Errors.Add(field, fmt.Sprintf("Field [%s] cannot be blanc", field))
		return false
	}

	return true
}