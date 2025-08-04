// Package errors fornece um sistema de tratamento de erros padronizado para APIs
package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TipoErro representa os diferentes tipos de erros da aplicação
type TipoErro string

const (
	TipoValidacao      TipoErro = "VALIDACAO"
	TipoBancoDados     TipoErro = "BANCO_DADOS"
	TipoConexao        TipoErro = "CONEXAO"
	TipoNaoEncontrado  TipoErro = "NAO_ENCONTRADO"
	TipoMetodoInvalido TipoErro = "METODO_INVALIDO"
	TipoAutenticacao   TipoErro = "AUTENTICACAO"
	TipoAutorizacao    TipoErro = "AUTORIZACAO"
	TipoInterno        TipoErro = "INTERNO"
	TipoConflito       TipoErro = "CONFLITO"
	TipoLimiteExcedido TipoErro = "LIMITE_EXCEDIDO"
)

// ErroAPI representa um erro estruturado da aplicação
type ErroAPI struct {
	Tipo       TipoErro `json:"tipo"`
	Codigo     string   `json:"codigo"`
	Mensagem   string   `json:"mensagem"`
	Detalhes   string   `json:"detalhes,omitempty"`
	StatusHTTP int      `json:"-"`
	Causa      error    `json:"-"`
}

// Error implementa a interface error
func (e *ErroAPI) Error() string {
	if e.Detalhes != "" {
		return fmt.Sprintf("[%s] %s: %s - %s", e.Tipo, e.Codigo, e.Mensagem, e.Detalhes)
	}
	return fmt.Sprintf("[%s] %s: %s", e.Tipo, e.Codigo, e.Mensagem)
}

// JSON retorna a representação JSON do erro para resposta da API
func (e *ErroAPI) JSON() []byte {
	data, _ := json.Marshal(e)
	return data
}

// ComDetalhes adiciona detalhes ao erro
func (e *ErroAPI) ComDetalhes(detalhes string) *ErroAPI {
	e.Detalhes = detalhes
	return e
}

// ComCausa adiciona a causa raiz do erro
func (e *ErroAPI) ComCausa(causa error) *ErroAPI {
	e.Causa = causa
	if e.Detalhes == "" && causa != nil {
		e.Detalhes = causa.Error()
	}
	return e
}

// Construtores de erros específicos

// NovoErroValidacao cria um erro de validação
func NovoErroValidacao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoValidacao,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusBadRequest,
	}
}

// NovoErroBancoDados cria um erro de banco de dados
func NovoErroBancoDados(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoBancoDados,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusInternalServerError,
	}
}

// NovoErroConexao cria um erro de conexão
func NovoErroConexao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoConexao,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusServiceUnavailable,
	}
}

// NovoErroNaoEncontrado cria um erro de recurso não encontrado
func NovoErroNaoEncontrado(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoNaoEncontrado,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusNotFound,
	}
}

// NovoErroMetodoInvalido cria um erro de método não permitido
func NovoErroMetodoInvalido(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoMetodoInvalido,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusMethodNotAllowed,
	}
}

// NovoErroAutenticacao cria um erro de autenticação
func NovoErroAutenticacao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoAutenticacao,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusUnauthorized,
	}
}

// NovoErroAutorizacao cria um erro de autorização
func NovoErroAutorizacao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoAutorizacao,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusForbidden,
	}
}

// NovoErroInterno cria um erro interno do servidor
func NovoErroInterno(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoInterno,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusInternalServerError,
	}
}

// NovoErroConflito cria um erro de conflito (ex: recurso já existe)
func NovoErroConflito(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoConflito,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusConflict,
	}
}

// NovoErroLimiteExcedido cria um erro de limite excedido (rate limit)
func NovoErroLimiteExcedido(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:       TipoLimiteExcedido,
		Codigo:     codigo,
		Mensagem:   mensagem,
		StatusHTTP: http.StatusTooManyRequests,
	}
}

// Erros pré-definidos comuns
var (
	ErrUsuarioNaoEncontrado   = NovoErroNaoEncontrado("USER_NOT_FOUND", "Usuário não encontrado")
	ErrEmailJaExiste          = NovoErroConflito("EMAIL_EXISTS", "Este email já está em uso")
	ErrSenhaInvalida          = NovoErroValidacao("INVALID_PASSWORD", "Senha inválida")
	ErrTokenInvalido          = NovoErroAutenticacao("INVALID_TOKEN", "Token de acesso inválido")
	ErrSemPermissao           = NovoErroAutorizacao("NO_PERMISSION", "Você não tem permissão para esta ação")
	ErrBancoDadosIndisponivel = NovoErroConexao("DB_UNAVAILABLE", "Banco de dados indisponível")
	ErrLimiteRequisicoes      = NovoErroLimiteExcedido("RATE_LIMIT", "Limite de requisições excedido")
)

// Funções utilitárias para verificação de tipos de erro

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

// EhErroNaoEncontrado verifica se o erro é do tipo não encontrado
func EhErroNaoEncontrado(err error) bool {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoNaoEncontrado
	}
	return false
}

// ObterStatusHTTP retorna o status HTTP apropriado para um erro
func ObterStatusHTTP(err error) int {
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.StatusHTTP
	}
	return http.StatusInternalServerError
}

// ConverterErro converte um erro genérico para ErroAPI
func ConverterErro(err error) *ErroAPI {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr
	}

	// Converte erro genérico para erro interno
	return NovoErroInterno("GENERIC_ERROR", "Erro interno do servidor").ComCausa(err)
}
