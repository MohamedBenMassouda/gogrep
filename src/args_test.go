package src

//
// import (
// 	"os"
// 	"testing"
// )
//
// type Args struct {
// 	input string
// 	path  string
// }
//
// func TestCheckArgs(t *testing.T) {
// 	tests := []Args{
// 		{input: "test", path: "test"},
// 		{input: "input", path: "main.go"},
// 	}
//
// 	for _, test := range tests {
// 		os.Args = []string{test.input, test.path}
// 		input, path := CheckArgs()
//
// 		if input != test.input {
// 			t.Errorf("Expected %s, got %s", test.input, input)
// 		}
//
// 		if path != test.path {
// 			t.Errorf("Expected %s, got %s", test.path, path)
// 		}
//
// 	}
//
// }
