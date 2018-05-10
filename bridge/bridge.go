package bridge

import (
	"errors"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err errors) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = error.New("Content received on writer was empty")
}
