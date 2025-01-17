package fkinternet_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkinternet"
)

func TestDomainName(t *testing.T) {
	fmt.Println(fkinternet.Domain())
}
