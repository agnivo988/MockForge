package main

import (
	"log"

	"github.com/agnivo988/MockForge/backend/internal/api"
	"github.com/agnivo988/MockForge/backend/internal/mock"
	"github.com/agnivo988/MockForge/backend/internal/parser"
	"github.com/agnivo988/MockForge/backend/internal/recorder"
	"github.com/agnivo988/MockForge/backend/internal/store"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	spec, err := parser.LoadSpec("spec.yaml")
	if err != nil {
		log.Fatal(err)

	}
	store.SetSpec(spec)

	mock.RegisterMockRoutes(r, spec)

	r.Use(cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:3000"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	AllowHeaders:     []string{"Origin", "Content-Type"},
	AllowCredentials: true,
}))

	r.Use(recorder.Middleware())
	r.POST("/api/spec/upload", api.UploadSpec)

	log.Println("MockForage running on :8080")
	r.Run(":8080")

}

func loadSpecs() {
	files := map[string]string{
		"dev":     "specs/dev.yaml",
		"staging": "specs/staging.yaml",
		"prod":    "specs/prod.yaml",
	}

	for env, path := range files {
		spec, err := parser.LoadSpec(path)
		if err != nil {
			log.Fatalf("[%s] %v", env, err)
		}
		store.Manager.Set(env, spec)
	}

	store.Manager.SetActive("dev")
}
