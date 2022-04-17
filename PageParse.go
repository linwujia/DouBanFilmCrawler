package main

type PageParse struct {
	dataCh chan []byte
}

func NewPageParse() *PageParse {
	parse := new(PageParse)
	parse.dataCh = make(chan []byte)
	return parse
}

func (p PageParse) Run()  {
	
}
