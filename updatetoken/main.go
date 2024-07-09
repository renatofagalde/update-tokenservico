package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbHost     = os.Getenv("host")
	dbUser     = os.Getenv("user")
	dbPassword = os.Getenv("password")
	dbName     = os.Getenv("database")
)

func handleRequest(ctx context.Context) (string, error) {
	fmt.Println("cross-07")
	// Conectar ao banco de dados
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
		return "", err
	}
	defer db.Close()

	// Verificar se o select retorna null
	var tokenServico sql.NullString
	err = db.QueryRow("SELECT tokenServico FROM Usuario WHERE idUsuario = 1 AND tokenServico IS NOT NULL").
		Scan(&tokenServico)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error querying database: %s", err)
		return "", err
	}

	// Se o select retornar null, fazer o update
	if !tokenServico.Valid {
		_, err = db.Exec("UPDATE Usuario SET fotoFacebookURL = 'xpto' WHERE idUsuario = 1")
		if err != nil {
			log.Fatalf("Error updating database: %s", err)
			return "", err
		}
		return "Update successful", nil
	}

	return "No update needed", nil
}

func main() {
	lambda.Start(handleRequest)
}
