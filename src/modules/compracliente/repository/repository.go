package repository

import (
	"../model"
)

//CompraClienteRepository
type CompraClienteRepository interface {
	InsertCompra(string) error
	Save(*model.CompraCliente) error
	Update(string, *model.CompraCliente) error
	Delete(string) error
	FindByCPF(string) (*model.CompraCliente, error)
	FindAllTmp() (model.CompraClientes, error)
	FindAll() (model.CompraClientes, error)
	DeleteAll() error
	DeleteAllTmp() error
}
