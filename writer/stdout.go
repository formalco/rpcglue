package writer

import (
	"github.com/formalco/rpcglue/log"
)

type StdoutWriter struct{}

func NewStdoutWriter() *StdoutWriter {
	return &StdoutWriter{}
}

func (s *StdoutWriter) Write(_path string, data []byte) error {
	log.Print(string(data))
	return nil
}
