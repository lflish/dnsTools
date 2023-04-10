package main

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/urfave/cli/v2"
	"time"
)

func show(msg *dns.Msg, t time.Duration) {
	for _, v := range msg.Answer {
		record, ok := v.(*dns.A)
		if ok {
			fmt.Printf("domain:%s, server:%s, delay:%dms\n", msg.Question[0].Name, record.A, t.Milliseconds())
			break
		}
	}
}

func Speed(ctx *cli.Context) error {
	domain := ctx.String("domain")
	count := ctx.Int("count")
	server := ctx.String("server")

	every := ctx.Int("every")

	var sucTime, sunCount int = 0, 0

	for i := 0; i < count; i++ {
		r, t, err := calcTime(domain, server)
		time.Sleep(time.Duration(every) * time.Millisecond)

		if err != nil {
			fmt.Printf("Timeout:domain:%s\n", domain)
			continue
		}

		if len(r.Answer) <= 0 {
			continue
		}

		sucTime += int(t.Milliseconds())
		sunCount += 1

		show(r, t)
	}

	if sunCount > 0 {
		fmt.Printf("dnsTools Total:%d, Succ:%d, Error:%d, averageTime:%dms\n", count, sunCount, count-sunCount, sucTime/sunCount)
	} else {
		fmt.Printf("dnsTools Total:%d, Succ:%d, Error:%d\n", count, sunCount, count-sunCount)
	}
	return nil
}
