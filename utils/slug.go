package utils

import (
	"strings"
	"github.com/google/uuid"
)

func GenerateSlug(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

func GenerateUUID() string {
	return uuid.New().String()
}