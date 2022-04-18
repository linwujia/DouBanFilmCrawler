package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type PageParse struct {
	dataCh chan *PageData
}

func NewPageParse() *PageParse {
	parse := new(PageParse)
	parse.dataCh = make(chan *PageData)
	return parse
}

func (p *PageParse) Run()  {
	select {
		case data := <- p.dataCh:
			p.parsePageData(data)
	}
}

func (p *PageParse) SendPageData(data *PageData)  {
	p.dataCh <- data
}

func (p *PageParse) parsePageData(data *PageData)  {
	filmNameReg := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	filmNames := filmNameReg.FindAllStringSubmatch(data.data, -1)
	fmt.Println(filmNames)

	filmScoreReg := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*)</span>`)
	filmScores := filmScoreReg.FindAllStringSubmatch(data.data, -1)
	fmt.Println(filmScores)

	filmScoreNumReg := regexp.MustCompile(`<span>(.*)人评价</span>`)
	filmScoreNum := filmScoreNumReg.FindAllStringSubmatch(data.data, -1)
	fmt.Println(filmScoreNum)

	films := make([]FilmData, len(filmNames))
	for i := 0; i < len(films); i++ {
		films[i] = FilmData{
			Name:     filmNames[i][1],
			Score:    filmScores[i][1],
			ScoreNum: filmScoreNum[i][1],
		}
	}

	err := p.save2File(films, data.index)
	if err != nil {
		fmt.Errorf("%d page data save to File error %e", data.index, err)
	}
}

func (p *PageParse) save2File(films []FilmData, page uint) error {
	f, err := os.Create(fmt.Sprintf("page %d.json", page))
	if err != nil {
		return err
	}

	defer f.Close()

	bytes, err := json.MarshalIndent(films, "", "\t")
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(bytes))
	if err != nil {
		return err
	}

	return nil
}
