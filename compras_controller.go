package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// GetOrdersHandler returns all orders
// REQUEST:
// {
//   "order_id": 1
// }
// RESPONSE:
// { "orders": [
//  {
//   "status": "pendente",
//   "produtos": [
//     {
//        "produto_id": 1,
//        "quantidade": 2
//     },
//     {
//        "produto_id": 2,
//        "quantidade": 1
//     }
//
//    ]
//   }
//  ]
// }
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}
	os := &OrdensDeCompra{}
	os.All()
	fmt.Println(os)
	body, _ := json.Marshal(os)
	w.Write(body)
}

// GetOrderHandler returns a correspondent order from an "order_id"
// REQUEST:
// {
//   "order_id": 1
// }
// RESPONSE:
// {
//   "status": "pendente",
//   "produtos": [
//     {
//        "produto_id": 1,
//        "quantidade": 2
//     },
//     {
//        "produto_id": 2,
//        "quantidade": 1
//     }
//   ]
// }
func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}
	orderID, err := strconv.Atoi(r.URL.Query().Get("order_id"))
	if err != nil {
		http.Error(w, "Invalid order_id", 400)
		return
	}
	o := &OrdemDeCompra{}
	o.GetByID(orderID)
	body, _ := json.Marshal(o)
	w.Write(body)
}

type productRequest struct {
	ProdutoID  int     `json:"produto_id"`
	Quantidade int     `json:"quantidade"`
	Valor      float64 `json:"valor"`
}

type productsRequest struct {
	Produtos []productRequest `json:"produtos"`
}

type newOrderResponse struct {
	OrdemID int `json:"ordem_id"`
}

// NewOrderHandler creates a new order from the params expected below:
// REQUEST:
// {
//   "produtos": [
//     {
//        "produto_id": 1,
//        "valor": 1.24,
//        "quantidade": 2
//     },
//     {
//        "produto_id": 2,
//        "valor": 12.04,
//        "quantidade": 1
//     }
//   ]
// }
// RESPONSE:
// {
//   "order_id": 1
// }
func NewOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	var products productsRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&products)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	o := &OrdemDeCompra{}
	o.Status = StatusPendente
	o.CreatedAt = time.Now()
	var items []Item
	for _, p := range products.Produtos {
		items = append(items, Item{ProdutoId: p.ProdutoID, Quantidade: p.Quantidade, Valor: p.Valor})
	}

	o.Items = items
	o.Create()
	id := newOrderResponse{OrdemID: o.Id}
	respSer, _ := json.Marshal(id)
	w.Write(respSer)
}

type alterOrderRequest struct {
	OrdemID int `json:"ordem_id"`
}

// FinishOrderHandler finishes an order
// REQUEST:
// {
//   "order_id": 1
// }
// RESPONSE:
// {
//   "message": "Order finished successfully"
// }
func FinishOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	var order alterOrderRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	o := &OrdemDeCompra{}
	o.GetByID(order.OrdemID)
	if err := o.Finish(); err != nil {
		http.Error(w, err.Error(), 404)
	}
}

// ApproveOrderHandler approves an order
// REQUEST:
// {
//   "order_id": 1
// }
// RESPONSE:
// {
//   "message": "Order approved successfully"
// }
func ApproveOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	var order alterOrderRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	o := &OrdemDeCompra{}
	o.GetByID(order.OrdemID)
	if err := o.Approve(); err != nil {
		http.Error(w, err.Error(), 404)
	}
}

// CancelOrderHandler calcels an order
// REQUEST:
// {
//   "order_id": 1
// }
// RESPONSE:
// {
//   "message": "Order canceled successfully"
// }
func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "X-PINGOTHER, Origin, X-Requested-With, Content-Type, Accept")
	if r.ParseForm() != nil {
		http.Error(w, "Invalid parameters", 400)
		return
	}

	var order alterOrderRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	o := &OrdemDeCompra{}
	o.GetByID(order.OrdemID)
	if err := o.Cancel(); err != nil {
		http.Error(w, err.Error(), 404)
	}
}
