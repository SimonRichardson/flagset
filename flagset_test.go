package flagset

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"testing/quick"
)

func TestReadingFromEnv(t *testing.T) {
	fn := func(envValue, defaultValue, cmdValue ASCII) bool {
		os.Setenv("TEST", envValue.String())

		flagset := NewFlagSet("test", flag.ExitOnError)

		test := flagset.String("test", defaultValue.String(), "test value")

		args := []string{fmt.Sprintf("-test=%s", cmdValue.String())}
		if err := flagset.Parse(args); err != nil {
			t.Fatal(err)
		}

		return *test == envValue.String()
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func TestEnvName(t *testing.T) {
	for _, testcase := range []struct {
		value string
		want  string
	}{
		{"name", "NAME"},
		{"name.subname", "NAME_SUBNAME"},
		{"name..SubName", "NAME__SUBNAME"},
		{".NAmE.", "_NAME_"},
	} {
		t.Run(testcase.value, func(t *testing.T) {
			if expected, actual := testcase.want, envName(testcase.value); expected != actual {
				t.Errorf("expected: %s, actual: %s", expected, actual)
			}
		})
	}
}

func TestBuildFlag(t *testing.T) {
	for _, testcase := range []struct {
		title       string
		name, value string
		want        string
	}{
		{"normal", "a", "value", "-a=value"},
		{"no value", "a", "", "-a"},
		{"number value", "a", "'1'", "-a='1'"},
		{"bool value", "a", "true", "-a=true"},
		{"quoted value", "a", "\"value\"", "-a=\"value\""},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			if expected, actual := testcase.want, buildFlag(testcase.name, testcase.value); expected != actual {
				t.Errorf("expected: %s, actual: %s", expected, actual)
			}
		})
	}
}

// ASCII creates a series of tags that are ascii compliant.
type ASCII []byte

// Generate allows ASCII to be used within quickcheck scenarios.
func (ASCII) Generate(r *rand.Rand, size int) reflect.Value {
	var (
		chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		res   = make([]byte, size+1)
	)

	for k := range res {
		res[k] = byte(chars[r.Intn(len(chars)-1)])
	}

	return reflect.ValueOf(res)
}

func (a ASCII) Slice() []byte {
	return a
}

func (a ASCII) String() string {
	return string(a)
}
