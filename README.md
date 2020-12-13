Desafio - Carga de dados através de arquivo - Compra cliente
=======================

diretório raiz dos codigos ./src
diretório de arquivos de carga, confifuração e sql ./files

### Cadastro de compras por Cpf

Projeto de carga de dados através de arquivo texto utilizando linguagem GO e banco de dados PostgresSql.
<br>
<p>Foco do projeto - > Carga de dados direto do arquivo base_teste.txt em tabela gateway temp_compracliente, Inclusão, validação e tratamento dos registros na tabela final compracliente.</p> 

### Instalação
Clonar projeto em um diretório local com o comando "git clone https://github.com/luiznd/compracliente"
<br>
Instalar banco de dados postgresSql : https://www.postgresql.org/download/
<br>
Instalar o o Golang : https://golang.org/doc/install
<br>
Executar os scripts SQLs <u>tmp_compracliente.sql</u> e  <u>compracliente.sql</u> para criar as tabelas que estão na pasta ./files no pgAdmin 
<br>
Instalar o Visual Studio Code para visualizar e executar o projeto.
<br>

### Execução
No terminal do Visual Studio Code executar o comando:  <u>go run main.go </u>

### Entidade
Controller - > module/Application/src/Application/Controller/IndexController.php
<p>View - > module/Application/view/application/index/index.phtml</p>

### Adição de Tarefas 
Controller - > module/Application/src/Application/Controller/IndexController.php
<p>View - > module/Application/view/application/index/adicionar.phtml</p>

### Edição de Tarefas 
Controller - > module/Application/src/Application/Controller/IndexController.php
<p>View - > module/Application/view/application/index/editar.phtml</p>

