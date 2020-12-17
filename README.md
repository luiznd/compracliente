Desafio - Compra cliente
=======================

### Requisitos

* Docker
https://www.docker.com/products/docker-desktop


Diretório raiz dos codigos ./src
<br>
Diretório de arquivos de carga, confifuração e sql ./files

### Cadastro de compras por Cpf

* Projeto de carga de dados ETL através de arquivo texto utilizando linguagem GO e banco de dados PostgresSql, há um CRUD também para facilitar a manipulação dos dados posteriormente.

* <p>Foco do projeto - > Carga de dados direto do arquivo base_teste.txt em tabela gateway temp_compracliente, Inclusão, validação e tratamento dos registros na tabela final compracliente.</p> 

### Instalação / Execução
```
* Clonar projeto em um diretório local com o comando "git clone https://github.com/luiznd/compracliente"
```
```
* Acessar a pasta ./src do projeto no seu terminal
```
```
* Build dos serviços do docker com o comando:
$ docker-compose up --build -d
```
```
* Para verificar o log e status do carregamento do arquivo digitar o comando:
$ docker logs --follow full_app
```
```
* Para acessar o banco de dados e verificar as tabelas acesse o pgAdmin do Docker

    1-  http://localhost:5050
        Usuario : luiznd@hotmail.com
        senha: root
        
    2 - Encontrar o IP do postgres no Docker para criar server no pgAdimn conforme comando e link da imagem abaixo:
    
        $ docker inspect full_db_postgres
    
        https://1drv.ms/u/s!Apyd3zQWtpsdgh3Ymc3QXj8Zj1T8?e=CmtTMq


    3 - Criar um novo server conforme a imagem do link abaixo, com o ip listado no passo anterior, senha "root"

        https://1drv.ms/u/s!Apyd3zQWtpsdghwBAhdYb2jiWXGF?e=KlfbmA
```
```  
* Parar o Docker se quiser
$ docker-compose stop
```  
```  
* Excluir todas as imagens do Docker se necessário
$ docker system prune -a --volumes
``` 

### Execução local fora do Docker
```
* Alterar as configurações de banco postgres se for trabalhar local:
  Editar aquivo na pasta 'src/files/config.ini'
  Linha 1 : usuário
  Linha 2 : senha
  Linha 3 : servidor
  Linha 4 : porta
  Linha 5 : database
```  
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


### Conector Go com o banco de dados postgresSql
\src\config\postgres.go

### Controle
\src\utils\main.go

