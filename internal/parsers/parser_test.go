package parsers_test

import (
	"testing"

	"github.com/balabanovds/jsonrpc/internal/parsers"
	"github.com/balabanovds/jsonrpc/internal/parsers/iter"
	"github.com/balabanovds/jsonrpc/internal/parsers/std"
	"github.com/stretchr/testify/require"
)

type person struct {
	Name  string  `json:"name"`
	Email string  `json:"e-mail"`
	Year  int     `json:"birth_year,omitempty"`
	Float float64 `json:"float,omitempty"`
}

func TestParser_Unmarshal(t *testing.T) {
	input := []byte(`{"name":"Den","e-mail":"den@mail.com","birth_year":1970,"float":0.123456789}`)
	want := person{Name: "Den", Email: "den@mail.com", Year: 1970, Float: 0.123456789}

	tests := []struct {
		name   string
		parser parsers.JSONParser
	}{
		{
			name:   "std ok",
			parser: std.New(),
		},
		{
			name:   "iter compatible to std ok",
			parser: iter.New(),
		},
		{
			name:   "iter fastest ok",
			parser: iter.NewFastest(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var got person
			err := tt.parser.Unmarshal(input, &got)

			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	}
}

func TestParser_Marshal(t *testing.T) {
	tests := []struct {
		name   string
		input  person
		want   string
		parser parsers.JSONParser
	}{
		{
			name:   "std",
			input:  person{Name: "Den", Email: "den@mail.com", Year: 1970},
			want:   `{"name":"Den","e-mail":"den@mail.com","birth_year":1970}`,
			parser: std.New(),
		},
		{
			name:   "std omitempty",
			input:  person{Name: "Den", Email: "den@mail.com"},
			want:   `{"name":"Den","e-mail":"den@mail.com"}`,
			parser: std.New(),
		},
		{
			name:   "std float",
			input:  person{Name: "Den", Email: "den@mail.com", Float: 0.123456789},
			want:   `{"name":"Den","e-mail":"den@mail.com","float":0.123456789}`,
			parser: std.New(),
		},
		{
			name:   "iter",
			input:  person{Name: "Den", Email: "den@mail.com", Year: 1970},
			want:   `{"name":"Den","e-mail":"den@mail.com","birth_year":1970}`,
			parser: iter.New(),
		},
		{
			name:   "iter omitempty",
			input:  person{Name: "Den", Email: "den@mail.com"},
			want:   `{"name":"Den","e-mail":"den@mail.com"}`,
			parser: iter.New(),
		},
		{
			name:   "iter float compatible to std",
			input:  person{Name: "Den", Email: "den@mail.com", Float: 0.123456789},
			want:   `{"name":"Den","e-mail":"den@mail.com","float":0.123456789}`,
			parser: iter.New(),
		},
		{
			name:   "iter float fastest",
			input:  person{Name: "Den", Email: "den@mail.com", Float: 0.123456789},
			want:   `{"name":"Den","e-mail":"den@mail.com","float":0.123457}`,
			parser: iter.NewFastest(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.parser.Marshal(&tt.input)

			require.NoError(t, err)
			require.Equal(t, tt.want, string(got))
		})
	}
}
