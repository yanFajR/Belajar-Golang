package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i:= 0; i<b.N ; i++ {
		HelloWorld("Eko")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Eko",
			request: "Eko",
			expected: "Hello Eko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip(("Cannot run on Mac OS"))
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be Hello Eko")
	fmt.Println("Tes hello world with assert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		t.Error("Hasilnya harusnya Hello Eko")
	}
}