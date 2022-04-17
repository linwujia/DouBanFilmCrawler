package main

import (
	"fmt"
	"os"
	"sync"
)

type DouBanManager struct {

	pageGetters []*PageGetter
	pageParse []*PageParse
}

func NewDouBanManager(start, end uint) *DouBanManager {
	manager := &DouBanManager{
		pageGetters: make([]*PageGetter, end - start + 1),
		pageParse:   make([]*PageParse, end - start + 1),
	}

	for i := range manager.pageGetters {
		manager.pageGetters[i] = NewPageGetter(getPageUrl(start + uint(i)), start + uint(i),  manager)
	}

	for i := range manager.pageParse {
		manager.pageParse[i] = NewPageParse()
	}

	return manager
}

func (m DouBanManager) Run()  {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(len(m.pageParse))
	for _, parse := range m.pageParse {
		go func(parse *PageParse) {
			defer waitGroup.Done()
			parse.Run()
		}(parse)
	}

	waitGroup.Add(len(m.pageGetters))
	for _, getter := range m.pageGetters {
		go func(getter *PageGetter) {
			defer waitGroup.Done()
			getter.Run()
		}(getter)
	}

	waitGroup.Wait()
}

func (m *DouBanManager) SendPageData(data *PageData)  {
	file, err := os.Create(fmt.Sprintf("page%d.html", data.index))
	if err != nil {
		fmt.Println("create file error", err)
	}
	defer file.Close()
	file.WriteString(data.data)
}
