package model

import (
	"time"
)

// CompraCliente struct
type CompraCliente struct {
	Cpf                  string
	Private              bool
	Incompleto           bool
	Data_ultima_compra   string
	Compra_ticket_medio  float32
	Ticket_ultima_compra float32
	Loja_mais_frequente  string
	Loja_ultima_compra   string
	Data_criacao         time.Time
	Data_modificacao     time.Time
}

//CompraClientes type CompraCliente list
type CompraClientes []CompraCliente

//Constructor NewCompraCliente de CompraCliente's
func NewCompraCliente() *CompraCliente {
	return &CompraCliente{
		Data_criacao:     time.Now(),
		Data_modificacao: time.Now(),
	}
}
