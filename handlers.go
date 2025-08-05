package main

import "strconv"

// Erros de Validação
var (
	ErrCPFInvalido         = NovoErroValidacao("CPF_INVALIDO", "CPF inválido")
	ErrCNPJInvalido        = NovoErroValidacao("CNPJ_INVALIDO", "CNPJ inválido")
	ErrEmailInvalido       = NovoErroValidacao("EMAIL_INVALIDO", "Email inválido")
	ErrTelefoneInvalido    = NovoErroValidacao("TELEFONE_INVALIDO", "Telefone inválido")
	ErrCEPInvalido         = NovoErroValidacao("CEP_INVALIDO", "CEP inválido")
	ErrSenhaInvalida       = NovoErroValidacao("SENHA_INVALIDA", "Senha inválida")
	ErrUUIDInvalido        = NovoErroValidacao("UUID_INVALIDO", "UUID inválido")
	ErrIDInvalido          = NovoErroValidacao("ID_INVALIDO", "ID inválido")
	ErrCampoObrigatorio    = NovoErroValidacao("CAMPO_OBRIGATORIO", "Campo obrigatório não informado")
	ErrFormatoJSON         = NovoErroValidacao("JSON_INVALIDO", "Formato JSON inválido")
	ErrTamanhoInvalido     = NovoErroValidacao("TAMANHO_INVALIDO", "Tamanho inválido")
	ErrValorDuplicado      = NovoErroValidacao("VALOR_DUPLICADO", "Valor duplicado")
	ErrValorInvalido       = NovoErroValidacao("VALOR_INVALIDO", "Valor inválido")
	ErrFormatoDataInvalida = NovoErroValidacao("DATA_INVALIDA", "Formato de data inválido")
	ErrDataFutura          = NovoErroValidacao("DATA_FUTURA", "Data não pode ser no futuro")
	ErrDataPassada         = NovoErroValidacao("DATA_PASSADA", "Data não pode ser no passado")
	ErrIdadeInvalida       = NovoErroValidacao("IDADE_INVALIDA", "Idade inválida")
)

// Erros de Banco de Dados
var (
	ErrConexaoBanco        = NovoErroBancoDados("DB_CONNECTION_ERROR", "Erro de conexão com banco de dados")
	ErrTimeoutBanco        = NovoErroBancoDados("DB_TIMEOUT", "Timeout na operação do banco de dados")
	ErrQueryInvalida       = NovoErroBancoDados("DB_INVALID_QUERY", "Query inválida")
	ErrConstraintViolation = NovoErroBancoDados("DB_CONSTRAINT_VIOLATION", "Violação de constraint do banco")
	ErrTransacaoFalhou     = NovoErroBancoDados("DB_TRANSACTION_FAILED", "Falha na transação")
	ErrBancoIndisponivel   = NovoErroBancoDados("DB_UNAVAILABLE", "Banco de dados indisponível")
	ErrTabelaNaoEncontrada = NovoErroBancoDados("DB_TABLE_NOT_FOUND", "Tabela não encontrada")
	ErrColunaNaoEncontrada = NovoErroBancoDados("DB_COLUMN_NOT_FOUND", "Coluna não encontrada")
	ErrChaveEstrangeira    = NovoErroBancoDados("DB_FOREIGN_KEY_ERROR", "Erro de chave estrangeira")
	ErrChavePrimaria       = NovoErroBancoDados("DB_PRIMARY_KEY_ERROR", "Erro de chave primária")
)

// Erros de Conexão
var (
	ErrConexaoRecusada        = NovoErroConexao("CONNECTION_REFUSED", "Conexão recusada")
	ErrTimeoutConexao         = NovoErroConexao("CONNECTION_TIMEOUT", "Timeout de conexão")
	ErrConexaoInterrompida    = NovoErroConexao("CONNECTION_INTERRUPTED", "Conexão interrompida")
	ErrServicoIndisponivel    = NovoErroConexao("SERVICE_UNAVAILABLE", "Serviço indisponível")
	ErrRedeInacessivel        = NovoErroConexao("NETWORK_UNREACHABLE", "Rede inacessível")
	ErrDNSFalhou              = NovoErroConexao("DNS_RESOLUTION_FAILED", "Falha na resolução DNS")
	ErrSSLFalhou              = NovoErroConexao("SSL_HANDSHAKE_FAILED", "Falha no handshake SSL")
	ErrAPIExternaIndisponivel = NovoErroConexao("EXTERNAL_API_UNAVAILABLE", "API externa indisponível")
)

