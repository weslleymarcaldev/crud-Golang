**CRUD em Golang**
**Repositório**: [CRUD Golang](https://github.com/weslleymarcaldev/crud-Golang)

```markdown
# CRUD em Golang 🛠️

Este projeto implementa um sistema CRUD (Create, Read, Update, Delete) utilizando a linguagem Go (Golang) e um banco de dados MySQL. O objetivo é criar uma API RESTful funcional com operações básicas de persistência de dados.

## 🚀 Funcionalidades
- Cadastro de registros.
- Listagem de registros individuais ou em massa.
- Atualização de dados existentes.
- Exclusão de registros.

## 🛠️ Tecnologias Utilizadas
- **Golang**: Linguagem de programação.
- **MySQL**: Banco de dados relacional.
- **Docker**: Gerenciamento de contêineres (opcional).
- **Postman**: Para testes da API (opcional).

## 📦 Como Executar
1. Clone o repositório:
   ```bash
   git clone https://github.com/weslleymarcaldev/crud-Golang.git
2. Configure o banco de dados MySQL:
  Crie um banco de dados chamado crud_golang.
  Atualize o arquivo config.go com suas credenciais de banco de dados.
3. Execute a aplicação:
  go run main.go
4. Teste a API:
  POST: /create - Criar registro.
  GET: /read - Ler registros.
  PUT: /update - Atualizar registro.
  DELETE: /delete - Deletar registro.

🧑‍💻 Contribuições
Contribuições são incentivadas! Abra uma issue para sugestões ou envie um pull request.

📄 Licença
Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.
