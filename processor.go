package jsonrpc

import (
	"context"

	"github.com/balabanovds/jsonrpc/internal/parsers"
)

type processor struct {
	methods map[string]Method
	json    parsers.JSONParser
}

func newProcessor(methods map[string]Method, json parsers.JSONParser) *processor {
	return &processor{methods, json}
}

func (p *processor) Call(ctx context.Context, data []byte) []byte {
	return nil
}
