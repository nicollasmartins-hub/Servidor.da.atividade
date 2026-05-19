package utils

import (
	"log"
)

// Patient representa a estrutura de dados do paciente
type Patient struct {
	FullName  string
	CPF       string
	BirthDate string
	Symptoms  string
	Status    string
}

// GetPatientByCPF busca um paciente no banco de dados usando o CPF
func GetPatientByCPF(cpf string) (*Patient, error) {
	// A query agora busca na tabela patients usando o CPF
	query := `SELECT full_name, cpf, birth_date, symptoms, status FROM patients WHERE cpf = $1`

	var patient Patient

	// Executa a query e copia (Scan) os resultados para dentro da nossa estrutura Patient
	err := DB.QueryRow(query, cpf).Scan(&patient.FullName, &patient.CPF, &patient.BirthDate, &patient.Symptoms, &patient.Status)
	if err != nil {
		log.Printf("Erro ao buscar paciente no banco de dados: %v", err)
		return nil, err
	}

	return &patient, nil
}
