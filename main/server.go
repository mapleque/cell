package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mapleque/cell/server"
)

func main() {
	config := server.NewConfig()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	db, err := sql.Open(
		"mysql",
		config.Get("DB_DSN"),
	)
	if err != nil {
		logger.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(config.Int("DB_MAX_IDLE"))
	db.SetMaxOpenConns(config.Int("DB_MAX_OPEN"))

	oidc := server.NewOidc(
		db,
		logger,
	)
	oidc.AddKeyPair(
		"key1",
		config.Get("OIDC_PUBLIC_KEY_FILE"),
		config.Get("OIDC_PRIVATE_KEY_FILE"),
	)

	mail := server.NewMail(
		logger,
		config.Get("MAIL_USERNAME"),
		config.Get("MAIL_PASSWORD"),
		config.Get("MAIL_HOST"),
		config.Get("MAIL_ADDRESS"),
		config.Get("MAIL_FROM"),
		config.Get("MAIL_TEMPLATE_PATTERN"),
	)

	auth := server.NewAuth(
		db,
		logger,
		mail,
	)

	kerberos := server.NewKerberos(
		db,
		logger,
		config.Get("KERBEROS_TGS_SECRET_KEY"),
		config.Get("KERBEROS_APP_SECRET_KEY"),
	)

	s := server.NewServer(
		config.Get("HTTP_LISTEN"),
		config.Get("HTTP_STATIC_PATH"),
		db,
		logger,
		oidc,
		auth,
		kerberos,
	)

	s.Run(
		config.Get("HTTPS_KEY_FILE"),
		config.Get("HTTPS_CERT_FILE"),
	)
}
