# GoError - Biblioteca de Tratamento de Erros para Go

Uma biblioteca simples e eficiente para padronizar o tratamento de erros em aplicações Go, com tipos pré-definidos, códigos HTTP automáticos e serialização JSON nativa.

## Características

- ✅ **8 tipos de erro** organizados por categoria
- ✅ **Códigos HTTP automáticos** baseados no tipo do erro
- ✅ **Serialização JSON** nativa
- ✅ **Encadeamento de métodos** para adicionar detalhes
- ✅ **Erros pré-definidos** para casos comuns
- ✅ **Funções utilitárias** para verificação de tipos
- ✅ **Integração fácil** com handlers HTTP

## Instalação

```bash
go get github.com/nespadoni/goerror
```

## Estrutura do Erro

```go
type ErroAPI struct {
Tipo           TipoErro `json:"tipo"`
CodigoPersonal string   `json:"codigo_personal"`
Mensagem       string   `json:"mensagem"`
StatusHTTP     int      `json:"status_http"`
Detalhes       string   `json:"detalhes,omitempty"`
ErroOriginal   error    `json:"-"`
}
```

## Tipos de Erro Disponíveis

- `TipoBancoDados` - Erros relacionados ao banco de dados (500)
- `TipoValidacao` - Erros de validação de dados (400)
- `TipoNaoEncontrado` - Recursos não encontrados (404)
- `TipoNaoAutorizado` - Problemas de autenticação/autorização (401/403)
- `TipoConexao` - Erros de conectividade (503)
- `TipoInterno` - Erros internos do servidor (500)
- `TipoMetodoInvalido` - Método HTTP não permitido (405)
- `TipoFormatoInvalido` - Formato de dados inválido (400)

## Funções Principais

### ErroBancoDados
Retorna `nil` se não houver erro, ou `ErroAPI` se houver erro de banco de dados.

```go
func ExemploBancoDados() error {
err := db.Query("SELECT * FROM usuarios")
return goerror.ErroBancoDados("DB_QUERY_FAILED", err)
}

// Uso:
if err := ExemploBancoDados(); err != nil {
// Trata o erro
fmt.Println(err.Error())
// Output: [BANCO_DADOS:DB_QUERY_FAILED] Erro na operação do banco de dados - connection refused
}
```

### ErroValidacao
Sempre retorna erro para casos de validação de dados.

```go
func ValidarIdade(idade int) error {
if idade < 18 {
return goerror.ErroValidacao("IDADE_INSUFICIENTE", "Idade deve ser maior que 18 anos")
}
return nil
}

// Uso:
if err := ValidarIdade(16); err != nil {
fmt.Println(err.Error())
// Output: [VALIDACAO:IDADE_INSUFICIENTE] Idade deve ser maior que 18 anos
}
```

### ErroNaoEncontrado
Sempre retorna erro para recursos não encontrados.

```go
func BuscarUsuario(id int) error {
// Simula busca que não encontrou o usuário
return goerror.ErroNaoEncontrado("USER_NOT_FOUND", "Usuário")
}

// Uso:
if err := BuscarUsuario(123); err != nil {
fmt.Println(err.Error())
// Output: [NAO_ENCONTRADO:USER_NOT_FOUND] Usuário não encontrado
}
```

### ErroNaoAutorizado
Sempre retorna erro para problemas de autenticação/autorização.

```go
func VerificarPermissao(token string) error {
if token == "" {
return goerror.ErroNaoAutorizado("MISSING_TOKEN", "Token de acesso obrigatório")
}
return nil
}

// Uso:
if err := VerificarPermissao(""); err != nil {
fmt.Println(err.Error())
// Output: [NAO_AUTORIZADO:MISSING_TOKEN] Token de acesso obrigatório
}
```

### ErroConexao
Retorna `nil` se não houver erro, ou `ErroAPI` se houver erro de conexão.

```go
func ConectarAPI() error {
_, err := http.Get("https://api.externa.com/dados")
return goerror.ErroConexao("API_CONNECTION_FAILED", err)
}

// Uso:
if err := ConectarAPI(); err != nil {
fmt.Println(err.Error())
// Output: [CONEXAO:API_CONNECTION_FAILED] Erro de conexão - dial tcp: connection refused
}
```

