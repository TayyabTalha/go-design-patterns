package adapter

import (
	"fmt"
)

// LagacyPrinter Interface
type LagacyPrinter interface {
	Print(s string) string
}

// MyLagacyPrinter implenents LagacyPrinter interface
type MyLagacyPrinter struct{}

// Print implement of LagacyPrinter
func (m *MyLagacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

// ModrenPrinter Interface
type ModrenPrinter interface {
	PrintStored() string
}

// PrinterAdapter will make LagacyPrinter Compatible
type PrinterAdapter struct {
	OldPrinter LagacyPrinter
	Msg        string
}

// PrintStored implementation of ModrenPrinter
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
