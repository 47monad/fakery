package fklorem_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fklorem"
	"github.com/47monad/fakery/fk/fkopts"
)

func TestSolo(t *testing.T) {
	fmt.Println(fklorem.Words())
	fmt.Println(fklorem.Words(fkopts.Words().SetLang("fa").SetCount(10)))
	fmt.Println(fklorem.Sentence())
	fmt.Println(fklorem.Sentence(fkopts.Sentence().SetLang("fa").SetWordCount(10)))
}