### ErroInterno
Retorna `nil` se não houver erro, ou `ErroAPI` se houver erro interno.

```go
func ProcessarDados() error {
err := errors.New("falha no processamento interno")
return goerror.ErroInterno("PROCESSING_FAILED", err)
}

// Uso:
if err := ProcessarDados(); err != nil {
fmt.Println(err.Error())
// Output: [INTERNO:PROCESSING_FAILED] Erro interno do servidor - falha no processamento interno
}
```

### ErroMetodoInvalido
Sempre retorna erro para métodos HTTP não permitidos.

```go
func ValidarMetodo(metodo string) error {
if metodo != "POST" {
return goerror.ErroMetodoInvalido("WRONG_METHOD", "POST")
}
return nil
}

// Uso:
if err := ValidarMetodo("GET"); err != nil {
fmt.Println(err.Error())
// Output: [METODO_INVALIDO:WRONG_METHOD] Método não permitido. Esperado: POST
}
```

### ErroFormatoInvalido
Retorna `nil` se não houver erro, ou `ErroAPI` se houver erro de formato.

```go
func ParsearJSON(data []byte) error {
var result map[string]interface{}
err := json.Unmarshal(data, &result)
return goerror.ErroFormatoInvalido("JSON_PARSE_ERROR", err)
}

// Uso:
if err := ParsearJSON([]byte("json inválido")); err != nil {
fmt.Println(err.Error())
// Output: [FORMATO_INVALIDO:JSON_PARSE_ERROR] Formato de dados inválido - invalid character 'j'...
}
```

## Funções de Criação Personalizada

### NovoErroBancoDados
Cria um novo erro de banco de dados com mensagem personalizada.

```go
func ExemploPersonalizado() error {
return goerror.NovoErroBancoDados("CUSTOM_DB_ERROR", "Falha na conexão com PostgreSQL")
}

// Uso:
err := ExemploPersonalizado()
fmt.Println(err.Error())
// Output: [BANCO_DADOS:CUSTOM_DB_ERROR] Falha na conexão com PostgreSQL
```

### NovoErroValidacao
Cria um novo erro de validação com mensagem personalizada.

```go
func ValidarEmail(email string) error {
if !strings.Contains(email, "@") {
return goerror.NovoErroValidacao("INVALID_EMAIL", "Email deve conter @")
}
return nil
}

// Uso:
if err := ValidarEmail("emailinvalido"); err != nil {
fmt.Println(err.Error())
// Output: [VALIDACAO:INVALID_EMAIL] Email deve conter @
}
```

### NovoErroNaoEncontrado
Cria um novo erro de recurso não encontrado.

```go
func BuscarProduto(id string) error {
return goerror.NovoErroNaoEncontrado("PRODUCT_NOT_FOUND", "Produto")
}

// Uso:
err := BuscarProduto("123")
fmt.Println(err.Error())
// Output: [NAO_ENCONTRADO:PRODUCT_NOT_FOUND] Produto não encontrado
```

### NovoErroConexao
Cria um novo erro de conexão com mensagem personalizada.

```go
func TestarConexao() error {
return goerror.NovoErroConexao("REDIS_CONN_FAILED", "Não foi possível conectar ao Redis")
}

// Uso:
err := TestarConexao()
fmt.Println(err.Error())
// Output: [CONEXAO:REDIS_CONN_FAILED] Não foi possível conectar ao Redis
```

### NovoErroInterno
Cria um novo erro interno com mensagem personalizada.

```go
func ProcessarConfig() error {
return goerror.NovoErroInterno("CONFIG_LOAD_FAILED", "Falha ao carregar configurações")
}

// Uso:
err := ProcessarConfig()
fmt.Println(err.Error())
// Output: [INTERNO:CONFIG_LOAD_FAILED] Falha ao carregar configurações
```

### NovoErroMetodoInvalido
Cria um novo erro de método inválido com mensagem personalizada.

