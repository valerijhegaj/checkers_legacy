package file

import "io"

func NewReadCloserFromString(data string) io.ReadCloser {
	return &TestStringReadCloser{data, 0}
}

func NewReadCloserFromBytes(data []byte) io.ReadCloser {
	return &BytesReadCloser{data, 0}
}

type TestStringReadCloser struct {
	data string
	ptr  int
}

func (c *TestStringReadCloser) Read(b []byte) (int, error) {
	if len(c.data) == 0 {
		return 0, io.EOF
	}
	n := copy(b, c.data)
	c.data = c.data[n:]
	return n, nil
}

func (c *TestStringReadCloser) Close() error {
	return nil
}

type BytesReadCloser struct {
	data []byte
	ptr  int
}

func (c *BytesReadCloser) Read(b []byte) (int, error) {
	if len(c.data) == 0 {
		return 0, io.EOF
	}
	n := copy(b, c.data)
	c.data = c.data[n:]
	return n, nil
}

func (c *BytesReadCloser) Close() error {
	return nil
}
