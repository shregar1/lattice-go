// Package main provides the CLI Module Generator for lattice-go.
// Usage: go run scripts/create_module.go Product
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Error: Please provide a module name.")
		fmt.Println("Example: go run scripts/create_module.go Product")
		os.Exit(1)
	}

	raw := os.Args[1]
	pascal := strings.Title(raw)
	snake := strings.ToLower(raw)
	kebab := strings.ReplaceAll(snake, "_", "-")

	fmt.Printf("🚀 Scaffolding Lattice-Go Clean-Architecture Module for: %s...\n\n", pascal)

	// 1. Model
	writeFile(filepath.Join("internal", "models", snake+".go"), fmt.Sprintf(`package models

import "time"

type %s struct {
	ID        string    `+"`json:\"id\"`"+`
	URN       string    `+"`json:\"urn\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	OwnerID   string    `+"`json:\"ownerId\"`"+`
	TenantID  string    `+"`json:\"tenantId\"`"+`
	IsDeleted bool      `+"`json:\"isDeleted\"`"+`
	CreatedAt time.Time `+"`json:\"createdAt\"`"+`
	UpdatedAt time.Time `+"`json:\"updatedAt\"`"+`
}
`, pascal))

	// 2. DTO
	writeFile(filepath.Join("internal", "dto", snake+".go"), fmt.Sprintf(`package dto

type Create%sRequest struct {
	Name string `+"`json:\"name\" validate:\"required,min=1,max=255\"`"+`
}

type %sResponse struct {
	ID      string `+"`json:\"id\"`"+`
	URN     string `+"`json:\"urn\"`"+`
	Name    string `+"`json:\"name\"`"+`
	OwnerID string `+"`json:\"ownerId\"`"+`
}
`, pascal, pascal))

	// 3. Repository Interface & Implementation
	writeFile(filepath.Join("internal", "repositories", snake+"_repository.go"), fmt.Sprintf(`package repositories

import (
	"context"
	"github.com/shregar1/lattice-go/internal/models"
)

type I%sRepository interface {
	FindByID(ctx context.Context, id string) (*models.%s, error)
	Create(ctx context.Context, entity *models.%s) (*models.%s, error)
}
`, pascal, pascal, pascal, pascal))

	// 4. Service Layer
	writeFile(filepath.Join("internal", "services", snake+"_service.go"), fmt.Sprintf(`package services

import (
	"context"
	"github.com/shregar1/lattice-go/internal/dto"
	"github.com/shregar1/lattice-go/internal/repositories"
)

type I%sService interface {
	Create(ctx context.Context, input dto.Create%sRequest, ownerID string) (*dto.%sResponse, error)
}

type %sService struct {
	repo repositories.I%sRepository
}
`, pascal, pascal, pascal, pascal, pascal))

	// 5. Orchestrator
	writeFile(filepath.Join("internal", "orchestrators", snake+"_orchestrator.go"), fmt.Sprintf(`package orchestrators

import (
	"context"
	"github.com/shregar1/lattice-go/internal/dto"
	"github.com/shregar1/lattice-go/internal/services"
)

type %sOrchestrator struct {
	service services.I%sService
}

func (o *%sOrchestrator) Create(ctx context.Context, input dto.Create%sRequest, ownerID string) (*dto.%sResponse, error) {
	return o.service.Create(ctx, input, ownerID)
}
`, pascal, pascal, pascal, pascal, pascal))

	fmt.Printf("\n✨ Scaffolding complete for lattice-go module '%s'!\n", pascal)
}

func writeFile(path, content string) {
	os.MkdirAll(filepath.Dir(path), 0755)
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("  ⚠️ Skipping existing file: %s\n", path)
		return
	}
	os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0644)
	fmt.Printf("  ✓ Created: %s\n", path)
}