```go
func ValidarMetodoCustom() error {
return goerror.NovoErroMetodoInvalido("ONLY_PUT_ALLOWED", "Apenas PUT é permitido neste endpoint")
}

// Uso:
err := ValidarMetodoCustom()
fmt.Println(err.Error())
// Output: [METODO_INVALIDO:ONLY_PUT_ALLOWED] Apenas PUT é permitido neste endpoint
```

## Métodos de Encadeamento

### ComDetalhes
Adiciona detalhes adicionais ao erro.

```go
func ExemploComDetalhes() error {
return goerror.NovoErroValidacao("FIELD_REQUIRED", "Campo obrigatório").
ComDetalhes("Campo 'email' é obrigatório para cadastro")
}

// Uso:
err := ExemploComDetalhes()
fmt.Println(err.Error())
// Output: [VALIDACAO:FIELD_REQUIRED] Campo obrigatório - Campo 'email' é obrigatório para cadastro
```

### ComCausa
Adiciona o erro original que causou este erro.

```go
func ExemploComCausa() error {
originalErr := errors.New("connection timeout")
return goerror.NovoErroBancoDados("DB_TIMEOUT", "Timeout na consulta").
ComCausa(originalErr)
}

// Uso:
err := ExemploComCausa()
fmt.Println(err.Error())
// Output: [BANCO_DADOS:DB_TIMEOUT] Timeout na consulta - connection timeout (Original: connection timeout)
```

## Erros Pré-definidos

### ErrUsuarioNaoEncontrado
```go
func BuscarUsuarioExemplo(id int) error {
// Simula usuário não encontrado
return goerror.ErrUsuarioNaoEncontrado
}

// Uso:
if err := BuscarUsuarioExemplo(123); err != nil {
fmt.Println(err.Error())
// Output: [NAO_ENCONTRADO:USER_NOT_FOUND] Usuário não encontrado
}
```

### ErrEmailJaExiste
```go
func CadastrarUsuario(email string) error {
// Simula email já existente
return goerror.ErrEmailJaExiste
}

// Uso:
if err := CadastrarUsuario("test@test.com"); err != nil {
fmt.Println(err.Error())
// Output: [VALIDACAO:EMAIL_EXISTS] Email já está em uso
}
```

### ErrTokenInvalido
```go
func ValidarToken(token string) error {
if token == "invalid" {
return goerror.ErrTokenInvalido
}
return nil
}

// Uso:
if err := ValidarToken("invalid"); err != nil {
fmt.Println(err.Error())
// Output: [NAO_AUTORIZADO:INVALID_TOKEN] Token de autenticação inválido
}
```

### ErrPermissaoNegada
```go
func VerificarAdmin(userRole string) error {
if userRole != "admin" {
return goerror.ErrPermissaoNegada
}
return nil
}

// Uso:
if err := VerificarAdmin("user"); err != nil {
fmt.Println(err.Error())
// Output: [NAO_AUTORIZADO:PERMISSION_DENIED] Permissão negada para esta operação
}
```

## Funções Utilitárias

### ConverterErro
Converte um erro comum para ErroAPI.

```go
func ExemploConverter() error {
err := errors.New("erro desconhecido")
return goerror.ConverterErro(err)
}

// Uso:
err := ExemploConverter()
if apiErr := goerror.ConverterErro(err); apiErr != nil {
fmt.Println(apiErr.Error())
// Output: [INTERNO:UNKNOWN_ERROR] Erro interno não identificado - erro desconhecido
}
```

### ResponderComErro
Helper para responder requisições HTTP com erro.

```go
func HandlerExemplo(w http.ResponseWriter, r *http.Request) {
err := goerror.ErroValidacao("INVALID_INPUT", "Dados inválidos")
goerror.ResponderComErro(w, err)
// Resposta HTTP 400 com JSON do erro
}

// JSON retornado:
// {
//   "erro": true,
//   "tipo": "VALIDACAO",
//   "codigo": "INVALID_INPUT",
//   "mensagem": "Dados inválidos",
//   "status_http": 400
// }
```

