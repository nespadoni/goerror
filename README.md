# GoError - Biblioteca de Tratamento de Erros para Go

Uma biblioteca robusta e estruturada para tratamento de erros em aplicações Go, com tipos pré-definidos, códigos HTTP automáticos e informações detalhadas para debugging.

## Características

- ✅ **65+ tipos de erro pré-definidos** organizados por categoria
- ✅ **Códigos HTTP automáticos** baseados no tipo do erro
- ✅ **Informações de debugging** (arquivo, linha, timestamp)
- ✅ **Sugestões automáticas** para resolução de problemas
- ✅ **Serialização JSON** nativa
- ✅ **Trace ID** para rastreamento distribuído
- ✅ **Detalhes customizáveis** para contexto adicional
- ✅ **Funções utilitárias** para verificação de tipos

## Instalação

```bash
go get github.com/nespadoni/goerror
```

## Uso Básico

### Criando um Erro Simples

```go
package main

import (
    "fmt"
    "github.com/nespadoni/goerror"
)

func main() {
    // Criar um erro básico
    err := goerror.Novo(goerror.EmailInvalido)
    
    fmt.Println(err.Error()) // [EMAIL_INVALIDO] Email informado é inválido
    fmt.Println(err.StatusHTTP) // 400
    fmt.Println(err.Sugestao) // Use o formato: usuario@dominio.com
}
```

### Exemplo em Handler HTTP

```go
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario Usuario
    if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
        erro := goerror.Novo(goerror.FormatoJSON)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(erro.StatusHTTP)
        w.Write(erro.JSON())
        return
    }
    
    if usuario.Email == "" {
        erro := goerror.Novo(goerror.CampoObrigatorio).
            ComDetalhe("campo", "email").
            ComDetalhe("valor_recebido", usuario.Email)
        
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(erro.StatusHTTP)
        w.Write(erro.JSON())
        return
    }
    
    // Lógica de criação...
}
```

## Categorias de Erros

### 🔍 Erros de Validação (400)
Para problemas de entrada de dados:

```go
// Campos obrigatórios
goerror.Novo(goerror.CampoObrigatorio)
goerror.Novo(goerror.CampoInvalido)

// Formatos específicos
goerror.Novo(goerror.EmailInvalido)
goerror.Novo(goerror.CPFInvalido) 
goerror.Novo(goerror.TelefoneInvalido)
goerror.Novo(goerror.DataInvalida)

// Validações de tamanho
goerror.Novo(goerror.TamanhoMinimo)
goerror.Novo(goerror.TamanhoMaximo)
goerror.Novo(goerror.IdadeMinima)

// Outros
goerror.Novo(goerror.SenhaFraca)
goerror.Novo(goerror.FormatoJSON)
```

### 🔐 Erros de Autenticação (401)
Para problemas de login e tokens:

```go
goerror.Novo(goerror.CredenciaisInvalidas)
goerror.Novo(goerror.TokenInvalido)
goerror.Novo(goerror.TokenExpirado)
goerror.Novo(goerror.SessaoExpirada)
goerror.Novo(goerror.LoginObrigatorio)
```

### 🚫 Erros de Autorização (403)
Para problemas de permissão:

```go
goerror.Novo(goerror.SemPermissao)
goerror.Novo(goerror.AcessoNegado)
goerror.Novo(goerror.PerfilInsuficiente)
goerror.Novo(goerror.RecursoProtegido)
```

### 🔍 Erros de Recurso Não Encontrado (404)
Para recursos inexistentes:

```go
goerror.Novo(goerror.UsuarioNaoEncontrado)
goerror.Novo(goerror.ProdutoNaoEncontrado)
goerror.Novo(goerror.PedidoNaoEncontrado)
goerror.Novo(goerror.ArquivoNaoEncontrado)
goerror.Novo(goerror.PaginaNaoEncontrada)
goerror.Novo(goerror.EndpointNaoEncontrado)
```

### ⚔️ Erros de Conflito (409)
Para recursos duplicados:

```go
goerror.Novo(goerror.EmailJaExiste)
goerror.Novo(goerror.CPFJaExiste)
goerror.Novo(goerror.UsuarioJaExiste)
goerror.Novo(goerror.RecursoJaExiste)
goerror.Novo(goerror.EstadoInvalido)
goerror.Novo(goerror.VersaoConflito)
```

### 💼 Erros de Regra de Negócio (422)
Para violações de regras específicas:

```go
goerror.Novo(goerror.SaldoInsuficiente)
goerror.Novo(goerror.EstoqueIndisponivel)
goerror.Novo(goerror.LimiteExcedido)
goerror.Novo(goerror.OperacaoNaoPermitida)
goerror.Novo(goerror.RegraViolada)
goerror.Novo(goerror.StatusInvalido)
goerror.Novo(goerror.PrazoVencido)
```

