package stringutil

import "testing"

func TestReverse(t *testing.T) {
     cases := []struct {
     	   in, want string
	   }{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, äç", "çä ,olleH"},
		{"", ""},
	   }
	   for _, c := range cases {
	       got := Reverse(c.in)
	       if got != c.want {
	       	  t.Errorf("AGHHH")
	       }
	   }
}