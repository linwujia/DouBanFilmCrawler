package main

import (
	"strconv"
)

var PageUrlBase = "https://movie.douban.com/top250?start=%s&filter="

func getPageUrl(page uint) string {
	return "https://movie.douban.com/top250?start=" + strconv.Itoa(int((page-1)*25)) + "&filter="
}
