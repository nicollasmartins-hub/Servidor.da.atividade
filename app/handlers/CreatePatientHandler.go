package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
	"net/http"               // Usado para lidar com requisições e respostas HTTP
	"servidorHTTP/app/utils" // Importa funções utilitárias, como inserção no banco de dados
)

// CreatePatientHandler é responsável por processar os dados enviados pelo formulário de triagem
func CreatePatientHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		// Retorna um erro caso o método não seja suportado
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores enviados pelo formulário HTML (baseado no atributo 'name' das tags <input>)
	fullName := request.FormValue("full_name")   // Nome completo do paciente
	cpf := request.FormValue("cpf")              // CPF do paciente
	birthDate := request.FormValue("birth_date") // Data de nascimento
	symptoms := request.FormValue("symptoms")    // Sintomas relatados na triagem

	// Insere os dados do paciente no banco de dados chamando o nosso utilitário
	err := utils.CreatePatientDB(fullName, cpf, birthDate, symptoms)
	if err != nil {
		// Retorna um erro caso ocorra falha ao salvar os dados no banco de dados
		http.Error(response, "Erro ao registrar paciente no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página inicial (menu principal) após o registro ser concluído com sucesso
	http.Redirect(response, request, "/", http.StatusSeeOther)
}
