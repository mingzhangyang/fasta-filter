package lib

import "io"

// Record represent one sequence
type Record struct {
	Head, Body []byte
}

// NewRecord create an new record
func NewRecord() *Record {
	return &Record{
		Head: make([]byte, 0, 512),
		Body: make([]byte, 0, 1024 * 16),
	}
}

// LoadHead load a header line
func (r *Record) LoadHead(line []byte) {
	r.Head = append(r.Head, line...)
}

// LoadBody load a body line
func (r *Record) LoadBody(line []byte) {
	r.Body = append(r.Body, line...)
}

// WriteTo a io.Writer
func (r *Record) WriteTo(out io.Writer) (int64, error) {
	if len(r.Head) == 0 {
		return 0, nil
	}
	n, err := out.Write(r.Head)
	if err != nil {
		r.Reset()
		return 0, err
	}
	i, err := out.Write([]byte{'\n'})
	if err != nil {
		r.Reset()
		return 0, err
	}
	n += i
	i, err = out.Write(r.Body)
	if err != nil {
		r.Reset()
		return 0, err
	}
	n += i
	i, err = out.Write([]byte{'\n'})
	if err != nil {
		r.Reset()
		return 0, err
	}
	n += i
	r.Reset()
	return int64(n), nil
}

// Reset a record to reuse it
func (r *Record) Reset() {
	r.Head = r.Head[:0]
	r.Body = r.Body[:0]
}