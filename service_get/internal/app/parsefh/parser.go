package parsefh

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type ParserFH struct {
	resultMap map[int][]string
	c         *colly.Collector
}

var hrefPref = "https://www.fitnesshouse.ru/%s"

func CollectorInit() *ParserFH {
	return &ParserFH{
		resultMap: make(map[int][]string),
		c: colly.NewCollector(
			colly.Async(true),
			colly.AllowURLRevisit(),
			//colly.Debugger(&debug.LogDebugger{}),
			colly.IgnoreRobotsTxt(),
		),
	}
}

func (p *ParserFH) GetData(clubUrl string) (map[int][]string, error) {
	if err := p.parsing(clubUrl); err != nil {
		return nil, err
	}
	return p.resultMap, nil
}

func (p *ParserFH) parsing(clubHref string) error {

	count := 1
	p.c.OnHTML(".shedule", func(e *colly.HTMLElement) {
		e.ForEach(".shedule th", func(i int, element *colly.HTMLElement) {
			p.resultMap[i] = append(p.resultMap[i], element.Text)
		})
		e.ForEach(".shedule td", func(i int, element *colly.HTMLElement) {
			if element.Attr("rowspan") != "" {
			} else {
				if count > 7 {
					count = 1
				}
				if strings.TrimSpace(element.Text)[5:] != "" {
					p.resultMap[count] = append(p.resultMap[count], strings.TrimSpace(element.Text), "\n")
				}
				count++
			}
		})
	})

	if err := p.c.Visit(fmt.Sprintf(hrefPref, clubHref)); err != nil {
		return err
	}
	p.c.Wait()
	return nil
}
