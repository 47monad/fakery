package fkinternet

import (
	"fmt"
	"strings"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkcompany"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkopts"
	"github.com/47monad/fakery/fk/fkperson"
	"github.com/47monad/fakery/internal/fkgeneral"
	"github.com/47monad/fakery/internal/sampler"
)

func ipv4Address() string {
	num := func() int { return fkgeneral.RandomInt(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

func ipv6Address() string {
	num := func() int { return fkgeneral.RandomInt(65536) }
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", num(), num(), num(), num(), num(), num(), num(), num())
}

func macAddress() string {
	num := func() int { return fkgeneral.RandomInt(255) }
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", num(), num(), num(), num(), num(), num())
}

func browser(ctx *fk.Context[fkdata.Internet]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Internet) []string { return d.Browser })
}

func browserEngine(ctx *fk.Context[fkdata.Internet]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.Internet) []string { return d.BrowserEngine })
}

func domainSuffix(ctx *fk.Context[fkdata.Internet], safe bool) string {
	keyer := func(d *fkdata.Internet) []string {
		if safe {
			return d.SafeDomainSuffix
		}
		return d.DomainSuffix
	}
	return sampler.Run(ctx.Data, keyer)
}

func domainName(ctx *fk.Context[fkdata.Internet]) string {
	companyName := strings.ReplaceAll(fkcompany.Buzzword(), " ", "")
	suffix := domainSuffix(ctx, true)
	return fmt.Sprintf("%s.%s", strings.ToLower(companyName), suffix)
}

func email(ctx *fk.Context[fkdata.Internet]) string {
	domain := domainName(ctx)
	firstName := fkperson.FirstName(fkopts.Name().SetLang("en").SetGender("male"))
	lastName := fkperson.LastName(fkopts.LastName().SetLang("en"))

	return fmt.Sprintf("%s.%s@%s", strings.ToLower(firstName), strings.ToLower(lastName), domain)
}

func username(ctx *fk.Context[fkdata.Internet]) string {
	lastName := fkperson.LastName(fkopts.LastName().SetLang("en"))

	return fmt.Sprintf("%s%d", strings.ToLower(lastName), fkgeneral.RandomInt(9999))
}
