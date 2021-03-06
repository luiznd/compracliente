package repository

import (
	"database/sql"

	"../model"
)

type compraclienteRepositoryPostgres struct {
	db *sql.DB
}

func NewCompraClienteRepositoryPostgres(db *sql.DB) *compraclienteRepositoryPostgres {
	return &compraclienteRepositoryPostgres{db}
}

// Os dados são inseridos na tabela final compracliente transformando os dados em mauisculo e retirados possíveis acentos
func (r *compraclienteRepositoryPostgres) Save(compracliente *model.CompraCliente) error {

	query := `INSERT INTO "compracliente"("cpf", "private", "incompleto", "data_ultima_compra", "compra_ticket_medio", "ticket_ultima_compra", "loja_mais_frequente", "loja_ultima_compra", "data_criacao","data_modificacao")
        VALUES(
			UPPER(translate($1, 'áéíóúàèìòùãõâêîôôäëïöüçÁÉÍÓÚÀÈÌÒÙÃÕÂÊÎÔÛÄËÏÖÜÇ', 'aeiouaeiouaoaeiooaeioucAEIOUAEIOUAOAEIOOAEIOUC')), 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			UPPER(translate($7, 'áéíóúàèìòùãõâêîôôäëïöüçÁÉÍÓÚÀÈÌÒÙÃÕÂÊÎÔÛÄËÏÖÜÇ', 'aeiouaeiouaoaeiooaeioucAEIOUAEIOUAOAEIOOAEIOUC')), 
			UPPER(translate($8, 'áéíóúàèìòùãõâêîôôäëïöüçÁÉÍÓÚÀÈÌÒÙÃÕÂÊÎÔÛÄËÏÖÜÇ', 'aeiouaeiouaoaeiooaeioucAEIOUAEIOUAOAEIOOAEIOUC')), 
			$9, 
			$10)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(compracliente.Cpf, compracliente.Private, compracliente.Incompleto, compracliente.Data_ultima_compra, compracliente.Compra_ticket_medio, compracliente.Ticket_ultima_compra, compracliente.Loja_mais_frequente, compracliente.Loja_ultima_compra, compracliente.Data_criacao, compracliente.Data_modificacao)

	if err != nil {
		return err
	}

	return nil
}

// Insere os registros que vem do arquivo em uma tabela tmp, faz o split das colunas e aplica uma regra para os nulos.
func (r *compraclienteRepositoryPostgres) InsertCompra(linha []string) error {

	query := `INSERT INTO "tmp_compracliente"("cpf", "private", "incompleto", "data_ultima_compra", "compra_ticket_medio", "ticket_ultima_compra", "loja_mais_frequente", "loja_ultima_compra", "data_criacao","data_modificacao")
	VALUES($1, $2, $3, 
		CASE $4 WHEN 'NULL' THEN NULL ELSE TO_DATE($4,'YYYY-MM-DD') END, 
		CASE $5 WHEN 'NULL' THEN NULL ELSE cast(REPLACE($5,',','.') as float) END, 
		CASE $6 WHEN 'NULL' THEN NULL ELSE cast(REPLACE($6,',','.') as float) END, 
		CASE $7 WHEN 'NULL' THEN NULL ELSE $7 END, 
		CASE $8 WHEN 'NULL' THEN NULL ELSE $8 END, 
		now(), 
		now())`
	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(linha[0], linha[1], linha[2], linha[3], linha[4], linha[5], linha[6], linha[7])

	if err != nil {
		return err
	}

	return nil
}

// Atualiza os dados na tabela final compracliente por cpf
func (r *compraclienteRepositoryPostgres) Update(cpf string, compracliente *model.CompraCliente) error {

	query := `UPDATE "compracliente" SET "private"=$1, "incompleto"=$2, "data_ultima_compra"=$3, "compra_ticket_medio"=$4, "ticket_ultima_compra"=$5, "loja_mais_frequente"=$6, "loja_ultima_compra"=$7, "data_criacao"=$8, "data_modificacao"=$9 WHERE "cpf"=$10`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(compracliente.Private, compracliente.Incompleto, compracliente.Data_ultima_compra, compracliente.Compra_ticket_medio, compracliente.Ticket_ultima_compra, compracliente.Loja_mais_frequente, compracliente.Loja_ultima_compra, compracliente.Data_criacao, compracliente.Data_modificacao, cpf)

	if err != nil {
		return err
	}

	return nil
}

// Deleta os dados na tabela final compracliente por cpf
func (r *compraclienteRepositoryPostgres) Delete(cpf string) error {

	query := `DELETE FROM "compracliente" WHERE "CPF" = $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(cpf)

	if err != nil {
		return err
	}

	return nil
}

func (r *compraclienteRepositoryPostgres) FindByCPF(cpf string) (*model.CompraCliente, error) {
	query := `SELECT * FROM "compracliente" WHERE "CPF" = $1`

	var compracliente model.CompraCliente

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(cpf).Scan(&compracliente.Cpf, &compracliente.Private, &compracliente.Incompleto, &compracliente.Data_ultima_compra, &compracliente.Compra_ticket_medio, &compracliente.Ticket_ultima_compra, &compracliente.Loja_mais_frequente, &compracliente.Loja_ultima_compra, &compracliente.Data_criacao, &compracliente.Data_criacao, &compracliente.Data_modificacao)

	if err != nil {
		return nil, err
	}

	return &compracliente, nil

}

