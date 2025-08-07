package goerror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TipoErro define os tipos de erro suportados
type TipoErro string

const (
	TipoBancoDados      TipoErro = "BANCO_DADOS"
	TipoValidacao       TipoErro = "VALIDACAO"
	TipoNaoEncontrado   TipoErro = "NAO_ENCONTRADO"
	TipoNaoAutorizado   TipoErro = "NAO_AUTORIZADO"
	TipoConexao         TipoErro = "CONEXAO"
	TipoInterno         TipoErro = "INTERNO"
	TipoMetodoInvalido  TipoErro = "METODO_INVALIDO"
	TipoFormatoInvalido TipoErro = "FORMATO_INVALIDO"
)

// ErroAPI representa um erro padronizado da API
type ErroAPI struct {
	Tipo           TipoErro `json:"tipo"`
	CodigoPersonal string   `json:"codigo_personal"`
	Mensagem       string   `json:"mensagem"`
	StatusHTTP     int      `json:"status_http"`
	Detalhes       string   `json:"detalhes,omitempty"`
	ErroOriginal   error    `json:"-"`
}

// Error implementa a interface error
func (e *ErroAPI) Error() string {
	if e.ErroOriginal != nil {
		return fmt.Sprintf("[%s:%s] %s - %s (Original: %v)",
			e.Tipo, e.CodigoPersonal, e.Mensagem, e.Detalhes, e.ErroOriginal)
	}
	return fmt.Sprintf("[%s:%s] %s - %s", e.Tipo, e.CodigoPersonal, e.Mensagem, e.Detalhes)
}

// JSON retorna a representação JSON do erro
func (e *ErroAPI) JSON() []byte {
	resposta := map[string]interface{}{
		"erro":        true,
		"tipo":        e.Tipo,
		"codigo":      e.CodigoPersonal,
		"mensagem":    e.Mensagem,
		"status_http": e.StatusHTTP,
	}

	if e.Detalhes != "" {
		resposta["detalhes"] = e.Detalhes
	}

	jsonBytes, _ := json.Marshal(resposta)
	return jsonBytes
}

// ComDetalhes adiciona detalhes ao erro
func (e *ErroAPI) ComDetalhes(detalhes string) *ErroAPI {
	e.Detalhes = detalhes
	return e
}

// ComCausa adiciona o erro original que causou este erro
func (e *ErroAPI) ComCausa(err error) *ErroAPI {
	e.ErroOriginal = err
	if e.Detalhes == "" && err != nil {
		e.Detalhes = err.Error()
	}
	return e
}

// FUNÇÕES PRINCIPAIS - SEM IF!

// ErroBancoDados - Retorna nil se não houver erro, ou ErroAPI se houver
func ErroBancoDados(codigoPersonal string, err error) error {
	if err == nil {
		return nil
	}

	return &ErroAPI{
		Tipo:           TipoBancoDados,
		CodigoPersonal: codigoPersonal,
		Mensagem:       "Erro na operação do banco de dados",
		StatusHTTP:     http.StatusInternalServerError,
		ErroOriginal:   err,
		Detalhes:       err.Error(),
	}
}

// ErroValidacao - Sempre retorna erro (para validações)
func ErroValidacao(codigoPersonal, mensagem string) error {
	return &ErroAPI{
		Tipo:           TipoValidacao,
		CodigoPersonal: codigoPersonal,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusBadRequest,
	}
}

// ErroNaoEncontrado - Sempre retorna erro
func ErroNaoEncontrado(codigoPersonal, recurso string) error {
	return &ErroAPI{
		Tipo:           TipoNaoEncontrado,
		CodigoPersonal: codigoPersonal,
		Mensagem:       fmt.Sprintf("%s não encontrado", recurso),
		StatusHTTP:     http.StatusNotFound,
	}
}

// ErroNaoAutorizado - Sempre retorna erro
func ErroNaoAutorizado(codigoPersonal, mensagem string) error {
	return &ErroAPI{
		Tipo:           TipoNaoAutorizado,
		CodigoPersonal: codigoPersonal,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusUnauthorized,
	}
}

// ErroConexao - Retorna nil se não houver erro, ou ErroAPI se houver
func ErroConexao(codigoPersonal string, err error) error {
	if err == nil {
		return nil
	}

	return &ErroAPI{
		Tipo:           TipoConexao,
		CodigoPersonal: codigoPersonal,
		Mensagem:       "Erro de conexão",
		StatusHTTP:     http.StatusServiceUnavailable,
		ErroOriginal:   err,
		Detalhes:       err.Error(),
	}
}

// ErroInterno - Retorna nil se não houver erro, ou ErroAPI se houver
func ErroInterno(codigoPersonal string, err error) error {
	if err == nil {
		return nil
	}

	return &ErroAPI{
		Tipo:           TipoInterno,
		CodigoPersonal: codigoPersonal,
		Mensagem:       "Erro interno do servidor",
		StatusHTTP:     http.StatusInternalServerError,
		ErroOriginal:   err,
		Detalhes:       err.Error(),
	}
}

