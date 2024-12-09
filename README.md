# Go Modular Backend Template

Este é um projeto template para um backend modular em Go, demonstrando uma arquitetura limpa e escalável com exemplos de conexões a diferentes bancos de dados e sistemas de mensageria.

## Estrutura do Projeto

O projeto segue uma arquitetura limpa e modular:

- `cmd/`: Contém os pontos de entrada da aplicação.

  - `api/`: Contém o arquivo `main.go` que inicia o servidor HTTP.

- `internal/`: Contém o código interno da aplicação, não destinado a ser importado por outros projetos.

  - `app/`: Contém a lógica de aplicação.
    - `services/`: Implementa a lógica de negócios.
  - `domain/`: Define as entidades de domínio e interfaces de repositório.
    - `entities/`: Define as estruturas de dados principais.
    - `repositories/`: Define as interfaces para acesso a dados.
  - `infra/`: Implementa as interfaces definidas no domínio.
    - `database/`: Contém as implementações de acesso a banco de dados.
      - `postgres/`: Implementação do repositório PostgreSQL.
      - `mongodb/`: Implementação do repositório MongoDB.
    - `messaging/`: Contém as implementações de sistemas de mensageria.
      - `rabbitmq/`: Implementação da conexão e operações com RabbitMQ.
  - `interfaces/`: Lida com as entradas externas para a aplicação.
    - `http/`: Contém os handlers HTTP.
      - `handlers/`: Implementa os handlers para as rotas HTTP.

- `pkg/`: Contém código que pode ser utilizado por aplicações externas.
  - `logger/`: Implementa funcionalidades de logging.

Esta estrutura permite uma clara separação de responsabilidades e facilita a manutenção e expansão do projeto.

## Instalação e Execução

Siga estas etapas para instalar e executar o projeto:

1. **Pré-requisitos**

   - Go 1.16 ou superior
   - PostgreSQL
   - MongoDB
   - RabbitMQ

2. **Clone o repositório**
   \`\`\`
   git clone https://github.com/romulosm/go-modular-backend-template.git
   cd go-modular-backend-template
   \`\`\`

3. **Instale as dependências**
   \`\`\`
   go mod tidy
   \`\`\`

4. **Configure as variáveis de ambiente**
   Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:
   \`\`\`

# Configurações do PostgreSQL

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=mypassword
POSTGRES_DB=mydb

# Configurações do MongoDB

MONGODB_URI=mongodb://localhost:27017/mydb

# Configurações do RabbitMQ

RABBITMQ_URL=amqp://guest:guest@localhost:5672/

# Configuração de Logging

LOG_LEVEL=info

# Configuração do Servidor

SERVER_PORT=8080
\`\`\`
Substitua os valores acima pelos corretos para o seu ambiente.

5. **Inicie os serviços de banco de dados e mensageria**
   Certifique-se de que PostgreSQL, MongoDB e RabbitMQ estejam em execução em sua máquina.

6. **Execute o projeto**
   \`\`\`
   go run cmd/api/main.go
   \`\`\`

O servidor deve iniciar e você verá uma mensagem como:
\`\`\`
Servidor iniciado na porta 8080
\`\`\`
Note que a porta pode ser diferente se você configurou SERVER_PORT no arquivo .env.

7. **Teste a API**
   Você pode testar a API usando curl ou qualquer cliente HTTP de sua preferência:
   \`\`\`
   curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}'
   \`\`\`

Agora o projeto está em execução e pronto para uso!
