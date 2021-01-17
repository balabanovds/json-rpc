package parsers_test

import (
	"testing"

	"github.com/balabanovds/jsonrpc/internal/parsers/iter"
	"github.com/balabanovds/jsonrpc/internal/parsers/std"
	"github.com/stretchr/testify/require"
)

type person struct {
	Name  string `json:"name"`
	Email string `json:"e-mail"`
	Year  int    `json:"birth_year"`
	Meta  string
}

var input = `{"name":"Den","e-mail":"den@mail.com","birth_year":1970}`

var want = person{Name: "Den", Email: "den@mail.com", Year: 1970}

func TestParser_Unmarshal(t *testing.T) {
	t.Run("std ok", func(t *testing.T) {
		j := std.New()
		var got person
		err := j.Unmarshal([]byte(input), &got)

		require.NoError(t, err)
		require.Equal(t, want, got)
	})

	t.Run("iter ok", func(t *testing.T) {
		j := iter.New()
		var got person
		err := j.Unmarshal([]byte(input), &got)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})

}
