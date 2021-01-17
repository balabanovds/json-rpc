package parsers_test

import (
	"testing"

	"github.com/balabanovds/jsonrpc/internal/parsers"
	"github.com/balabanovds/jsonrpc/internal/parsers/iter"
	"github.com/balabanovds/jsonrpc/internal/parsers/std"
)

func BenchmarkParserUnmarshal(b *testing.B) {
	tests := []struct {
		name   string
		parser parsers.JSONParser
		input  []byte
	}{
		{
			name:   "std",
			parser: std.New(),
			input:  []byte(input),
		}, {
			name:   "iter",
			parser: iter.New(),
			input:  []byte(input),
		},
	}

	for _, tt := range tests {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			var got person

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = tt.parser.Unmarshal(tt.input, &got)
			}
		})
	}

}
