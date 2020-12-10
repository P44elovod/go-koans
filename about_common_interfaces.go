package go_koans

import (
	"bytes"
	"io"
	"log"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		if _, err := io.Copy(out, bytes.NewReader(in.Bytes())); err != nil {
			log.Fatal(err)
		}

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		if _, err := io.CopyN(out, bytes.NewReader(in.Bytes()), 5); err != nil {
			log.Fatal(err)
		}

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
