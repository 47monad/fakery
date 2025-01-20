package fkinternet_test

import (
	"fmt"
	"testing"

	"github.com/47monad/fakery/fk/fkinternet"
)

func TestIPv4(t *testing.T) {
	fmt.Println(fkinternet.IPv4())
}

func TestIPv6(t *testing.T) {
	fmt.Println(fkinternet.IPv6())
}

func TestDomain(t *testing.T) {
	fmt.Println(fkinternet.Domain())
}

func TestMacAddr(t *testing.T) {
	fmt.Println(fkinternet.MacAddr())
}
