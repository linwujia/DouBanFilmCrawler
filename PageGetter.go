package main

import (
	"fmt"
	"io"
	"net/http"
)

type PageGetter struct {
	url string
	index uint
	manager *DouBanManager
}

func NewPageGetter(url string, index uint, manager *DouBanManager) (getter *PageGetter) {
	getter = new(PageGetter)
	getter.url = url
	getter.index = index
	getter.manager = manager
	return
}

func (g *PageGetter) Run() {
	resp, err := http.Get(g.url)
	if err != nil {
		fmt.Errorf("PageGetter Run get data error %e", err)
		return
	}

	defer resp.Body.Close()

	buff := make([]byte, 1024)
	var result string

	for {
		n, err1 := resp.Body.Read(buff)
		if n <= 0 {
			break
		}

		if err1 != nil && err1 != io.EOF {
			err = err1
			break
		}

		result += string(buff[:n])
	}

	if err != nil {
		fmt.Errorf("PageGetter Run read data error %e", err)
	}

	g.manager.SendPageData(&PageData{
		data:  result,
		index: g.index,
	})
}