// ErroMetodoInvalido - Sempre retorna erro
func ErroMetodoInvalido(codigoPersonal, metodoEsperado string) error {
	return &ErroAPI{
		Tipo:           TipoMetodoInvalido,
		CodigoPersonal: codigoPersonal,
		Mensagem:       fmt.Sprintf("Método não permitido. Esperado: %s", metodoEsperado),
		StatusHTTP:     http.StatusMethodNotAllowed,
	}
}

// ErroFormatoInvalido - Retorna nil se não houver erro, ou ErroAPI se houver
func ErroFormatoInvalido(codigoPersonal string, err error) error {
	if err == nil {
		return nil
	}

	return &ErroAPI{
		Tipo:           TipoFormatoInvalido,
		CodigoPersonal: codigoPersonal,
		Mensagem:       "Formato de dados inválido",
		StatusHTTP:     http.StatusBadRequest,
		ErroOriginal:   err,
		Detalhes:       err.Error(),
	}
}

// VERSÕES PERSONALIZÁVEIS

// NovoErroBancoDados - Para criar com mensagem personalizada
func NovoErroBancoDados(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoBancoDados,
		CodigoPersonal: codigo,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusInternalServerError,
	}
}

// NovoErroValidacao - Para criar com mensagem personalizada
func NovoErroValidacao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoValidacao,
		CodigoPersonal: codigo,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusBadRequest,
	}
}

// NovoErroNaoEncontrado - Para criar com mensagem personalizada
func NovoErroNaoEncontrado(codigo, recurso string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoNaoEncontrado,
		CodigoPersonal: codigo,
		Mensagem:       fmt.Sprintf("%s não encontrado", recurso),
		StatusHTTP:     http.StatusNotFound,
	}
}

// NovoErroConexao - Para criar com mensagem personalizada
func NovoErroConexao(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoConexao,
		CodigoPersonal: codigo,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusServiceUnavailable,
	}
}

// NovoErroInterno - Para criar com mensagem personalizada
func NovoErroInterno(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoInterno,
		CodigoPersonal: codigo,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusInternalServerError,
	}
}

// NovoErroMetodoInvalido - Para criar com mensagem personalizada
func NovoErroMetodoInvalido(codigo, mensagem string) *ErroAPI {
	return &ErroAPI{
		Tipo:           TipoMetodoInvalido,
		CodigoPersonal: codigo,
		Mensagem:       mensagem,
		StatusHTTP:     http.StatusMethodNotAllowed,
	}
}

// ERROS PRÉ-DEFINIDOS
var (
	ErrUsuarioNaoEncontrado = &ErroAPI{
		Tipo:           TipoNaoEncontrado,
		CodigoPersonal: "USER_NOT_FOUND",
		Mensagem:       "Usuário não encontrado",
		StatusHTTP:     http.StatusNotFound,
	}

	ErrEmailJaExiste = &ErroAPI{
		Tipo:           TipoValidacao,
		CodigoPersonal: "EMAIL_EXISTS",
		Mensagem:       "Email já está em uso",
		StatusHTTP:     http.StatusConflict,
	}

	ErrTokenInvalido = &ErroAPI{
		Tipo:           TipoNaoAutorizado,
		CodigoPersonal: "INVALID_TOKEN",
		Mensagem:       "Token de autenticação inválido",
		StatusHTTP:     http.StatusUnauthorized,
	}

	ErrPermissaoNegada = &ErroAPI{
		Tipo:           TipoNaoAutorizado,
		CodigoPersonal: "PERMISSION_DENIED",
		Mensagem:       "Permissão negada para esta operação",
		StatusHTTP:     http.StatusForbidden,
	}
)

// UTILITÁRIOS

// ConverterErro converte um erro comum para ErroAPI
func ConverterErro(err error) *ErroAPI {
	if err == nil {
		return nil
	}

	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr
	}

	return &ErroAPI{
		Tipo:           TipoInterno,
		CodigoPersonal: "UNKNOWN_ERROR",
		Mensagem:       "Erro interno não identificado",
		StatusHTTP:     http.StatusInternalServerError,
		ErroOriginal:   err,
		Detalhes:       err.Error(),
	}
}

// Helper para responder requisições HTTP com erro
func ResponderComErro(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	apiErr := ConverterErro(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.StatusHTTP)
	w.Write(apiErr.JSON())
}

// Funções de verificação de tipo de erro
func EhErroValidacao(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoValidacao
	}
	return false
}

func EhErroBancoDados(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoBancoDados
	}
	return false
}

func EhErroNaoEncontrado(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoNaoEncontrado
	}
	return false
}

func EhErroNaoAutorizado(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoNaoAutorizado
	}
	return false
}

func EhErroConexao(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*ErroAPI); ok {
		return apiErr.Tipo == TipoConexao
	}
	return false
}
