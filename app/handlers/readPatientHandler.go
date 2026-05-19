package handlers

import (
	"fmt"
	"net/http"
)

// ReadPatientHandler é a rota de teste/leitura para listar os pacientes da clínica
func ReadPatientHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se a rota acessada está correta
	if request.URL.Path != "/api/read-patients" {
		http.Error(response, "Página não encontrada", http.StatusNotFound)
		return
	}

	// Verifica se o método é GET (usado para buscar/ler dados)
	if request.Method != http.MethodGet {
		http.Error(response, "Método não suportado, use GET", http.StatusMethodNotAllowed)
		return
	}

	// Como ainda não criamos uma tela HTML complexa para listar os pacientes,
	// podemos retornar uma mensagem simples ou um texto formatado por enquanto:
	fmt.Fprintf(response, "🏥 Servidor de Triagem Ativo! Aqui será exibida a lista de pacientes cadastrados.")
}