## Funções de Verificação de Tipo

### EhErroValidacao
Verifica se o erro é do tipo validação.

```go
func TratarErro(err error) {
if goerror.EhErroValidacao(err) {
log.Println("Erro de validação:", err.Error())
}
}

// Uso:
err := goerror.ErroValidacao("FIELD_INVALID", "Campo inválido")
TratarErro(err)
// Output: Erro de validação: [VALIDACAO:FIELD_INVALID] Campo inválido
```

### EhErroBancoDados
Verifica se o erro é do tipo banco de dados.

```go
func TratarErroBD(err error) {
if goerror.EhErroBancoDados(err) {
log.Println("Problema no banco:", err.Error())
}
}

// Uso:
err := goerror.NovoErroBancoDados("QUERY_FAILED", "Falha na query")
TratarErroBD(err)
// Output: Problema no banco: [BANCO_DADOS:QUERY_FAILED] Falha na query
```

### EhErroNaoEncontrado
Verifica se o erro é do tipo não encontrado.

```go
func TratarNotFound(err error) {
if goerror.EhErroNaoEncontrado(err) {
log.Println("Recurso não encontrado:", err.Error())
}
}

// Uso:
err := goerror.ErrUsuarioNaoEncontrado
TratarNotFound(err)
// Output: Recurso não encontrado: [NAO_ENCONTRADO:USER_NOT_FOUND] Usuário não encontrado
```

### EhErroNaoAutorizado
Verifica se o erro é do tipo não autorizado.

```go
func TratarAuth(err error) {
if goerror.EhErroNaoAutorizado(err) {
log.Println("Problema de autorização:", err.Error())
}
}

// Uso:
err := goerror.ErrTokenInvalido
TratarAuth(err)
// Output: Problema de autorização: [NAO_AUTORIZADO:INVALID_TOKEN] Token de autenticação inválido
```

### EhErroConexao
Verifica se o erro é do tipo conexão.

```go
func TratarConexao(err error) {
if goerror.EhErroConexao(err) {
log.Println("Problema de conectividade:", err.Error())
}
}

// Uso:
connErr := errors.New("timeout")
err := goerror.ErroConexao("CONN_TIMEOUT", connErr)
TratarConexao(err)
// Output: Problema de conectividade: [CONEXAO:CONN_TIMEOUT] Erro de conexão - timeout
```

## Serialização JSON

### Método JSON()
Retorna a representação JSON do erro.

```go
func ExemploJSON() {
err := goerror.NovoErroValidacao("EMAIL_REQUIRED", "Email é obrigatório").
ComDetalhes("O campo email não pode estar vazio")

jsonBytes := err.JSON()
fmt.Println(string(jsonBytes))
}

// Output:
// {
//   "erro": true,
//   "tipo": "VALIDACAO",
//   "codigo": "EMAIL_REQUIRED", 
//   "mensagem": "Email é obrigatório",
//   "status_http": 400,
//   "detalhes": "O campo email não pode estar vazio"
// }
```

## Exemplo Completo em Handler HTTP

```go
func CriarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
// Validar método
if err := goerror.ErroMetodoInvalido("WRONG_METHOD", "POST"); r.Method != "POST" && err != nil {
goerror.ResponderComErro(w, err)
return
}

// Parse JSON
var dados map[string]interface{}
if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
erro := goerror.ErroFormatoInvalido("JSON_PARSE_ERROR", err)
goerror.ResponderComErro(w, erro)
return
}

// Validar email
email, ok := dados["email"].(string)
if !ok || email == "" {
erro := goerror.ErroValidacao("EMAIL_REQUIRED", "Email é obrigatório")
goerror.ResponderComErro(w, erro)
return
}

// Verificar se usuário já existe
if emailExiste(email) {
goerror.ResponderComErro(w, goerror.ErrEmailJaExiste)
return
}

// Simular erro de banco
if err := salvarUsuario(dados); err != nil {
erro := goerror.ErroBancoDados("USER_SAVE_FAILED", err)
goerror.ResponderComErro(w, erro)
return
}

// Sucesso
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}
```

## Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.