package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// StatusOrdemDeCompra represents the order's status
type StatusOrdemDeCompra int

const (
	// StatusPendente represents a pending order
	StatusPendente StatusOrdemDeCompra = iota
	// StatusCancelado represents an canceled order
	StatusCancelado
	// StatusAprovado represents an approved order. Approved means that the manager has
	// approved the order but it isn't paid.
	StatusAprovado
	//StatusTerminado means that the order is successfully completed.
	StatusTerminado
)

// Item holds reference to the produto
type Item struct {
	ID              int     `json:"id" sql:"autoincrement"`
	ProdutoID       int     `json:"produto_id"`
	OrdemDeCompraID int     `sql:"index"`
	Quantidade      int     `json:"quantidade"`
	Valor           float64 `json:"valor"`
}

// OrdemDeCompra stores relevant data from a single order.
type OrdemDeCompra struct {
	Id         int                 `json:"ordem_id" sql:"autoincrement"`
	Status     StatusOrdemDeCompra `json:"status"`
	Items      []Item              `json:"itens"`
	CreatedAt  time.Time           `json:"criado_em"`
	FinishedAt time.Time           `json:"terminado_em"`
}

// OrdensDeCompra type is a slice of OrdemDeCompra
type OrdensDeCompra []OrdemDeCompra

// Approve approves a pending order
func (o *OrdemDeCompra) Approve() error {
	if o.Status != StatusPendente {
		return errors.New("This order cannot be approved.")
	}
	o.Status = StatusAprovado
	db := OpenDB()
	db.Save(o)
	return nil
}

// Cancel only cancels a not finished order
func (o *OrdemDeCompra) Cancel() error {
	if o.Status == StatusTerminado {
		return errors.New("This order is already finished.")
	}
	o.Status = StatusCancelado
	db := OpenDB()
	db.Save(o)
	return nil
}

// GetByID returns the correspondent OrdemDeCompra for the given ID
func (o *OrdemDeCompra) GetByID(id int) {
	db := OpenDB()
	db.First(o, id)
}

// Finish finishes an already approved order
func (o *OrdemDeCompra) Finish() error {
	if o.Status != StatusAprovado {
		return errors.New("This order must be approved before finishing.")
	}
	o.Status = StatusTerminado
	db := OpenDB()
	db.Save(o)
	return nil
}

// AddItem appends a new items to the exitent items
func (o *OrdemDeCompra) AddItem(item Item) {
	item.OrdemDeCompraID = o.Id
	db := OpenDB()
	db.Create(&item)
	o.Items = append(o.Items, item)
}

// SetItems overrides the default items with the items passed
func (o *OrdemDeCompra) SetItems(items []Item) {
	o.Items = items
	db := OpenDB()
	db.Save(o)
}

// Create persists a new OrdemDeCompra
func (o *OrdemDeCompra) Create() {
	db := OpenDB()
	db.Debug().Create(o)
}

// GetByStatus returns all OrdensDeCompra with the given status
func (os *OrdensDeCompra) GetByStatus(s StatusOrdemDeCompra) {
	db := OpenDB()
	db.Where("status = ?", s).Find(os)
}

// NewOrdemDeCompra returns a new instance of OrdemDeCompra from an array of Items
// returns nil if there was a problem
func NewOrdemDeCompra(items []Item) *OrdemDeCompra {
	o := &OrdemDeCompra{
		Status: StatusPendente,
		Items:  items,
	}
	db := OpenDB()
	if db.NewRecord(o) {
		return o
	}
	return nil
}

func OpenDB() gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/compras.db")
	if err != nil {
		panic(err)
	}
	db.DB()
	return db
}
