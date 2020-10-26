package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"proxy-collect/component"
	"proxy-collect/component/logger"
	"proxy-collect/config"
	"proxy-collect/dto"
	"strings"
)

func NewGetProxyIp3366() *getProxyIp3366 {
	return &getProxyIp3366{}
}

type getProxyIp3366 struct {
}

func (s *getProxyIp3366) GetUrlList() []string {
	list := []string{
		"http://www.ip3366.net/free/?stype=1",
		"http://www.ip3366.net/free/?stype=2",
	}
	for i := 2; i < 6; i++ {
		list = append(list, fmt.Sprintf("http://www.ip3366.net/free/?stype=1&page=%d", i))
		list = append(list, fmt.Sprintf("http://www.ip3366.net/free/?stype=2&page=%d", i))
	}
	return list
}
func (s *getProxyIp3366) GetContentHtml(requestUrl string) string {
	h := dto.RequestHeaderDto{
		UserAgent:               config.USER_AGENT,
		UpgradeInsecureRequests: "1",
	}

	logger.Info("get proxy from ip3366", requestUrl)
	return component.WebGet(requestUrl, h)
}

func (s *getProxyIp3366) ParseHtml(body string) [][]string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		logger.Error(err)
		return nil
	}
	var proxyList [][]string
	doc.Find("tbody > tr").Each(func(i int, selection *goquery.Selection) {
		td := selection.ChildrenFiltered("td").First()
		host := strings.TrimSpace(td.Text())
		td2 := selection.ChildrenFiltered("td").Eq(1)
		port := strings.TrimSpace(td2.Text())

		if !ProxyService.CheckProxyFormat(host, port) {
			logger.Error("格式错误:", host, port)
			return
		}
		proxyArr := []string{host, port}
		proxyList = append(proxyList, proxyArr)
	})
	return proxyList
}
