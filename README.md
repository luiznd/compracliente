Desafio - Carga de dados através de arquivo - Compra cliente
=======================

Diretório raiz dos codigos ./src
<br>
Diretório de arquivos de carga, confifuração e sql ./files

### Cadastro de compras por Cpf

* Projeto de carga de dados através de arquivo texto utilizando linguagem GO e banco de dados PostgresSql.

* <p>Foco do projeto - > Carga de dados direto do arquivo base_teste.txt em tabela gateway temp_compracliente, Inclusão, validação e tratamento dos registros na tabela final compracliente.</p> 

### Instalação
```
* Clonar projeto em um diretório local com o comando "git clone https://github.com/luiznd/compracliente"

* Instalar banco de dados postgresSql : https://www.postgresql.org/download/

* Instalar o o Golang : https://golang.org/doc/install

* Executar os scripts SQLs `tmp_compracliente.sql` e  `compracliente.sql` para criar as tabelas
  que estão na pasta ./files no pgAdmin

* Instalar o Visual Studio Code para visualizar o código e executar o projeto.
```

### Execução
* No terminal do Visual Studio Code acessar a pasta `cd src`  e executar o comando:  `go run main.go`


### Entidade
\src\modules\compracliente\model\compracliente.go


### Repositório model
\src\modules\compracliente\repository\repository.go

\src\modules\compracliente\repository\compracliente_repository_postgres.go


### Manipulação de arquivos
\src\fileinput\inputfile.go


### Utilitários, validação de cnpj/cpf
\src\utils\cpfcnpj.go


### Controle
\src\utils\main.go

