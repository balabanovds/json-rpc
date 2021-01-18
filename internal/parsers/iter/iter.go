package iter

import (
	"io"

	"github.com/balabanovds/jsonrpc/internal/parsers"
	jsoniter "github.com/json-iterator/go"
)

type Parser struct {
	json jsoniter.API
}

// New creates new instance of json-iter
// compatible to standard library
func New() parsers.JSONParser {
	return &Parser{jsoniter.ConfigCompatibleWithStandardLibrary}
}

// NewFastest creates instance thats marshals float with only 6 digits precision
func NewFastest() parsers.JSONParser {
	return &Parser{jsoniter.ConfigFastest}
}

// Customize you parser with config
func NewCustom(config jsoniter.Config) parsers.JSONParser {
	return &Parser{config.Froze()}
}

func (p *Parser) Marshal(v interface{}) ([]byte, error) {
	return p.json.Marshal(v)
}

func (p *Parser) Unmarshal(data []byte, v interface{}) error {
	return p.json.Unmarshal(data, v)
}

func (p *Parser) Write(w io.Writer, data interface{}) error {
	return p.json.NewEncoder(w).Encode(data)
}

func (p *Parser) Read(r io.Reader, v interface{}) error {
	return p.json.NewDecoder(r).Decode(v)
}
