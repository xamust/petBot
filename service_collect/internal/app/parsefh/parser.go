package parsefh

import (
	"github.com/gocolly/colly"
	"strings"
)

type ParserFH struct {
	clubMap map[string]string
	c       *colly.Collector
}

var clubString = "https://www.fitnesshouse.ru/club.html"

func CollectorInit() *ParserFH {
	return &ParserFH{
		clubMap: make(map[string]string),
		c: colly.NewCollector(
			colly.Async(true),
			colly.AllowURLRevisit(),
			//colly.Debugger(&debug.LogDebugger{}),
			colly.IgnoreRobotsTxt(),
		),
	}
}

func (p *ParserFH) GetData() (map[string]string, error) {
	if err := p.collectFHClub(); err != nil {
		return nil, err
	}
	return p.clubMap, nil
}

func (p *ParserFH) collectFHClub() error {
	p.c.OnHTML(".tab-content", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			if element.ChildAttr("img", "alt") != "" {
				p.clubMap[strings.ToLower(element.ChildAttr("img", "alt"))] = element.Attr("href")
			}
		})
	})
	if err := p.c.Visit(clubString); err != nil {
		return err
	}
	p.c.Wait()
	return nil
}
