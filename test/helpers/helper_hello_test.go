package helper_test

import (
	"testing"

	helper "github.com/adit1o/train3/helpers"
	"github.com/stretchr/testify/assert"
)

// test
func TestSayHello(t *testing.T) {
	assert.Equal(t, "Hello adit", helper.SayHello("adit"), "output must be `Hello adit`")
}

// table test
func TestSayHelloTable(t *testing.T) {
	tables := []struct {
		Name     string
		Expected string
		Request  string
	}{
		{
			Name:     "adit",
			Request:  "adit",
			Expected: "Hello adit",
		},

		{
			Name:     "budi",
			Request:  "budi",
			Expected: "Hello budi",
		},
	}

	for _, table := range tables {
		t.Run(table.Name, func(t *testing.T) {
			assert.Equal(t, table.Expected, helper.SayHello(table.Request), "output must be `"+table.Expected+"`")
		})

	}
}

// benchmark
func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.SayHello("adit")
	}
}

// table benchmark
func BenchmarkSayHelloTable(b *testing.B) {
	tables := []struct {
		Name     string
		Expected string
		Request  string
	}{
		{
			Name:    "adit",
			Request: "adit",
		},
		{
			Name:    "budi",
			Request: "budi",
		},
	}

	for _, benchmark := range tables {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				helper.SayHello(benchmark.Request)
			}
		})
	}
}
