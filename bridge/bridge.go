package bridge

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI TODO
type PrinterAPI interface {
	PrintMessage(string) error
}

// PrinterImpl1 TODO
type PrinterImpl1 struct{}

// PrintMessage TODO
func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// PrinterImpl2 TODO
type PrinterImpl2 struct {
	Writer io.Writer
}

// PrintMessage TODO
func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

// TestWriter TODO
type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on writer was empty")
	return
}

// PrinterAbstraction TODO
type PrinterAbstraction interface {
	Print() error
}

// NormalPrinter TODO
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print TODO
func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// PacktPrinter TODO
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print TODO
func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil

}
