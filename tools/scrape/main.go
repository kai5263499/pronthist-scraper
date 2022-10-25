package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/caarlos0/env/v6"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
	"github.com/headzoo/surf/jar"
	"github.com/sirupsen/logrus"
)

type config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

var (
	cfg config
)

func main() {
	cfg = config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithError(err).Fatal("parse config")
	}

	if level, err := logrus.ParseLevel(cfg.LogLevel); err != nil {
		logrus.WithError(err).Fatal("parse level")
	} else {
		logrus.SetLevel(level)
	}

	for char := 97; char <= 122; char++ {
		uri := fmt.Sprintf("https://phrontistery.info/%c.html", char)
		if err := extractWords(uri); err != nil {
			logrus.WithError(err).Errorf("error extracting words from %s", uri)
			continue
		}
	}
}

func extractWords(uri string) error {
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Chrome())
	bow.SetAttribute(browser.SendReferer, true)
	bow.SetAttribute(browser.MetaRefreshHandling, true)
	bow.SetAttribute(browser.FollowRedirects, true)
	bow.SetCookieJar(jar.NewMemoryCookies())
	if err := bow.Open(uri); err != nil {
		return err
	}

	bow.Find("table.words tr").Each(func(trIdx int, trSel *goquery.Selection) {
		if trIdx < 1 {
			return
		}

		trSel.Find("td").Each(func(tdIdx int, tdSel *goquery.Selection) {
			fmt.Printf(tdSel.Text())
			if tdIdx == 0 {
				fmt.Printf("\t")
			}
		})
	})

	return nil
}
