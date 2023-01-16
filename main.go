package main

import (
	"net/http"
	"os"

	"lib.dev/schemacafe"
	"lib.dev/web"
)

func main() {
	port := os.Getenv("PORT")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	registrationKey := ""
	schemacafeOrg := schemacafe.Org{
		Members: map[string]bool{
			"admin": true,
		},
		Schemas: map[string]*web.Schema{
			"app": {
				Fields: []web.Field{},
			},
		},
	}
	app := schemacafe.App{
		Auth: web.NewAuthDB(adminPassword, registrationKey),
		Orgs: map[string]*schemacafe.Org{
			"schema-cafe": &schemacafeOrg,
		},
	}
	err := http.ListenAndServe(":"+port, &app)
	if err != nil {
		panic(err)
	}
}
