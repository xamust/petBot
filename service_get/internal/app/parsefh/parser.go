package parsefh

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type ParserFH struct {
	resultMap map[string][]string
	c         *colly.Collector
	data      Trainees
}

type Storage struct {
	DayOfWeek string
	Trainee   []string
}

type Trainees struct {
	Store []Storage
}

var hrefPref = "https://www.fitnesshouse.ru/%s"

func CollectorInit() *ParserFH {
	return &ParserFH{
		resultMap: make(map[string][]string),
		c: colly.NewCollector(
			colly.Async(true),
			colly.AllowURLRevisit(),
			//colly.Debugger(&debug.LogDebugger{}),
			colly.IgnoreRobotsTxt(),
		),
	}
}

func (p *ParserFH) Search(clubUrl string) (map[string][]string, error) {
	pD, err := p.parsing(clubUrl)
	if err != nil {
		return nil, err
	}
	p.CollectStruct(pD)
	return p.resultMap, nil
}

func (p *ParserFH) parsing(clubHref string) (map[int][]string, error) {

	parseData := make(map[int][]string)

	count := 1
	p.c.OnHTML(".shedule", func(e *colly.HTMLElement) {
		e.ForEach(".shedule th", func(i int, element *colly.HTMLElement) {
			parseData[i] = append(parseData[i], element.Text)
		})
		e.ForEach(".shedule td", func(i int, element *colly.HTMLElement) {
			if element.Attr("rowspan") != "" {
			} else {
				if count > 7 {
					count = 1
				}
				if strings.TrimSpace(element.Text)[5:] != "" {
					parseData[count] = append(parseData[count], strings.TrimSpace(element.Text), "\n")
				}
				count++
			}
		})
	})

	if err := p.c.Visit(fmt.Sprintf(hrefPref, clubHref)); err != nil {
		return nil, err
	}
	p.c.Wait()
	return parseData, nil
}

func (p *ParserFH) CollectStruct(inputMap map[int][]string) {
	for i, i2 := range inputMap {
		switch i {
		case 1:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "понедельник",
				Trainee:   i2,
			})
		case 2:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "вторник",
				Trainee:   i2,
			})
		case 3:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "среда",
				Trainee:   i2,
			})
		case 4:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "четверг",
				Trainee:   i2,
			})
		case 5:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "пятница",
				Trainee:   i2,
			})
		case 6:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "суббота",
				Trainee:   i2,
			})
		case 7:
			p.data.Store = append(p.data.Store, Storage{
				DayOfWeek: "воскресенье",
				Trainee:   i2,
			})
		}
	}
}

func (p *ParserFH) GetData(day string) []string {
	for _, storage := range p.data.Store {
		if storage.DayOfWeek == strings.ToLower(day) {
			return storage.Trainee
		}
	}
	return nil
}

func trimSpaces(input string) string {
	sb := strings.Builder{}
	for _, b := range strings.Split(input, "\n") {
		if b != "" {
			temp := string(bytes.TrimSpace([]byte(b)))
			if temp != "" {
				sb.WriteString(temp)
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}