// Erros de Não Encontrado
var (
	ErrUsuarioNaoEncontrado   = NovoErroNaoEncontrado("USER_NOT_FOUND", "Usuário não encontrado")
	ErrRecursoNaoEncontrado   = NovoErroNaoEncontrado("RESOURCE_NOT_FOUND", "Recurso não encontrado")
	ErrArquivoNaoEncontrado   = NovoErroNaoEncontrado("FILE_NOT_FOUND", "Arquivo não encontrado")
	ErrPaginaNaoEncontrada    = NovoErroNaoEncontrado("PAGE_NOT_FOUND", "Página não encontrada")
	ErrEndpointNaoEncontrado  = NovoErroNaoEncontrado("ENDPOINT_NOT_FOUND", "Endpoint não encontrado")
	ErrClienteNaoEncontrado   = NovoErroNaoEncontrado("CLIENT_NOT_FOUND", "Cliente não encontrado")
	ErrProdutoNaoEncontrado   = NovoErroNaoEncontrado("PRODUCT_NOT_FOUND", "Produto não encontrado")
	ErrPedidoNaoEncontrado    = NovoErroNaoEncontrado("ORDER_NOT_FOUND", "Pedido não encontrado")
	ErrCategoriaNaoEncontrada = NovoErroNaoEncontrado("CATEGORY_NOT_FOUND", "Categoria não encontrada")
	ErrEmpresaNaoEncontrada   = NovoErroNaoEncontrado("COMPANY_NOT_FOUND", "Empresa não encontrada")
)

// Erros de Autenticação
var (
	ErrTokenInvalido         = NovoErroAutenticacao("INVALID_TOKEN", "Token inválido")
	ErrTokenExpirado         = NovoErroAutenticacao("TOKEN_EXPIRED", "Token expirado")
	ErrCredenciaisInvalidas  = NovoErroAutenticacao("INVALID_CREDENTIALS", "Credenciais inválidas")
	ErrSessaoExpirada        = NovoErroAutenticacao("SESSION_EXPIRED", "Sessão expirada")
	ErrUsuarioInativo        = NovoErroAutenticacao("USER_INACTIVE", "Usuário inativo")
	ErrContaBloqueada        = NovoErroAutenticacao("ACCOUNT_BLOCKED", "Conta bloqueada")
	ErrMuitasTentativas      = NovoErroAutenticacao("TOO_MANY_ATTEMPTS", "Muitas tentativas de login")
	ErrAutenticacaoRequerida = NovoErroAutenticacao("AUTHENTICATION_REQUIRED", "Autenticação necessária")
)

// Erros de Autorização
var (
	ErrSemPermissao         = NovoErroAutorizacao("NO_PERMISSION", "Sem permissão para esta ação")
	ErrAcessoNegado         = NovoErroAutorizacao("ACCESS_DENIED", "Acesso negado")
	ErrNivelInsuficiente    = NovoErroAutorizacao("INSUFFICIENT_LEVEL", "Nível de acesso insuficiente")
	ErrRecursoProtegido     = NovoErroAutorizacao("PROTECTED_RESOURCE", "Recurso protegido")
	ErrOperacaoNaoPermitida = NovoErroAutorizacao("OPERATION_NOT_ALLOWED", "Operação não permitida")
	ErrAdminRequerido       = NovoErroAutorizacao("ADMIN_REQUIRED", "Privilégios de administrador necessários")
)

// Erros de Conflito
var (
	ErrEmailJaExiste       = NovoErroConflito("EMAIL_EXISTS", "Este email já está em uso")
	ErrCPFJaExiste         = NovoErroConflito("CPF_EXISTS", "Este CPF já está cadastrado")
	ErrCNPJJaExiste        = NovoErroConflito("CNPJ_EXISTS", "Este CNPJ já está cadastrado")
	ErrUsuarioJaExiste     = NovoErroConflito("USER_EXISTS", "Usuário já existe")
	ErrRecursoJaExiste     = NovoErroConflito("RESOURCE_EXISTS", "Recurso já existe")
	ErrNomeJaUtilizado     = NovoErroConflito("NAME_ALREADY_USED", "Nome já está sendo utilizado")
	ErrCodigoJaExiste      = NovoErroConflito("CODE_EXISTS", "Código já existe")
	ErrVersaoConflito      = NovoErroConflito("VERSION_CONFLICT", "Conflito de versão")
	ErrOperacaoEmAndamento = NovoErroConflito("OPERATION_IN_PROGRESS", "Operação já em andamento")
)

// Erros Internos
var (
	ErrServidorInterno      = NovoErroInterno("INTERNAL_SERVER_ERROR", "Erro interno do servidor")
	ErrConfiguracaoInvalida = NovoErroInterno("INVALID_CONFIGURATION", "Configuração inválida")
	ErrMemoriaInsuficiente  = NovoErroInterno("INSUFFICIENT_MEMORY", "Memória insuficiente")
	ErrArquivoSistema       = NovoErroInterno("FILE_SYSTEM_ERROR", "Erro no sistema de arquivos")
	ErrProcessamentoFalhou  = NovoErroInterno("PROCESSING_FAILED", "Falha no processamento")
	ErrManutencao           = NovoErroInterno("MAINTENANCE_MODE", "Sistema em manutenção")
)

