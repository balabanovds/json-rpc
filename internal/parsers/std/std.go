package std

import (
	"encoding/json"
	"io"

	"github.com/balabanovds/jsonrpc/internal/parsers"
)

type Parser struct{}

func New() parsers.JSONParser {
	return &Parser{}
}

func (p *Parser) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (p *Parser) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (p *Parser) Write(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

func (p *Parser) Reader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
