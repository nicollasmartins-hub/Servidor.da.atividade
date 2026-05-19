package handlers

import (
	"net/http"
	"servidorHTTP/app/utils" // Se o nome do seu módulo for diferente, lembre-se de ajustar aqui
)

func DeletePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// 1. Em vez de email e senha, vamos pegar o CPF que virá do HTML
	cpf := request.FormValue("cpf")

	// 2. Apagamos a parte de criptografia (utils.Encrypt) e validação (utils.ValidateUser)

	// 3. Chama a função do banco de dados para deletar (Dar Alta) pelo CPF
	err := utils.DeletePatientDB(cpf)
	if err != nil {
		http.Error(response, "Erro ao remover o paciente do sistema", http.StatusInternalServerError)
		return
	}

	// 4. Redireciona de volta para a página inicial após o sucesso
	http.Redirect(response, request, "/", http.StatusSeeOther)
}
