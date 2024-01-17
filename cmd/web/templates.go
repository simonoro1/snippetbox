package main

import (
	"snippetbox.simonoro1/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
