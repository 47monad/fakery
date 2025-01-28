package fkphone_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/fk/fkphone"
)

func TestNumber(t *testing.T) {
	fmt.Println(fkphone.Number())
	fmt.Println(fkphone.MobileNumber())
	fmt.Println(fkphone.Number(fkopts.PhoneNumber().SetLang("fa")))
	fmt.Println(fkphone.MobileNumber(fkopts.MobileNumber().SetLang("fa")))
}

func TestE164(t *testing.T) {
	fmt.Println(fkphone.E164())
	fmt.Println(fkphone.E164(fkopts.MobileNumber().SetLang("fa")))
}
