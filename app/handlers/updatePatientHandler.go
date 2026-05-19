package handlers

import (
	"fmt"
	"net/http"
	"servidorHTTP/app/utils"
	"text/template"
)

// UpdatePatientHandler é responsável por atualizar os dados do paciente
func UpdatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém o identificador principal do paciente (obrigatório para achar no banco)
	cpf := request.FormValue("cpf")

	// Obtém os valores opcionais que a recepção pode querer atualizar
	status := request.FormValue("status")
	symptoms := request.FormValue("symptoms")

	// Cria um mapa para armazenar apenas os campos que realmente vieram preenchidos
	updates := make(map[string]string)

	if status != "" {
		updates["status"] = status
	}
	if symptoms != "" {
		updates["symptoms"] = symptoms
	}

	// Atualiza os campos informados no banco de dados chamando nosso utilitário
	err := utils.UpdatePatientDB(cpf, updates)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados do paciente no banco", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Busca as informações recém-atualizadas do paciente para mostrar na tela
	patient, err := utils.GetPatientByCPF(cpf)
	if err != nil {
		http.Error(response, "Erro ao buscar informações atualizadas do paciente", http.StatusInternalServerError)
		return
	}

	// Carrega o template do prontuário (antigo profile.html)
	tmpl, err := template.ParseFiles("static/patientProfile.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template do prontuário", http.StatusInternalServerError)
		return
	}

	// Renderiza o template com os dados atualizados do paciente
	err = tmpl.Execute(response, patient)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