// Erros de Limite Excedido
var (
	ErrLimiteRequisicoes   = NovoErroLimiteExcedido("RATE_LIMIT", "Limite de requisições excedido")
	ErrLimiteUpload        = NovoErroLimiteExcedido("UPLOAD_LIMIT", "Limite de upload excedido")
	ErrLimiteUsuarios      = NovoErroLimiteExcedido("USER_LIMIT", "Limite de usuários excedido")
	ErrLimiteBandwidth     = NovoErroLimiteExcedido("BANDWIDTH_LIMIT", "Limite de bandwidth excedido")
	ErrLimiteArmazenamento = NovoErroLimiteExcedido("STORAGE_LIMIT", "Limite de armazenamento excedido")
	ErrLimiteConexoes      = NovoErroLimiteExcedido("CONNECTION_LIMIT", "Limite de conexões excedido")
)

// Erros de Método HTTP
var (
	ErrMetodoNaoPermitido = NovoErroMetodoInvalido("METHOD_NOT_ALLOWED", "Método HTTP não permitido")
	ErrMetodoNaoSuportado = NovoErroMetodoInvalido("METHOD_NOT_SUPPORTED", "Método HTTP não suportado")
	ErrPostRequerido      = NovoErroMetodoInvalido("POST_REQUIRED", "Método POST é obrigatório")
	ErrGetRequerido       = NovoErroMetodoInvalido("GET_REQUIRED", "Método GET é obrigatório")
	ErrPutRequerido       = NovoErroMetodoInvalido("PUT_REQUIRED", "Método PUT é obrigatório")
	ErrDeleteRequerido    = NovoErroMetodoInvalido("DELETE_REQUIRED", "Método DELETE é obrigatório")
)

// Erros de Arquivo
var (
	ErrArquivoMuitoGrande  = NovoErroArquivo("FILE_TOO_LARGE", "Arquivo muito grande")
	ErrTipoArquivoInvalido = NovoErroArquivo("INVALID_FILE_TYPE", "Tipo de arquivo inválido")
	ErrArquivoCorrempido   = NovoErroArquivo("CORRUPTED_FILE", "Arquivo corrompido")
	ErrFormatoArquivo      = NovoErroArquivo("INVALID_FILE_FORMAT", "Formato de arquivo inválido")
	ErrUploadFalhou        = NovoErroArquivo("UPLOAD_FAILED", "Falha no upload do arquivo")
)

// Funções auxiliares para criar erros com contexto específico

// CriarErroRecursoNaoEncontrado cria um erro personalizado de recurso não encontrado
func CriarErroRecursoNaoEncontrado(recurso, identificador string) *ErroAPI {
	return NovoErroNaoEncontrado("RESOURCE_NOT_FOUND", recurso+" não encontrado").
		ComDetalhes("Identificador: " + identificador)
}

// CriarErroValidacaoCampo cria um erro de validação para um campo específico
func CriarErroValidacaoCampo(campo, motivo string) *ErroAPI {
	return NovoErroValidacao("FIELD_VALIDATION_ERROR", "Erro de validação no campo").
		ComDetalhes("Campo: " + campo + ", Motivo: " + motivo)
}

// CriarErroConexaoServico cria um erro de conexão para um serviço específico
func CriarErroConexaoServico(servico, detalhes string) *ErroAPI {
	return NovoErroConexao("SERVICE_CONNECTION_ERROR", "Erro de conexão com "+servico).
		ComDetalhes(detalhes)
}

// CriarErroBancoOperacao cria um erro de banco para uma operação específica
func CriarErroBancoOperacao(operacao, tabela string, causa error) *ErroAPI {
	return NovoErroBancoDados("DB_OPERATION_ERROR", "Erro na operação de "+operacao).
		ComDetalhes("Tabela: " + tabela).
		ComCausa(causa)
}

// CriarErroPermissaoOperacao cria um erro de permissão para uma operação específica
func CriarErroPermissaoOperacao(operacao, recurso string) *ErroAPI {
	return NovoErroAutorizacao("OPERATION_NOT_AUTHORIZED", "Operação não autorizada").
		ComDetalhes("Operação: " + operacao + ", Recurso: " + recurso)
}

// CriarErroLimiteEspecifico cria um erro de limite excedido específico
func CriarErroLimiteEspecifico(tipoLimite string, limite, atual int) *ErroAPI {
	return NovoErroLimiteExcedido("LIMIT_EXCEEDED", "Limite de "+tipoLimite+" excedido").
		ComDetalhes("Limite: " + strconv.Itoa(limite) + ", Atual: " + strconv.Itoa(atual))
}
