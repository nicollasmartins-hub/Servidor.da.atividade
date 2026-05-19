package utils

import (
	"log"
)

// DeletePatientDB remove o paciente do banco de dados (Alta Médica)
func DeletePatientDB(cpf string) error {
	// Atualizamos a query para deletar da tabela patients usando o CPF
	query := `DELETE FROM patients WHERE cpf = $1`

	// Executamos a query
	_, err := DB.Exec(query, cpf)
	if err != nil {
		log.Printf("Erro ao remover paciente do banco de dados: %v", err)
		return err
	}

	log.Println("Paciente removido (alta) com sucesso!")
	return nil
}
