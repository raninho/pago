package pago

import (
	"bytes"
	"encoding/json"
)

type ContentPager interface {
	NumberPage() int
	LimitPage() int
	TotalElements() int
	Content() interface{}
}

type Page struct {
	Number           int         `json:"number"`
	TotalPages       int         `json:"totalPages"`
	NumberOfElements int         `json:"numberOfElements"`
	TotalElements    int         `json:"totalElements"`
	HasPreviousPage  bool        `json:"hasPreviousPage"`
	HasNextPage      bool        `json:"hasNextPage"`
	HasContent       bool        `json:"hasContent"`
	First            bool         `json:"first"`
	Last             bool        `json:"last"`
	NextPage         int         `json:"nextPage"`
	PreviousPage     int         `json:"previousPage"`
	Content          interface{} `json:"content"`
}

func (p Page) ToJSON() (*bytes.Buffer, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(p); err != nil {
		return nil, err
	}

	return &buf, nil
}

//BuildPage ...
func BuildPage(c ContentPager) (Page, error) {
	page := Page{}
	page.Number = c.NumberPage()
	page.TotalPages = c.TotalElements() / c.LimitPage()
	page.NumberOfElements = c.LimitPage()
	page.TotalElements = c.TotalElements()
	page.HasPreviousPage = c.NumberPage() > 0
	page.HasNextPage = c.NumberPage() < (page.TotalPages - 1)
	page.HasContent = c.TotalElements() > 0
	page.First = c.NumberPage() == 0
	page.Last = c.NumberPage() == page.TotalPages - 1

	page.NextPage = c.NumberPage() + 1
	if c.NumberPage() == page.TotalPages - 1 {
		page.NextPage = c.NumberPage()
	}

	page.PreviousPage = c.NumberPage() -1
	if c.NumberPage() == 0 {
		page.PreviousPage = 0
	}

	page.Content = c.Content()

	return page, nil
}

/*
   @Override
   public <T> Page<T> buildPage(List<T> list, Integer page, Integer limit, Long totalElements) {

        Page<T> pageResponse = new Page<>();

        pageResponse.number = page;
        pageResponse.totalPages = new BigDecimal(totalElements).divide(new BigDecimal(limit), BigDecimal.ROUND_UP, 0).intValue();
        pageResponse.numberOfElements = limit;
        pageResponse.totalElements = totalElements;
        pageResponse.hasPreviousPage = page > 0;
        pageResponse.hasNextPage = page < (pageResponse.totalPages - 1);
        pageResponse.hasContent = Objeto.notBlank(list);
        pageResponse.first = page == 0;
        pageResponse.last = page == (pageResponse.totalPages - 1);
        pageResponse.nextPage = page == (pageResponse.totalPages - 1) ? page : page + 1;
        pageResponse.previousPage = page == 0 ? 0 : page - 1;
        pageResponse.content = list;

        return pageResponse;
   }
*/
