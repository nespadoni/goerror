package goerror

import (
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Expressões regulares para validação
var (
	regexCPF      = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	regexCNPJ     = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?\d{4}-?\d{2}$`)
	regexTelefone = regexp.MustCompile(`^(\(\d{2}\)\s?|\d{2}\s?)?9?\d{4}-?\d{4}$`)
	regexCEP      = regexp.MustCompile(`^\d{5}-?\d{3}$`)
	regexUUID     = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

// ValidarCPF valida se um CPF é válido
func ValidarCPF(cpf string) error {
	if cpf == "" {
		return NovoErroValidacao("CPF_VAZIO", "CPF não pode estar vazio")
	}

	// Remove formatação
	cpfLimpo := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(cpf, ".", ""), "-", ""), " ", "")

	if len(cpfLimpo) != 11 {
		return NovoErroValidacao("CPF_TAMANHO_INVALIDO", "CPF deve ter 11 dígitos").
			ComDetalhes("CPF fornecido: " + cpf)
	}

	// Verifica se todos os dígitos são iguais
	todosIguais := true
	for i := 1; i < len(cpfLimpo); i++ {
		if cpfLimpo[i] != cpfLimpo[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return NovoErroValidacao("CPF_DIGITOS_IGUAIS", "CPF não pode ter todos os dígitos iguais").
			ComDetalhes("CPF fornecido: " + cpf)
	}

	// Validação dos dígitos verificadores
	soma := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpfLimpo[i]))
		soma += num * (10 - i)
	}
	digito1 := 11 - (soma % 11)
	if digito1 >= 10 {
		digito1 = 0
	}

	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpfLimpo[i]))
		soma += num * (11 - i)
	}
	digito2 := 11 - (soma % 11)
	if digito2 >= 10 {
		digito2 = 0
	}

	dv1, _ := strconv.Atoi(string(cpfLimpo[9]))
	dv2, _ := strconv.Atoi(string(cpfLimpo[10]))

	if dv1 != digito1 || dv2 != digito2 {
		return NovoErroValidacao("CPF_INVALIDO", "CPF inválido - dígitos verificadores incorretos").
			ComDetalhes("CPF fornecido: " + cpf)
	}

	return nil
}

// ValidarCNPJ valida se um CNPJ é válido
func ValidarCNPJ(cnpj string) error {
	if cnpj == "" {
		return NovoErroValidacao("CNPJ_VAZIO", "CNPJ não pode estar vazio")
	}

	// Remove formatação
	cnpjLimpo := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(cnpj, ".", ""), "/", ""), "-", ""), " ", "")

	if len(cnpjLimpo) != 14 {
		return NovoErroValidacao("CNPJ_TAMANHO_INVALIDO", "CNPJ deve ter 14 dígitos").
			ComDetalhes("CNPJ fornecido: " + cnpj)
	}

	// Validação dos dígitos verificadores do CNPJ
	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	soma := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpjLimpo[i]))
		soma += num * pesos1[i]
	}
	digito1 := soma % 11
	if digito1 < 2 {
		digito1 = 0
	} else {
		digito1 = 11 - digito1
	}

	soma = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpjLimpo[i]))
		soma += num * pesos2[i]
	}
	digito2 := soma % 11
	if digito2 < 2 {
		digito2 = 0
	} else {
		digito2 = 11 - digito2
	}

	dv1, _ := strconv.Atoi(string(cnpjLimpo[12]))
	dv2, _ := strconv.Atoi(string(cnpjLimpo[13]))

	if dv1 != digito1 || dv2 != digito2 {
		return NovoErroValidacao("CNPJ_INVALIDO", "CNPJ inválido - dígitos verificadores incorretos").
			ComDetalhes("CNPJ fornecido: " + cnpj)
	}

	return nil
}

// ValidarEmail valida se um email é válido
func ValidarEmail(email string) error {
	if email == "" {
		return NovoErroValidacao("EMAIL_VAZIO", "Email não pode estar vazio")
	}

	if len(email) > 254 {
		return NovoErroValidacao("EMAIL_MUITO_LONGO", "Email muito longo (máximo 254 caracteres)").
			ComDetalhes("Tamanho: " + strconv.Itoa(len(email)))
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return NovoErroValidacao("EMAIL_FORMATO_INVALIDO", "Formato de email inválido").
			ComDetalhes("Email fornecido: " + email).ComCausa(err)
	}

	return nil
}

// ValidarTelefone valida se um telefone brasileiro é válido
func ValidarTelefone(telefone string) error {
	if telefone == "" {
		return NovoErroValidacao("TELEFONE_VAZIO", "Telefone não pode estar vazio")
	}

	if !regexTelefone.MatchString(telefone) {
		return NovoErroValidacao("TELEFONE_FORMATO_INVALIDO", "Formato de telefone inválido").
			ComDetalhes("Telefone fornecido: " + telefone + " - Formato esperado: (11) 99999-9999 ou 11999999999")
	}

	// Remove formatação para validação adicional
	telefoneLimpo := regexp.MustCompile(`\D`).ReplaceAllString(telefone, "")

	if len(telefoneLimpo) < 10 || len(telefoneLimpo) > 11 {
		return NovoErroValidacao("TELEFONE_TAMANHO_INVALIDO", "Telefone deve ter 10 ou 11 dígitos").
			ComDetalhes("Telefone fornecido: " + telefone)
	}

	return nil
}

// ValidarCEP valida se um CEP é válido
func ValidarCEP(cep string) error {
	if cep == "" {
		return NovoErroValidacao("CEP_VAZIO", "CEP não pode estar vazio")
	}

	if !regexCEP.MatchString(cep) {
		return NovoErroValidacao("CEP_FORMATO_INVALIDO", "Formato de CEP inválido").
			ComDetalhes("CEP fornecido: " + cep + " - Formato esperado: 12345-678 ou 12345678")
	}

	return nil
}

// ValidarSenha valida se uma senha atende aos critérios de segurança
func ValidarSenha(senha string, minCaracteres int) error {
	if senha == "" {
		return NovoErroValidacao("SENHA_VAZIA", "Senha não pode estar vazia")
	}

	if minCaracteres == 0 {
		minCaracteres = 8 // padrão
	}

	if utf8.RuneCountInString(senha) < minCaracteres {
		return NovoErroValidacao("SENHA_MUITO_CURTA", "Senha deve ter pelo menos "+strconv.Itoa(minCaracteres)+" caracteres").
			ComDetalhes("Tamanho atual: " + strconv.Itoa(utf8.RuneCountInString(senha)))
	}

	temMaiuscula := regexp.MustCompile(`[A-Z]`).MatchString(senha)
	temMinuscula := regexp.MustCompile(`[a-z]`).MatchString(senha)
	temNumero := regexp.MustCompile(`[0-9]`).MatchString(senha)
	temEspecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(senha)

	if !temMaiuscula {
		return NovoErroValidacao("SENHA_SEM_MAIUSCULA", "Senha deve conter pelo menos uma letra maiúscula")
	}

	if !temMinuscula {
		return NovoErroValidacao("SENHA_SEM_MINUSCULA", "Senha deve conter pelo menos uma letra minúscula")
	}

	if !temNumero {
		return NovoErroValidacao("SENHA_SEM_NUMERO", "Senha deve conter pelo menos um número")
	}

	if !temEspecial {
		return NovoErroValidacao("SENHA_SEM_ESPECIAL", "Senha deve conter pelo menos um caractere especial")
	}

	return nil
}

// ValidarUUID valida se uma string é um UUID válido
func ValidarUUID(uuid string) error {
	if uuid == "" {
		return NovoErroValidacao("UUID_VAZIO", "UUID não pode estar vazio")
	}

	if !regexUUID.MatchString(strings.ToLower(uuid)) {
		return NovoErroValidacao("UUID_FORMATO_INVALIDO", "Formato de UUID inválido").
			ComDetalhes("UUID fornecido: " + uuid + " - Formato esperado: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	}

	return nil
}

// ValidarID valida se um ID é válido (maior que zero)
func ValidarID(id int, nomeRecurso string) error {
	if id <= 0 {
		return NovoErroValidacao("ID_INVALIDO", "ID deve ser maior que zero").
			ComDetalhes("Recurso: " + nomeRecurso + ", ID fornecido: " + strconv.Itoa(id))
	}
	return nil
}

// ValidarStringObrigatoria valida se uma string obrigatória não está vazia
func ValidarStringObrigatoria(valor, nomeCampo string) error {
	valor = strings.TrimSpace(valor)
	if valor == "" {
		return NovoErroValidacao("CAMPO_OBRIGATORIO", "Campo obrigatório não informado").
			ComDetalhes("Campo: " + nomeCampo)
	}
	return nil
}

// ValidarTamanhoString valida o tamanho de uma string
func ValidarTamanhoString(valor, nomeCampo string, min, max int) error {
	tamanho := utf8.RuneCountInString(valor)

	if tamanho < min {
		return NovoErroValidacao("CAMPO_MUITO_CURTO", "Campo muito curto").
			ComDetalhes("Campo: " + nomeCampo + ", tamanho mínimo: " + strconv.Itoa(min) + ", tamanho atual: " + strconv.Itoa(tamanho))
	}

	if max > 0 && tamanho > max {
		return NovoErroValidacao("CAMPO_MUITO_LONGO", "Campo muito longo").
			ComDetalhes("Campo: " + nomeCampo + ", tamanho máximo: " + strconv.Itoa(max) + ", tamanho atual: " + strconv.Itoa(tamanho))
	}

	return nil
}

// ValidarNumeroPositivo valida se um número é positivo
func ValidarNumeroPositivo(numero float64, nomeCampo string) error {
	if numero <= 0 {
		return NovoErroValidacao("NUMERO_DEVE_SER_POSITIVO", "Número deve ser positivo").
			ComDetalhes("Campo: " + nomeCampo + ", valor fornecido: " + strconv.FormatFloat(numero, 'f', -1, 64))
	}
	return nil
}

// ValidarFaixaNumerica valida se um número está dentro de uma faixa
func ValidarFaixaNumerica(numero float64, min, max float64, nomeCampo string) error {
	if numero < min {
		return NovoErroValidacao("NUMERO_ABAIXO_MINIMO", "Número abaixo do valor mínimo").
			ComDetalhes("Campo: " + nomeCampo + ", mínimo: " + strconv.FormatFloat(min, 'f', -1, 64) + ", valor: " + strconv.FormatFloat(numero, 'f', -1, 64))
	}

	if numero > max {
		return NovoErroValidacao("NUMERO_ACIMA_MAXIMO", "Número acima do valor máximo").
			ComDetalhes("Campo: " + nomeCampo + ", máximo: " + strconv.FormatFloat(max, 'f', -1, 64) + ", valor: " + strconv.FormatFloat(numero, 'f', -1, 64))
	}

	return nil
}
