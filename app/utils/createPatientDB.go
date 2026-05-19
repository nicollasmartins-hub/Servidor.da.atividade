package utils

import (
	"log"
)

// CreatePatientDB insere um novo paciente no banco de dados
func CreatePatientDB(fullName, cpf, birthDate, symptoms string) error {
	// Atualizamos a query para inserir na tabela de pacientes
	query := `INSERT INTO patients (full_name, cpf, birth_date, symptoms) VALUES ($1, $2, $3, $4)`

	// Executamos a query passando os dados do paciente
	_, err := DB.Exec(query, fullName, cpf, birthDate, symptoms)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return err
	}

	log.Println("Paciente inserido com sucesso!")
	return nil
}