// Lista todos os registros da tabela tmp para ser feita a higienização dos dados após persistência , e a validação de CPFs/CNPJs contidos (válidos e não válidos numericamente)
func (r *compraclienteRepositoryPostgres) FindAllTmp() (model.CompraClientes, error) {

	query := `SELECT 
				cpf, 
				private, 
				incompleto, 
				CASE WHEN data_ultima_compra IS NULL THEN 'NULL' ELSE to_char(data_ultima_compra, 'YYYY-MM-DD') END as data_ultima_compra,
				CASE WHEN compra_ticket_medio IS NULL THEN 0 ELSE compra_ticket_medio  END as compra_ticket_medio,
				CASE WHEN ticket_ultima_compra IS NULL THEN 0 ELSE ticket_ultima_compra END as ticket_ultima_compra,
				CASE WHEN loja_mais_frequente IS NULL THEN 'NULL' ELSE loja_mais_frequente  END as loja_mais_frequente, 
				CASE WHEN loja_ultima_compra IS NULL THEN 'NULL' ELSE loja_ultima_compra  END as loja_ultima_compra,
				data_criacao,
				data_modificacao 
				FROM tmp_compracliente`

	var compraclientes model.CompraClientes

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var compracliente model.CompraCliente

		err = rows.Scan(&compracliente.Cpf, &compracliente.Private, &compracliente.Incompleto, &compracliente.Data_ultima_compra, &compracliente.Compra_ticket_medio, &compracliente.Ticket_ultima_compra, &compracliente.Loja_mais_frequente, &compracliente.Loja_ultima_compra, &compracliente.Data_criacao, &compracliente.Data_modificacao)

		if err != nil {
			return nil, err
		}

		compraclientes = append(compraclientes, compracliente)
	}

	return compraclientes, nil
}

// Lista todos os registros da tabela final
func (r *compraclienteRepositoryPostgres) FindAll() (model.CompraClientes, error) {

	query := `SELECT 
				cpf, 
				private, 
				incompleto, 
				CASE WHEN data_ultima_compra IS NULL THEN 'NULL' ELSE to_char(data_ultima_compra, 'YYYY-MM-DD') END as data_ultima_compra,
				CASE WHEN compra_ticket_medio IS NULL THEN 0 ELSE compra_ticket_medio  END as compra_ticket_medio,
				CASE WHEN ticket_ultima_compra IS NULL THEN 0 ELSE ticket_ultima_compra END as ticket_ultima_compra,
				CASE WHEN loja_mais_frequente IS NULL THEN 'NULL' ELSE loja_mais_frequente  END as loja_mais_frequente, 
				CASE WHEN loja_ultima_compra IS NULL THEN 'NULL' ELSE loja_ultima_compra  END as loja_ultima_compra,
				data_criacao,
				data_modificacao 
				FROM compracliente`

	var compraclientes model.CompraClientes

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var compracliente model.CompraCliente

		err = rows.Scan(&compracliente.Cpf, &compracliente.Private, &compracliente.Incompleto, &compracliente.Data_ultima_compra, &compracliente.Compra_ticket_medio, &compracliente.Ticket_ultima_compra, &compracliente.Loja_mais_frequente, &compracliente.Loja_ultima_compra, &compracliente.Data_criacao, &compracliente.Data_modificacao)

		if err != nil {
			return nil, err
		}

		compraclientes = append(compraclientes, compracliente)
	}

	return compraclientes, nil
}

// Deleta todos os registros da tabela final
func (r *compraclienteRepositoryPostgres) DeleteAll() error {

	query := `DELETE FROM "compracliente"`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	return nil
}

// Deleta todos os registros da tabela tmp
func (r *compraclienteRepositoryPostgres) DeleteAllTmp() error {

	query := `DELETE FROM "tmp_compracliente"`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	return nil
}

// Cria tabela tmp se não existir
func (r *compraclienteRepositoryPostgres) CreateTableTmp() error {

	seq := `CREATE SEQUENCE IF NOT EXISTS public.sq_pk_tmp_compracliente START 1`

	query := `CREATE TABLE IF NOT EXISTS tmp_compracliente (
		id_compracliente BIGINT NOT NULL DEFAULT nextval('public.sq_pk_tmp_compracliente'),
		CPF VARCHAR(20) NOT NULL,
		private VARCHAR(20),
		incompleto VARCHAR(20),
		data_ultima_compra date,
		compra_ticket_medio float,
		ticket_ultima_compra float,
		loja_mais_frequente VARCHAR(20),
		loja_ultima_compra VARCHAR(20),
		data_criacao timestamp,
		data_modificacao timestamp,
		CONSTRAINT pk_tmp_compracliente PRIMARY KEY (id_compracliente)
	)`

	statement, err := r.db.Prepare(seq)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	statement2, err2 := r.db.Prepare(query)

	if err2 != nil {
		return err2
	}

	defer statement2.Close()

	_, err2 = statement2.Exec()

	if err2 != nil {
		return err2
	}

	return nil
}

// Cria tabela final se não existir
func (r *compraclienteRepositoryPostgres) CreateTable() error {

	seq := `CREATE SEQUENCE IF NOT EXISTS public.sq_pk_compracliente START 1`

	query := `CREATE TABLE IF NOT EXISTS compracliente (
		id_compracliente BIGINT NOT NULL DEFAULT nextval('public.sq_pk_compracliente'),
		CPF VARCHAR(20) NOT NULL,
		private VARCHAR(20),
		incompleto VARCHAR(20),
		data_ultima_compra date,
		compra_ticket_medio float,
		ticket_ultima_compra float,
		loja_mais_frequente VARCHAR(20),
		loja_ultima_compra VARCHAR(20),
		data_criacao timestamp,
		data_modificacao timestamp,
		CONSTRAINT pk_compracliente PRIMARY KEY (id_compracliente)
	)`

	statement, err := r.db.Prepare(seq)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	statement2, err2 := r.db.Prepare(query)

	if err2 != nil {
		return err2
	}

	defer statement2.Close()

	_, err2 = statement2.Exec()

	if err2 != nil {
		return err2
	}

	return nil
}
