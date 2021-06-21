package main

import (
	"bufio"
	"io"
)

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func run_Write(fd io.Writer, p0, p1, p2 []byte, a, b, c, d, e, f int) error {
	ew := &errWriter{w: fd}
	ew.write(p0[a:b])
	ew.write(p1[c:d])
	ew.write(p2[e:f])
	// and so on
	if ew.err != nil {
		return ew.err
	}

	return nil
}

func run_bufioWrite(fd io.Writer, p0, p1, p2 []byte, a, b, c, d, e, f int) error {
	buf := bufio.NewWriter(fd)
	buf.Write(p0[a:b])
	buf.Write(p1[c:d])
	buf.Write(p2[e:f])
	// and so on
	if buf.Flush() != nil {
		return buf.Flush()
	}

	return nil
}

func main() {
}

// https://blog.golang.org/errors-are-values
