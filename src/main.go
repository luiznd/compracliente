package main

import (
	"fmt"
	"strings"
	"time"

	"./config"

	"./fileinput"
	"./modules/compracliente/model"
	"./modules/compracliente/repository"
	"./utils"
)

func main() {
	var conteudo []string
	var inic time.Time
	var final time.Time

	conteudo = nil
	fmt.Println("====================================================")
	fmt.Println("Go! integrando registros do arquivo para table temp")
	inic = time.Now()
	fmt.Println("====================================================")

	// Cria a conexão com  o banco
	db, err := config.GetPostgresDB()

	compraclienteRepositoryPostgres := repository.NewCompraClienteRepositoryPostgres(db)

	if err != nil {
		fmt.Println(err)
	}

	//Limpa tabela gateway tmp
	err = deleteCompraClienteAllTmp(compraclienteRepositoryPostgres)
	if err != nil {
		fmt.Println(err)
	}

	//Limpa tabela final compracliente
	err = deleteCompraClienteAll(compraclienteRepositoryPostgres)
	if err != nil {
		fmt.Println(err)
	}

	// Carrega o arquivo com os dados
	conteudo, err = fileinput.GetText("../files/base_teste.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Insere todas as linhas do arquivo em uma tabela temporária, remove os espaços em branco até restar só um espaço e padronizar os campos
	for indice := range conteudo {
		if indice != 0 {
			registro := strings.Fields(conteudo[indice])
			err := insertCompraCliente(registro, compraclienteRepositoryPostgres)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	final = time.Now()

	fmt.Println("Fim da integração do arquivo")
	fmt.Println("==============================")

	//Cria uma lista de compra clientes
	compraclientes, err := getCompraClientesTmp(compraclienteRepositoryPostgres)

	if err != nil {
		fmt.Println(err)
	}

	//Cria um tipo de compra clientes
	cliente := model.NewCompraCliente()

	// Insere semente as linhas validas da tabela tmp na tabela final, valida os CPFs e CNPJs, descarta os inválidos
	for v := range compraclientes {

		cliente.Cpf = compraclientes[v].Cpf
		cliente.Private = compraclientes[v].Private
		cliente.Incompleto = compraclientes[v].Incompleto
		cliente.Data_ultima_compra = compraclientes[v].Data_ultima_compra
		cliente.Compra_ticket_medio = compraclientes[v].Compra_ticket_medio
		cliente.Ticket_ultima_compra = compraclientes[v].Ticket_ultima_compra
		cliente.Loja_mais_frequente = compraclientes[v].Loja_mais_frequente
		cliente.Loja_ultima_compra = compraclientes[v].Loja_ultima_compra
		cliente.Data_criacao = compraclientes[v].Data_criacao
		cliente.Data_modificacao = compraclientes[v].Data_modificacao

		if (utils.IsCPF(cliente.Cpf)) == false {

			fmt.Println(cliente.Cpf + " Não é um CPF válido!")

		} else if (utils.IsCNPJ(cliente.Loja_mais_frequente)) == false {

			if cliente.Loja_mais_frequente != "NULL" {
				fmt.Println(cliente.Loja_mais_frequente + " Não é um CNPJ válido para Loja_mais_frequente!")
			}

		} else if (utils.IsCNPJ(cliente.Loja_ultima_compra)) == false {

			if cliente.Loja_mais_frequente != "NULL" {
				fmt.Println(cliente.Loja_ultima_compra + " Não é um CNPJ válido para Loja_ultima_compra!")
			}

		} else {

			err = saveCompraCliente(cliente, compraclienteRepositoryPostgres)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	final = time.Now()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("==================" + inic.Format("2006-01-02 15:04:05") + "=======================")
	fmt.Println("Fim da inserção de dados na tabela final CompraCliente")
	fmt.Println("==================" + final.Format("2006-01-02 15:04:05") + "======================")
	fmt.Println("")
}

// Insere os registros na tabela temporária
func insertCompraCliente(p []string, repo repository.CompraClienteRepository) error {
	err := repo.InsertCompra(p)

	if err != nil {
		return err
	}

	return nil
}

// Inser os registros na tabela final CompraCliente
func saveCompraCliente(p *model.CompraCliente, repo repository.CompraClienteRepository) error {
	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

// Update na tabela final CompraClienteRepository
func updateCompraCliente(p *model.CompraCliente, repo repository.CompraClienteRepository) error {
	err := repo.Update(p.Cpf, p)

	if err != nil {
		return err
	}

	return nil
}

// Delete na tabela final CompraClienteRepository
func deleteCompraCliente(id string, repo repository.CompraClienteRepository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

// Seleciona um registros filtrando por cpf na tabela final CompraClienteRepository
func getCompraCliente(cpf string, repo repository.CompraClienteRepository) (*model.CompraCliente, error) {
	compracliente, err := repo.FindByCPF(cpf)

	if err != nil {
		return nil, err
	}

	return compracliente, nil
}

// Traz todos os registros da tabela temporária
func getCompraClientesTmp(repo repository.CompraClienteRepository) (model.CompraClientes, error) {
	compraclientes, err := repo.FindAllTmp()

	if err != nil {
		return nil, err
	}

	return compraclientes, nil
}

// Traz todos os registros da tabela final CompraCliente
func getCompraClientes(repo repository.CompraClienteRepository) (model.CompraClientes, error) {
	compraclientes, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return compraclientes, nil
}

// Delete na tabela final
func deleteCompraClienteAll(repo repository.CompraClienteRepository) error {
	err := repo.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}

// Delete na tabela tmp
func deleteCompraClienteAllTmp(repo repository.CompraClienteRepository) error {
	err := repo.DeleteAllTmp()

	if err != nil {
		return err
	}

	return nil
}
