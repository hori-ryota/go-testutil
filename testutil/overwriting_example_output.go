package testutil

import (
	"io/ioutil"
	"os"
)

func OverwritingExampleOutputWrapper(example func(), overwritingFunc func(s []byte) []byte) {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w

	go func() {
		defer w.Close()
		example()
	}()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	b = overwritingFunc(b)
	_, err = stdout.Write(b)
	if err != nil {
		panic(err)
	}
	os.Stdout = stdout
}
