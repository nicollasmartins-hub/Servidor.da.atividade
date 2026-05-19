package utils

import (
	"log"
)

func CreatePatientDB(fullName, cpf, birthDate, phone, symptoms string) error {
	query := `INSERT INTO patients (full_name, cpf, birth_date, phone, symptoms) VALUES ($1, $2, $3, $4, $5)`

	_, err := DB.Exec(query, fullName, cpf, birthDate, phone, symptoms)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return err
	}

	log.Println("Paciente inserido com sucesso!")
	return nil
}
