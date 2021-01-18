package parsers

import (
	"io"
)

type JSONParser interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Write(w io.Writer, data interface{}) error
	Read(r io.Reader, v interface{}) error
}
