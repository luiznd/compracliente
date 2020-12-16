CREATE SEQUENCE public.sq_pk_compracliente START 1;

CREATE TABLE compracliente (
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
)