package src

import (
	"testing"
)

func TestIsImageOrVideoFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"test.jpg"}, true},
		{"test2", args{"test.JPEG"}, true},
		{"test3", args{"test.jpeg"}, true},
		{"test4", args{"test.png"}, true},
		{"test5", args{"test.txt"}, false},
		{"test6", args{"test.gif"}, true},
		{"test7", args{"test.mp4"}, true},
		{"test8", args{"test"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImageOrVideoFile(tt.args.fileName); got != tt.want {
				t.Errorf("IsImageOrVideoFile() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestContains(t *testing.T) {
	type args struct {
		arr []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]string{"test1", "test2"}, "test1"}, true},
		{"test2", args{[]string{"test1", "test2"}, "test3"}, false},
		{"test3", args{[]string{"test1", "test2"}, "test2"}, true},
		{"test4", args{[]string{"test1", "test2"}, "test4"}, false},
		{"test5", args{[]string{"test1", "test2"}, "test5"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.arr, tt.args.str); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
