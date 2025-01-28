package fkphone

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/47monad/fakery/fk"
	"github.com/47monad/fakery/fk/fkdata"
	"github.com/47monad/fakery/fk/fkrand"
	"github.com/47monad/fakery/internal/sampler"
	"github.com/47monad/fakery/internal/templater"
)

func areaCode(ctx *fk.Context[fkdata.PhoneNumber]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.PhoneNumber) []string { return d.AreaCode })
}

func countryCode(ctx *fk.Context[fkdata.PhoneNumber]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.PhoneNumber) []string { return d.CountryCode })
}

func exchangeCode(ctx *fk.Context[fkdata.PhoneNumber]) string {
	return sampler.Run(ctx.Data, func(d *fkdata.PhoneNumber) []string { return d.ExchangeCode })
}

func extension(ctx *fk.Context[fkdata.PhoneNumber], length uint32) string {
	return fmt.Sprintf("%d", fkrand.Int(length))
}

func number(ctx *fk.Context[fkdata.PhoneNumber]) string {
	tmpl := sampler.Run(ctx.Data, func(d *fkdata.PhoneNumber) []string { return d.Format })
	return templater.MustExec("phoneNumber", tmpl, templater.MergeFuncMaps(TemplateFuncMap(ctx), fkrand.TemplateFuncMap()))
}

func mobileNumber(ctx *fk.Context[fkdata.PhoneNumber]) string {
	tmpl := sampler.Run(ctx.Data, func(d *fkdata.PhoneNumber) []string { return d.MobileFormat })
	return templater.MustExec("mobileNumber", tmpl, templater.MergeFuncMaps(TemplateFuncMap(ctx), fkrand.TemplateFuncMap()))
}

func e164(ctx *fk.Context[fkdata.PhoneNumber]) string {
	re := regexp.MustCompile(`[0-9]+`)
	number := strings.Join(re.FindAllString(mobileNumber(ctx), -1), "")
	return fmt.Sprintf("+%s%s", countryCode(ctx), number)
}
