package pago_test

import (
	"testing"
	"github.com/raninho/pago"
	"log"
)

type myContent []int

func (m myContent) NumberPage() int {return 0}
func (m myContent) LimitPage() int {return 4}
func (m myContent) TotalElements() int {return len(m)}
func (m myContent) Content() interface{} {return m}

func TestMakePagerWithSuccess(t *testing.T) {
	myContentPage :=  myContent{10, 1, 2, 3}
	page, err := pago.BuildPage(myContentPage)
	if err != nil {
		t.Errorf("Error make build page")
	}
	if page.TotalElements != 4 {
		t.Errorf("TotalElements != 4")
	}
	log.Println(myContentPage.Content())
	log.Println(page.ToJSON())
}
