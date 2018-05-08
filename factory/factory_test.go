package factory
import (
	"testing"
	"strings"
)
func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message not correct")
	}
	t.Log("Log:",msg)
}

func TestCreatePAymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("Debit payment method is not correct")
	}
	t.Log("Log:",msg)
}

func TestGetPaymentMethodDoesNotExistance(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 Must return error")
	}
	t.Log("Log:", err)
}