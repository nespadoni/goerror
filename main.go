package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Importe sua biblioteca de erros
	"github.com/nespadoni/goerror"
)

// Exemplo de estrutura de usuário
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Simulação de um service/repository
type UsuarioService struct {
	// db seria sua conexão com banco
}

// BuscarUsuario simula a busca de um usuário
func (s *UsuarioService) BuscarUsuario(id int) (*Usuario, error) {
	// Simula erro de banco de dados
	if id == 0 {
		return nil, errors.NovoErroBancoDados("DB_CONNECTION_FAILED", "Falha na conexão com banco de dados").
			ComDetalhes("Timeout na conexão após 30 segundos")
	}

	// Simula usuário não encontrado
	if id == 999 {
		return nil, errors.ErrUsuarioNaoEncontrado.ComDetalhes(fmt.Sprintf("ID: %d", id))
	}

	// Retorna usuário simulado
	return &Usuario{
		ID:    id,
		Nome:  "João Silva",
		Email: "joao@email.com",
	}, nil
}

// CriarUsuario simula a criação de um usuário
func (s *UsuarioService) CriarUsuario(usuario *Usuario) error {
	// Validações
	if usuario.Nome == "" {
		return errors.NovoErroValidacao("NOME_OBRIGATORIO", "Nome é obrigatório")
	}

	if usuario.Email == "" {
		return errors.NovoErroValidacao("EMAIL_OBRIGATORIO", "Email é obrigatório")
	}

	// Simula email já existente
	if usuario.Email == "admin@teste.com" {
		return errors.ErrEmailJaExiste.ComDetalhes(fmt.Sprintf("Email: %s", usuario.Email))
	}

	return nil
}

// Middleware para tratamento de erros
func tratarErros(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic capturado: %v", err)

				erroInterno := errors.NovoErroInterno("PANIC_ERROR", "Erro interno inesperado")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(erroInterno.StatusHTTP)
				w.Write(erroInterno.JSON())
			}
		}()

		next(w, r)
	}
}

// Handler para buscar usuário
func buscarUsuarioHandler(service *UsuarioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simula extração do ID da URL
		id := 1 // Normalmente viria de r.URL.Path ou query params

		usuario, err := service.BuscarUsuario(id)
		if err != nil {
			// Converte para ErroAPI se necessário
			apiErr := errors.ConverterErro(err)

			log.Printf("Erro ao buscar usuário: %s", apiErr.Error())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.StatusHTTP)
			w.Write(apiErr.JSON())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(usuario)
	}
}

// Handler para criar usuário
func criarUsuarioHandler(service *UsuarioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := errors.NovoErroMetodoInvalido("METHOD_NOT_ALLOWED", "Método não permitido").
				ComDetalhes(fmt.Sprintf("Método usado: %s, esperado: POST", r.Method))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(err.StatusHTTP)
			w.Write(err.JSON())
			return
		}

		var usuario Usuario
		if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
			apiErr := errors.NovoErroValidacao("JSON_INVALIDO", "JSON inválido").ComCausa(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.StatusHTTP)
			w.Write(apiErr.JSON())
			return
		}

		if err := service.CriarUsuario(&usuario); err != nil {
			apiErr := errors.ConverterErro(err)

			log.Printf("Erro ao criar usuário: %s", apiErr.Error())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.StatusHTTP)
			w.Write(apiErr.JSON())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensagem": "Usuário criado com sucesso"})
	}
}

func main() {
	service := &UsuarioService{}

	http.HandleFunc("/usuarios", tratarErros(criarUsuarioHandler(service)))
	http.HandleFunc("/usuarios/", tratarErros(buscarUsuarioHandler(service)))

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.Serve(nil))
}

// Exemplos de uso em outras situações:

// Exemplo de uso com database/sql
func exemploComBancoDados() {
	db, err := sql.Open("postgres", "connection-string")
	if err != nil {
		// Converte erro de conexão para ErroAPI
		apiErr := errors.NovoErroConexao("DB_CONNECTION_ERROR", "Erro ao conectar com banco").ComCausa(err)
		log.Fatal(apiErr.Error())
	}
	defer db.Close()
}

// Exemplo de validação de campos
func validarCampos(dados map[string]interface{}) error {
	if nome, ok := dados["nome"].(string); !ok || nome == "" {
		return errors.NovoErroValidacao("NOME_INVALIDO", "Nome deve ser uma string não vazia")
	}

	if idade, ok := dados["idade"].(float64); !ok || idade < 0 {
		return errors.NovoErroValidacao("IDADE_INVALIDA", "Idade deve ser um número positivo")
	}

	return nil
}

// Exemplo de verificação de tipo de erro
func exemploVerificacaoTipoErro(err error) {
	if errors.EhErroValidacao(err) {
		log.Println("Erro de validação detectado")
	} else if errors.EhErroBancoDados(err) {
		log.Println("Erro de banco de dados detectado")
	} else if errors.EhErroNaoEncontrado(err) {
		log.Println("Recurso não encontrado")
	}
}
