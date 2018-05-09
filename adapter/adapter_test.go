package adapter

import (
	"testing"
)

func TestPrinterAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{OldPrinter: &MyLagacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()

	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()

	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

}
