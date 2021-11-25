package greetings

import (
	"regexp"
	"testing"
)

func TestGreetName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)

	msg,err := Greet(name);

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T) {
	name := ""
	msg,err := Greet(name);

	if msg != "" || err == nil {
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}