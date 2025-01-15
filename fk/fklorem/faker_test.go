package fklorem_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fklorem"
	"github.com/47monad/fakery/fk/fkopts"
)

func TestWord(t *testing.T) {
	faker := fklorem.New()
	fmt.Println(faker.Word())
	faker2 := fklorem.New(fkopts.Faker().SetLang("fa"))
	fmt.Println(faker2.Word())
	fmt.Println(fklorem.Word())
	fmt.Println(fklorem.Word(fkopts.Word().SetLang("fa")))
}
