package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Tipos de erro
const (
	TipoValidacao      = "VALIDATION_ERROR"
	TipoBancoDados     = "DATABASE_ERROR"
	TipoConexao        = "CONNECTION_ERROR"
	TipoNaoEncontrado  = "NOT_FOUND_ERROR"
	TipoAutenticacao   = "AUTHENTICATION_ERROR"
	TipoAutorizacao    = "AUTHORIZATION_ERROR"
	TipoConflito       = "CONFLICT_ERROR"
	TipoInterno        = "INTERNAL_ERROR"
	TipoLimiteExcedido = "RATE_LIMIT_ERROR"
	TipoMetodoInvalido = "METHOD_ERROR"
	TipoArquivo        = "FILE_ERROR"
)

// ErroAPI representa um erro estruturado da API
type ErroAPI struct {
	Tipo       string    `json:"tipo"`
	Codigo     string    `json:"codigo"`
	Mensagem   string    `json:"mensagem"`
	Detalhes   string    `json:"detalhes,omitempty"`
	StatusHTTP int       `json:"-"`
	Timestamp  time.Time `json:"timestamp"`
	CausaRaiz  error     `json:"-"`
}

// Error implementa a interface error
func (e *ErroAPI) Error() string {
	if e.Detalhes != "" {
		return fmt.Sprintf("[%s] %s: %s - %s", e.Tipo, e.Codigo, e.Mensagem, e.Detalhes)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Tipo, e.Codigo, e.Mensagem)
}

// JSON retorna o erro como JSON bytes
func (e *ErroAPI) JSON() []byte {
	data, _ := json.Marshal(e)
	return data
}

// ComDetalhes adiciona detalhes ao erro e retorna o próprio erro para chaining
func (e *ErroAPI) ComDetalhes(detalhes string) *ErroAPI {
	e.Detalhes = detalhes
	return e
}

// ComCausa adiciona a causa raiz do erro
func (e *ErroAPI) ComCausa(causa error) *ErroAPI {
	e.CausaRaiz = causa
	if e.Detalhes == "" && causa != nil {
		e.Detalhes = causa.Error()
	}
	return e
}

// Unwrap retorna a causa raiz para compatibilidade com errors.Is e errors.As
func (e *ErroAPI) Unwrap() error {
	return e.CausaRaiz
}

// novoErro cria um novo erro base
func novoErro(tipo, codigo, mensagem string, statusHTTP int) *ErroAPI {
	return &ErroAPI{
		Tipo:       tipo,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: statusHTTP,
		Timestamp:  time.Now(),
	}
}

// Funções construtoras para cada tipo de erro

// NovoErroValidacao cria um novo erro de validação (400)
func NovoErroValidacao(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoValidacao, codigo, mensagem, http.StatusBadRequest)
}

// NovoErroBancoDados cria um novo erro de banco de dados (500)
func NovoErroBancoDados(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoBancoDados, codigo, mensagem, http.StatusInternalServerError)
}

// NovoErroConexao cria um novo erro de conexão (503)
func NovoErroConexao(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoConexao, codigo, mensagem, http.StatusServiceUnavailable)
}

// NovoErroNaoEncontrado cria um novo erro de recurso não encontrado (404)
func NovoErroNaoEncontrado(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoNaoEncontrado, codigo, mensagem, http.StatusNotFound)
}

// NovoErroAutenticacao cria um novo erro de autenticação (401)
func NovoErroAutenticacao(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoAutenticacao, codigo, mensagem, http.StatusUnauthorized)
}

// NovoErroAutorizacao cria um novo erro de autorização (403)
func NovoErroAutorizacao(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoAutorizacao, codigo, mensagem, http.StatusForbidden)
}

// NovoErroConflito cria um novo erro de conflito (409)
func NovoErroConflito(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoConflito, codigo, mensagem, http.StatusConflict)
}

// NovoErroInterno cria um novo erro interno (500)
func NovoErroInterno(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoInterno, codigo, mensagem, http.StatusInternalServerError)
}

// NovoErroLimiteExcedido cria um novo erro de limite excedido (429)
func NovoErroLimiteExcedido(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoLimiteExcedido, codigo, mensagem, http.StatusTooManyRequests)
}

// NovoErroMetodoInvalido cria um novo erro de método HTTP inválido (405)
func NovoErroMetodoInvalido(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoMetodoInvalido, codigo, mensagem, http.StatusMethodNotAllowed)
}

// NovoErroArquivo cria um novo erro de arquivo (400)
func NovoErroArquivo(codigo, mensagem string) *ErroAPI {
	return novoErro(TipoArquivo, codigo, mensagem, http.StatusBadRequest)
}

// Funções de verificação de tipo de erro

// EhErroValidacao verifica se o erro é do tipo validação
func EhErroValidacao(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoValidacao
	}
	return false
}

// EhErroBancoDados verifica se o erro é do tipo banco de dados
func EhErroBancoDados(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoBancoDados
	}
	return false
}

// EhErroConexao verifica se o erro é do tipo conexão
func EhErroConexao(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoConexao
	}
	return false
}

// EhErroNaoEncontrado verifica se o erro é do tipo não encontrado
func EhErroNaoEncontrado(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoNaoEncontrado
	}
	return false
}

// EhErroAutenticacao verifica se o erro é do tipo autenticação
func EhErroAutenticacao(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoAutenticacao
	}
	return false
}

// EhErroAutorizacao verifica se o erro é do tipo autorização
func EhErroAutorizacao(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoAutorizacao
	}
	return false
}

// EhErroConflito verifica se o erro é do tipo conflito
func EhErroConflito(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoConflito
	}
	return false
}

// EhErroInterno verifica se o erro é do tipo interno
func EhErroInterno(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoInterno
	}
	return false
}

// EhErroLimiteExcedido verifica se o erro é do tipo limite excedido
func EhErroLimiteExcedido(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoLimiteExcedido
	}
	return false
}

// EhErroMetodoInvalido verifica se o erro é do tipo método inválido
func EhErroMetodoInvalido(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoMetodoInvalido
	}
	return false
}

// EhErroArquivo verifica se o erro é do tipo arquivo
func EhErroArquivo(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoArquivo
	}
	return false
}

// ConverterErro converte um erro padrão em ErroAPI
func ConverterErro(err error) *ErroAPI {
	if err == nil {
		return nil
	}

	// Se já é um ErroAPI, retorna como está
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr
	}

	// Converte erro padrão para ErroAPI genérico
	return NovoErroInterno("UNKNOWN_ERROR", "Erro não categorizado").ComCausa(err)
}

// ObterStatusHTTP retorna o status HTTP de um erro
func ObterStatusHTTP(err error) int {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.StatusHTTP
	}
	return http.StatusInternalServerError
}

// ObterTipo retorna o tipo de um erro
func ObterTipo(err error) string {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo
	}
	return "UNKNOWN_ERROR"
}

// ObterCodigo retorna o código de um erro
func ObterCodigo(err error) string {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Codigo
	}
	return "UNKNOWN_CODE"
}
