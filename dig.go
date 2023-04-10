package main

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/urfave/cli/v2"
	"time"
)

func DnsQuery(domain, server string) (*dns.Msg, error) {

	c := dns.Client{Timeout: 5 * time.Second}

	m := dns.Msg{}
	m.SetQuestion(domain+".", dns.TypeA)

	r, _, err := c.Exchange(&m, server)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return nil, err
	}

	return r, nil
}

func calcTime(domain, server string) (*dns.Msg, time.Duration, error) {
	t := time.Now()

	r, err := DnsQuery(domain, server)
	if err != nil {
		return nil, 0, err
	}

	ms := time.Since(t)

	return r, ms, nil
}

func Dig(ctx *cli.Context) error {
	domain := ctx.String("domain")
	server := ctx.String("server")

	r, t, err := calcTime(domain, server)
	if err != nil {
		return fmt.Errorf("Timeout domain :%s, err:%s\n", domain, err.Error())
	}

	if len(r.Answer) <= 0 {
		return fmt.Errorf("No answer domain :%s\n", domain)
	}

	fmt.Println(r)
	fmt.Printf("dnsTools delay:%dms\n", t.Milliseconds())
	return nil
}