### 🗄️ Erros de Banco de Dados (500)
Para problemas de persistência:

```go
goerror.Novo(goerror.ErroBancoDados)
goerror.Novo(goerror.ConexaoBanco)
goerror.Novo(goerror.TimeoutBanco)
goerror.Novo(goerror.TransacaoFalhou)
goerror.Novo(goerror.ConstraintViolada)
```

### 🖥️ Erros de Sistema (500)
Para problemas internos:

```go
goerror.Novo(goerror.ErroInterno)
goerror.Novo(goerror.ServicoIndisponivel)
goerror.Novo(goerror.MemoriaInsuficiente)
goerror.Novo(goerror.ConfiguracaoInvalida)
goerror.Novo(goerror.ErroInesperado)
```

### 🌐 Erros de Rede (502/503/504)
Para problemas de conectividade:

```go
goerror.Novo(goerror.ServicoExternoIndisponivel)
goerror.Novo(goerror.TimeoutRede)
goerror.Novo(goerror.ConexaoRecusada)
goerror.Novo(goerror.APIExternaIndisponivel)
```

### ⏱️ Erros de Limite (429)
Para rate limiting:

```go
goerror.Novo(goerror.MuitasRequisicoes)
goerror.Novo(goerror.LimiteAPI)
goerror.Novo(goerror.QuotaExcedida)
```

### 🚷 Erros de Método (405/406)
Para problemas de protocolo:

```go
goerror.Novo(goerror.MetodoNaoPermitido)
goerror.Novo(goerror.VersaoNaoSuportada)
```

## Funcionalidades Avançadas

### Adicionando Detalhes

```go
// Detalhe único
err := goerror.Novo(goerror.CampoObrigatorio).
    ComDetalhe("campo", "email").
    ComDetalhe("valor_recebido", "")

// Múltiplos detalhes
detalhes := map[string]interface{}{
    "usuario_id": 123,
    "tentativas": 3,
    "ip": "192.168.1.1",
}
err := goerror.Novo(goerror.CredenciaisInvalidas).ComDetalhes(detalhes)
```

### Trace ID para Sistemas Distribuídos

```go
err := goerror.Novo(goerror.ServicoExternoIndisponivel).
    ComTraceID("req-abc123-def456")
```

### Encadeamento de Métodos

```go
err := goerror.Novo(goerror.SaldoInsuficiente).
    ComDetalhe("saldo_atual", 100.50).
    ComDetalhe("valor_tentativa", 250.00).
    ComDetalhe("usuario_id", 42).
    ComTraceID("txn-789xyz")
```

## Funções Utilitárias

### Extraindo Informações

```go
err := goerror.Novo(goerror.UsuarioNaoEncontrado)

// Obter status HTTP
status := goerror.ObterStatusHTTP(err) // 404

// Obter código do erro
codigo := goerror.ObterCodigo(err) // "USUARIO_NAO_ENCONTRADO"
```

### Verificação de Categorias

```go
err := goerror.Novo(goerror.EmailInvalido)

// Verificar tipo de erro
if goerror.EhValidacao(err) {
    // É um erro de validação (400)
}

if goerror.EhAutenticacao(err) {
    // É um erro de autenticação (401)
}

if goerror.EhAutorizacao(err) {
    // É um erro de autorização (403)
}

if goerror.EhNaoEncontrado(err) {
    // É um erro 404
}

if goerror.EhConflito(err) {
    // É um erro de conflito (409)
}

if goerror.EhNegocio(err) {
    // É um erro de regra de negócio (422)
}

if goerror.EhSistema(err) {
    // É um erro de sistema (500+)
}
```

## Serialização JSON

### Output JSON Automático

Quando você chama `err.JSON()`, a biblioteca retorna:

```json
{
  "codigo": "EMAIL_INVALIDO",
  "mensagem": "Email informado é inválido",
  "sugestao": "Use o formato: usuario@dominio.com",
  "detalhes": {
    "campo": "email",
    "valor_recebido": "email-inválido"
  },
  "timestamp": "2024-08-04T15:30:45Z",
  "arquivo": "/app/handlers/user.go",
  "linha": 45,
  "trace_id": "req-abc123"
}
```

### Middleware de Tratamento de Erros

