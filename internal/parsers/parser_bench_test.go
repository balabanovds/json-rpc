package parsers_test

import (
	"math"
	"testing"

	"github.com/balabanovds/jsonrpc/internal/parsers"
	"github.com/balabanovds/jsonrpc/internal/parsers/iter"
	"github.com/balabanovds/jsonrpc/internal/parsers/std"
	"github.com/stretchr/testify/require"
)

func BenchmarkParserUnmarshal(b *testing.B) {
	tests := []struct {
		name   string
		parser parsers.JSONParser
		input  []byte
		dest   interface{}
	}{
		{
			name:   "std small",
			parser: std.New(),
			input:  smallInput,
			dest:   smallPayload{},
		},
		{
			name:   "iter small",
			parser: iter.New(),
			input:  smallInput,
			dest:   smallPayload{},
		},
		{
			name:   "iter fastest small",
			parser: iter.NewFastest(),
			input:  smallInput,
			dest:   smallPayload{},
		},
		{
			name:   "std medium",
			parser: std.New(),
			input:  mediumInput,
			dest:   mediumPayload{},
		},
		{
			name:   "iter medium",
			parser: iter.New(),
			input:  mediumInput,
			dest:   mediumPayload{},
		},
		{
			name:   "iter fastest medium",
			parser: iter.NewFastest(),
			input:  mediumInput,
			dest:   mediumPayload{},
		},
	}

	for _, tt := range tests {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			var got smallPayload

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				err := tt.parser.Unmarshal(tt.input, &got)
				require.NoError(b, err)
			}
		})
	}
}

func BenchmarkParserMarshal(b *testing.B) {

	input := person{Name: "Den", Email: "den@mail.com", Float: math.MaxFloat64}

	tests := []struct {
		name   string
		parser parsers.JSONParser
	}{
		{
			name:   "std",
			parser: std.New(),
		},
		{
			name:   "iter",
			parser: iter.New(),
		},
		{
			name:   "iter fastest",
			parser: iter.NewFastest(),
		},
	}

	for _, tt := range tests {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, _ = tt.parser.Marshal(&input)
			}
		})
	}
}
