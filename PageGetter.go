package main

import (
	"github.com/golang/glog"
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

	request, err := http.NewRequest("GET", g.url, nil)
	//resp, err := http.Get(g.url)
	if err != nil {
		glog.Errorf("PageGetter Run new request error %e", err)
		return
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		glog.Errorf("PageGetter Run get data error %e", err)
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
		glog.Errorf("PageGetter Run read data error %e", err)
	}

	g.manager.SendPageData(&PageData{
		data:  result,
		index: g.index,
	})
}