```go
func ErrorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                var customErr *goerror.ErroCompleto
                
                switch e := err.(type) {
                case *goerror.ErroCompleto:
                    customErr = e
                case error:
                    customErr = goerror.Novo(goerror.ErroInterno).
                        ComDetalhe("erro_original", e.Error())
                default:
                    customErr = goerror.Novo(goerror.ErroInesperado).
                        ComDetalhe("panic_value", fmt.Sprintf("%v", err))
                }
                
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(customErr.StatusHTTP)
                w.Write(customErr.JSON())
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

## Exemplos Práticos

### Validação de Dados

```go
func ValidarUsuario(usuario *Usuario) error {
    if usuario.Email == "" {
        return goerror.Novo(goerror.CampoObrigatorio).
            ComDetalhe("campo", "email")
    }
    
    if !isValidEmail(usuario.Email) {
        return goerror.Novo(goerror.EmailInvalido).
            ComDetalhe("email_fornecido", usuario.Email)
    }
    
    if len(usuario.Senha) < 8 {
        return goerror.Novo(goerror.SenhaFraca).
            ComDetalhe("tamanho_atual", len(usuario.Senha)).
            ComDetalhe("tamanho_minimo", 8)
    }
    
    return nil
}
```

### Service Layer

```go
func (s *UserService) CriarUsuario(ctx context.Context, dados *CriarUsuarioRequest) (*Usuario, error) {
    // Verificar se email já existe
    existente, err := s.repo.BuscarPorEmail(ctx, dados.Email)
    if err != nil && !goerror.EhNaoEncontrado(err) {
        return nil, goerror.Novo(goerror.ErroBancoDados).
            ComDetalhe("operacao", "buscar_usuario").
            ComTraceID(ctx.Value("trace_id").(string))
    }
    
    if existente != nil {
        return nil, goerror.Novo(goerror.EmailJaExiste).
            ComDetalhe("email", dados.Email)
    }
    
    // Criar usuário
    usuario, err := s.repo.Criar(ctx, dados)
    if err != nil {
        return nil, goerror.Novo(goerror.TransacaoFalhou).
            ComDetalhe("operacao", "criar_usuario").
            ComTraceID(ctx.Value("trace_id").(string))
    }
    
    return usuario, nil
}
```

### Integração com APIs Externas

```go
func (c *PaymentClient) ProcessarPagamento(valor float64) error {
    resp, err := http.Post(c.baseURL+"/payments", "application/json", body)
    if err != nil {
        return goerror.Novo(goerror.ServicoExternoIndisponivel).
            ComDetalhe("servico", "payment_api").
            ComDetalhe("erro_http", err.Error())
    }
    
    defer resp.Body.Close()
    
    switch resp.StatusCode {
    case 401:
        return goerror.Novo(goerror.TokenInvalido).
            ComDetalhe("servico", "payment_api")
    case 429:
        return goerror.Novo(goerror.LimiteAPI).
            ComDetalhe("servico", "payment_api")
    case 422:
        return goerror.Novo(goerror.SaldoInsuficiente).
            ComDetalhe("valor_tentativa", valor)
    }
    
    return nil
}
```

## Informações de Debug

A biblioteca automaticamente captura:

- **Arquivo e linha** onde o erro foi criado
- **Timestamp** da criação do erro
- **Stack trace** implícito através do arquivo/linha
- **Trace ID** para rastreamento distribuído

```go
err := goerror.Novo(goerror.UsuarioNaoEncontrado)
fmt.Printf("Erro criado em: %s:%d\n", err.Arquivo, err.Linha)
fmt.Printf("Timestamp: %s\n", err.Timestamp.Format(time.RFC3339))
```

## Boas Práticas

### 1. Use os Tipos Corretos
```go
// ✅ Correto
return goerror.Novo(goerror.EmailJaExiste)

// ❌ Incorreto - use um tipo mais específico
return goerror.Novo(goerror.CampoInvalido)
```

### 2. Adicione Contexto Relevante
```go
// ✅ Correto
return goerror.Novo(goerror.SaldoInsuficiente).
    ComDetalhe("saldo_atual", conta.Saldo).
    ComDetalhe("valor_necessario", valor)

// ❌ Sem contexto suficiente
return goerror.Novo(goerror.SaldoInsuficiente)
```

### 3. Use Trace IDs em Sistemas Distribuídos
```go
// ✅ Correto
traceID := ctx.Value("trace_id").(string)
return goerror.Novo(goerror.ServicoExternoIndisponivel).
    ComTraceID(traceID)
```

### 4. Trate Erros na Camada Correta
```go
// ✅ Na camada de serviço - transforme erros técnicos em erros de negócio
if err != nil {
    if strings.Contains(err.Error(), "duplicate key") {
        return goerror.Novo(goerror.EmailJaExiste)
    }
    return goerror.Novo(goerror.ErroBancoDados)
}
```

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.
