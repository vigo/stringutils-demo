package stringutils_test

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/vigo/stringutils-demo"
)

func FuzzReverse(f *testing.F) {
	tcs := []string{"hello", "uğur", "Präzisionsmeßgerät"}
	for _, tc := range tcs {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, original string) {
		reversed, err1 := stringutils.Reverse(original)
		if err1 != nil {
			return
		}

		reversedAgain, err2 := stringutils.Reverse(reversed)
		if err2 != nil {
			return
		}

		if original != reversedAgain {
			t.Errorf("want: %q got: %q", reversed, reversedAgain)
		}
		if utf8.ValidString(original) && !utf8.ValidString(reversed) {
			t.Errorf("reverse produced invalid UTF-8 string %q", reversed)
		}
	})
}

func TestReverse(t *testing.T) {
	tcs := map[string]struct {
		input []string
		want  []string
		err   []error
	}{
		"none Turkish letters": {
			input: []string{"hello", "this is vigo"},
			want:  []string{"olleh", "ogiv si siht"},
			err:   []error{nil, nil},
		},
		"with Turkish letters": {
			input: []string{"uğur", "kırmızı şapka ve ÖĞRENCİ"},
			want:  []string{"ruğu", "İCNERĞÖ ev akpaş ızımrık"},
			err:   []error{nil, nil},
		},
		"with German letters": {
			input: []string{"Präzisionsmeßgerät"},
			want:  []string{"täregßemsnoisizärP"},
			err:   []error{nil, nil},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			for i, in := range tc.input {
				got, err := stringutils.Reverse(in)

				if !errors.Is(err, tc.err[i]) {
					t.Errorf("want: %v; got: %v", tc.err[i], err)
				}

				if got != tc.want[i] {
					t.Errorf("want: %v; got: %v", tc.want[i], got)
				}
			}
		})
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s, _ = stringutils.Reverse("merhaba dünya!")
	}

	gs = s
}

func ExampleReverse() {
	r, _ := stringutils.Reverse("vigo")
	fmt.Println(r)
	// Output: ogiv
}
