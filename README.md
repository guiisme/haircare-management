# HairCare Manager

Projeto HairCare Manager com Go + Angular + CSS + JWT

Um sistema (backend em Go e frontend em Angular) para gerenciar clientes, serviços e agendamentos de um salão de beleza. Este repositório contém a API em Go e a interface em Angular (estrutura de pastas pode variar — ajuste os comandos abaixo conforme a organização do seu repositório).

---

## Principais características

- Backend em Go (API REST)
- Frontend em Angular (SPA)
- Autenticação com JWT
- Containerização com Docker (Dockerfile incluído)
- Simples, pronto para extensões (serviços, clientes, agendamentos, permissões)

---

## Tecnologias

- Go (principal)
- Angular (frontend)
- CSS
- JWT para autenticação
- Docker (opcional)

---

## Requisitos (pré-instalação)

- Go (1.18+ recomendado)
- Node.js + npm (ou pnpm/yarn) — para o frontend Angular
- Angular CLI (se for desenvolver localmente o frontend)
- Docker & Docker Compose (opcional, para containerização)
- Sistema de banco de dados (Postgres, MySQL, SQLite — a configuração depende do que o projeto usa)

---

## Estrutura sugerida do repositório

Observação: o repositório pode ter uma organização diferente. Abaixo está uma estrutura comum que muitos projetos com Go + Angular usam:

- /backend (ou /server) — código Go da API
- /frontend (ou /web) — código Angular
- Dockerfile — imagem (geralmente do backend) e/ou root Dockerfile
- docker-compose.yml (opcional)

Ajuste os caminhos/nomes de pastas nos comandos abaixo conforme a estrutura real do projeto.

---

## Variáveis de ambiente esperadas

Configure as variáveis abaixo conforme o ambiente (desenvolvimento/produção). Nomes e necessárias exatas podem variar no código — confira os arquivos de configuração/README do backend.

Exemplo (.env):
```
PORT=8080
JWT_SECRET=uma_chave_secreta_forte
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=senha
DB_NAME=haircare_db
```

Principais variáveis:
- PORT — porta HTTP da API
- JWT_SECRET — segredo para assinar/validar tokens JWT
- DB_* — conexão com o banco de dados

---

## Executando localmente

Abaixo instruções genéricas. Se seu backend e frontend estiverem em pastas diferentes, execute dentro de cada pasta.

1. Preparar backend (Go)

```bash
# a partir da raiz, se o backend estiver em ./backend
cd backend

# instalar dependências (go modules)
go mod download

# configurar variáveis de ambiente (ex.: export JWT_SECRET=...)
# executar a API
go run ./cmd/main.go   # ou: go run main.go, depende da estrutura do projeto

# ou compilar
go build -o haircare-api ./...
./haircare-api
```

2. Preparar frontend (Angular)

```bash
# a partir da raiz, se o frontend estiver em ./frontend
cd frontend

# instalar dependências
npm install
# ou: yarn install

# executar em modo de desenvolvimento
ng serve --open
# ou (se não tiver Angular CLI instalado globalmente)
npx ng serve --open
```

Ao rodar ambos, acesse o frontend (por padrão http://localhost:4200) e verifique se o frontend aponta para a URL da API (ex.: http://localhost:8080). Ajuste as configurações de proxy/ambiente do Angular se necessário.

---

## Build para produção

Backend (Go):
```bash
# build estático/prod
cd backend
go build -o haircare-api ./...
# faça deploy do binário ou construa imagem Docker
```

Frontend (Angular):
```bash
cd frontend
ng build --configuration=production
# os arquivos gerados ficarão em dist/
```

---

## Executando com Docker (exemplo)

Se houver Dockerfile e/ou docker-compose.yml, você pode rodar via containers.

Exemplo simples (Dockerfile presente na raiz ou no backend):
```bash
# build da imagem
docker build -t haircare-manager:latest .

# rodar
docker run -e JWT_SECRET="sua_chave" -p 8080:8080 haircare-manager:latest
```

Exemplo docker-compose (ilustrativo — adapte conforme serviços e nomes reais):
```yaml
version: "3.8"
services:
  api:
    build: ./backend
    environment:
      - JWT_SECRET=uma_chave_secreta
      - DB_HOST=db
      - DB_USER=postgres
      - DB_PASSWORD=senha
      - DB_NAME=haircare_db
    ports:
      - "8080:8080"
    depends_on:
      - db

  frontend:
    build: ./frontend
    ports:
      - "4200:80"

  db:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: senha
      POSTGRES_DB: haircare_db
    volumes:
      - db-data:/var/lib/postgresql/data

volumes:
  db-data:
```

---

## Autenticação (JWT)

A API usa JWT para autenticação. Fluxo comum:
- Usuário faz login com credenciais -> API retorna token JWT
- Frontend salva token (ex.: localStorage) e o envia no header Authorization: Bearer <token> nas requisições protegidas
- O backend valida o token usando JWT_SECRET

Verifique o middleware/handler de autenticação no código Go para detalhes (nome das rotas, claims, expiração do token).

---

## Testes

Instruções genéricas (ajuste conforme testes existentes):

Backend (Go):
```bash
cd backend
go test ./...
```

Frontend (Angular):
```bash
cd frontend
ng test
# ou
npm test
```

---

## Como contribuir

1. Abra uma issue descrevendo o que deseja melhorar/bug encontrado.
2. Crie uma branch feature ou fix (ex.: feat/login, fix/jwt-expiry).
3. Faça commits pequenos e atômicos.
4. Abra um Pull Request apontando a branch base — descreva o que foi alterado e por quê.

Sinta-se à vontade para abrir issues pedindo features, documentação adicional ou relatando bugs.

---

## Boas práticas de segurança

- Nunca comite JWT_SECRET, senhas, ou credenciais no repositório.
- Use variáveis de ambiente ou serviços de segredo (Vault, GitHub Secrets).
- Revise bibliotecas e dependências regularmente.

---

## Licença

Se desejar, adicione uma licença (ex.: MIT). Atualmente nenhum arquivo de licença foi especificado neste README. Para aplicar MIT, crie um arquivo LICENSE com o texto apropriado.

---

## Contato

Autor: guiisme  
Descrição do repo: "Projeto HairCare Manager com Go + Angular + CSS + JWT"

Se quiser que eu gere arquivos auxiliares (por exemplo, docker-compose.yml, um exemplo de .env, ou um CONTRIBUTING.md), diga qual e eu gero um modelo adaptado ao seu projeto.
