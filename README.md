# API de Formulário de Contato

## Descrição

Esta API permite o envio de formulários de contato com validação de ReCaptcha e envio de emails tanto para o usuário quanto para a empresa. A API também implementa tratamento de erros conforme o padrão RFC7807.

## Funcionalidades

- Envio de formulários de contato.
- Validação de ReCaptcha ( compátivel com ReCaptcha v2)
- Envio de email para o usuário e a empresa.
- Retorno de erros detalhados no formato RFC7807.

## Endpoints

### 1. Enviar Formulário de Contato

` POST /contact`

Envia os dados do formulário de contato e realiza a validação ReCaptcha

### Request Body (JSON):

```typescript
{
  "g-recaptcha-response": "my-correct-captcha",
  "comment": "Minha mensagem",
  "name": "John Doe",
  "mail": "john.doe@example.com"
}
```

- `g_recaptcha_response`: A resposta do captcha. Para testes, use o valor valid-captcha-token.
- `comment`: O comentário enviado pelo usuário.
- `name`: O nome do usuario.
- `mail`: O email do usuario.

#### Response Status:

- `201 Created`: Quando o formulario foi enviado com sucesso.

##### Exemplo de Sucesso

```html
HTTP/1.1 201 Created
```

### 2. Tratamento de Erros

A API segue o padrão RFC7807 para retornar erros, com detalhes sobre o problema em questao.

##### 2.1 Erro de Captcha invalido

##### Request Body:

```JSON
{
  "g-recaptcha-response": "invalid-captcha",
  "comment": "Minha mensagem",
  "name": "John Doe",
  "mail": "john.doe@example.com"
}
```

#### resposne status:

- `401 Unauthorized`: Quando o captcha e invalido.

##### Exemplo de Resposta:

```json
{
  "type": "about:blank",
  "title": "UnauthorizedError",
  "detail": "The captcha is incorrect!",
  "instance": "/api-endpoint"
}
```

##### 2.2. Erro de Email invalido

##### Request body:

```json
{
  "g-recaptcha-response": "valid-captcha-token",
  "comment": "Minha mensagem",
  "name": "John Doe",
  "mail": "john.doe@"
}
```

#### Response Status:

- `400 Bad Request`: Quando o email enviado e valido.

#### Exemplo de Resposta:

```JSON
{
  "type": "about:blank",
  "title": "BadRequestError",
  "detail": "The email is invalid",
  "instance": "/api-endpoint"
}
```

##### 2.3 Erro de Nome Vazio

##### Rquest Body:

```JSON
{
  "g-recaptcha-response": "valid-captcha-token",
  "comment": "Minha mensagem",
  "name": "",
  "mail": "john.doe@example.com"
}
```

#### Response Status:

- `400 Bad Request`: Quando o nome do usuario esta vazio.

##### Exemplo de Resposta:

```JSON
{
  "type": "about:blank",
  "title": "BadRequestError",
  "detail": "The name is empty",
  "instance": "/api-endpoint"
}
```

| Variável          | Descrição                                                                                       |
| ----------------- | ----------------------------------------------------------------------------------------------- |
| `PORT`            | Porta onde o servidor vai rodar (e.g., 8080).                                                   |
| `RECAPTCHA_KEY`   | Chave secreta da API ReCaptcha.                                                                 |
| `RECAPTCHA_URL`   | URL para verificação do ReCaptcha.                                                              |
| `MAIL_HOST`       | Servidor SMTP para envio de emails.                                                             |
| `MAIL_PORT`       | Porta do servidor SMTP.                                                                         |
| `MAIL_SECURE`     | Se o SMTP usa SSL/TLS (true ou false).                                                          |
| `MAIL_AUTH_USER`  | Email do remetente (usado para autenticação SMTP).                                              |
| `MAIL_AUTH_PASS`  | Senha de autenticação SMTP.                                                                     |
| `TEXT_MAIL_TITLE` | Título do email.                                                                                |
| `TEXT_MAIL_BODY`  | Corpo do email em texto. Exemplo: `Contact from {name}, using mail: {email}, about: {comment}`. |
| `TEXT_MAIL_HTML`  | Corpo do email em HTML. Exemplo: `Contact from {name}, using mail: {email}, about: {comment}`.  |

#### Testes Automatizados

Para rodar os testes automatizados, use o comando:

```bash
go test ./test/NOME-DO-ARQUIVO
```

Os testes cobrem cenários como:

- Testa se o envio dos emails vai ser bem sucedido
- testa se existe as variaveis de ambientes estao presentes

#### Execução com Docker

Você pode executar a aplicação em um container Docker:

##### Contruir a imagem:

```bash
docker build -t backend_challenge_go .
```

### Executar o container:

```bash
docker run -p 8080:8080 backend_challenge_go
```

### Documentação com Swagger

A API está documentada utilizando Swagger. Para acessar a documentação interativa, acesse o endpoint `/api-docs` após iniciar a aplicação:

```bash
http://localhost:8080/api-docs
```
