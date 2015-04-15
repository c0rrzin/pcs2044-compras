package server_test

import (
	// "net/http/httptest"
	"fmt"
	// "io"
	//"io/ioutil"
	"encoding/json"
	"testing"

	compras "github.com/c0rrzin/pcs2044-compras"
	"github.com/rcmgleite/labSoft2_Estoque/requestHelper"
	. "gopkg.in/check.v1"
)

const RemoteAddr = "http://localhost:8080"

//helpers
func getJSON(object interface{}) ([]byte, error) {
	if object != nil {
		return json.Marshal(object)
	}
	return nil, nil
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestNotFound(c *C) {
	resp, err := requestHelper.MakeRequest("GET", RemoteAddr+"/invalid", nil, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	//body, _ := ioutil.ReadAll(resp.Body)
	c.Assert(resp.StatusCode, Equals, 404)
}

func (s *MySuite) TestNewOrder(c *C) {
	var o compras.OrdemDeCompra
	o.Status = compras.StatusPendente
	o.Items = []compras.Item{compras.Item{1, 2, 1, 2, 1.2}, compras.Item{2, 2, 1, 1, 2.50}}
	jsonBody, _ := getJSON(o)
	resp, err := requestHelper.MakeRequest("POST", RemoteAddr+"/order", jsonBody, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	c.Assert(resp.StatusCode, Equals, 200)
}

type orderRequest struct {
	OrdemId int `json:"ordem_id"`
}

func (s *MySuite) TestApproveOrder(c *C) {
	var o orderRequest
	o.OrdemId = 1
	jsonBody, _ := getJSON(o)
	resp, err := requestHelper.MakeRequest("POST", RemoteAddr+"/order/approve", jsonBody, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *MySuite) TestListOrders(c *C) {
	resp, err := requestHelper.MakeRequest("GET", RemoteAddr+"/orders", nil, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *MySuite) TestFinishOrder(c *C) {
	var o orderRequest
	o.OrdemId = 1
	jsonBody, _ := getJSON(o)
	resp, err := requestHelper.MakeRequest("POST", RemoteAddr+"/order/finish", jsonBody, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *MySuite) TestCancelOrder(c *C) {
	var o orderRequest
	o.OrdemId = 1
	jsonBody, _ := getJSON(o)
	resp, err := requestHelper.MakeRequest("POST", RemoteAddr+"/order/cancel", jsonBody, nil)
	if err != nil {
		fmt.Println("Ocorreu erro com o teste, verifique se o servidor foi inicializado com sucesso")
		c.Fail()
	}
	c.Assert(resp.StatusCode, Equals, 200)
}
