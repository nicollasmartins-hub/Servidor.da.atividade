package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
)

func CreatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	fullName := request.FormValue("nome")
	cpf := request.FormValue("cpf")
	birthDate := request.FormValue("data_nascimento")
	phone := request.FormValue("telefone")    // telefone vai para phone
	symptoms := request.FormValue("sintomas") // novo campo sintomas

	err := utils.CreatePatientDB(fullName, cpf, birthDate, phone, symptoms)
	if err != nil {
		http.Error(response, "Erro ao registrar paciente no banco de dados", http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/", http.StatusSeeOther)
}
