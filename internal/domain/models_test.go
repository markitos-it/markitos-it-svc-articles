package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestArticle_ShouldKeepAssignedDomainState(t *testing.T) {
	now := time.Date(2026, 2, 22, 10, 30, 0, 0, time.UTC)
	prefix := HelperRandomAlphaPrefix(t, 8)
	expectedID := prefix + "-article-001"
	expectedTitle := prefix + "-Servicio de Documentos"
	expectedDescription := prefix + "-Plantilla base"
	expectedCategory := prefix + "-backend"
	expectedContentB64 := prefix + "-Y29udGVudA=="
	expectedCoverImage := "https://example.com/" + prefix + "/cover.png"
	expectedTags := []string{prefix + "-go", prefix + "-grpc", prefix + "-postgres"}

	article := Article{
		ID:          expectedID,
		Title:       expectedTitle,
		Description: expectedDescription,
		Category:    expectedCategory,
		Tags:        expectedTags,
		UpdatedAt:   now,
		ContentB64:  expectedContentB64,
		CoverImage:  expectedCoverImage,
	}

	if article.ID != expectedID {
		t.Fatalf("expected ID %s, got %s", expectedID, article.ID)
	}
	if article.Title != expectedTitle {
		t.Fatalf("expected Title %s, got %s", expectedTitle, article.Title)
	}
	if article.Description != expectedDescription {
		t.Fatalf("expected Description %s, got %s", expectedDescription, article.Description)
	}
	if article.Category != expectedCategory {
		t.Fatalf("expected Category %s, got %s", expectedCategory, article.Category)
	}
	if !reflect.DeepEqual(article.Tags, expectedTags) {
		t.Fatalf("expected Tags %v, got %v", expectedTags, article.Tags)
	}
	if !article.UpdatedAt.Equal(now) {
		t.Fatalf("expected UpdatedAt %v, got %v", now, article.UpdatedAt)
	}
	if article.ContentB64 != expectedContentB64 {
		t.Fatalf("expected ContentB64 %s, got %s", expectedContentB64, article.ContentB64)
	}
	if article.CoverImage != expectedCoverImage {
		t.Fatalf("expected CoverImage %s, got %s", expectedCoverImage, article.CoverImage)
	}
}

func TestArticle_ShouldExposeZeroValueAsEmptyDomainState(t *testing.T) {
	var article Article

	if article.ID != "" {
		t.Fatalf("expected empty ID, got %s", article.ID)
	}
	if article.Title != "" {
		t.Fatalf("expected empty Title, got %s", article.Title)
	}
	if article.Description != "" {
		t.Fatalf("expected empty Description, got %s", article.Description)
	}
	if article.Category != "" {
		t.Fatalf("expected empty Category, got %s", article.Category)
	}
	if len(article.Tags) != 0 {
		t.Fatalf("expected empty Tags, got %v", article.Tags)
	}
	if !article.UpdatedAt.IsZero() {
		t.Fatalf("expected zero UpdatedAt, got %v", article.UpdatedAt)
	}
	if article.ContentB64 != "" {
		t.Fatalf("expected empty ContentB64, got %s", article.ContentB64)
	}
	if article.CoverImage != "" {
		t.Fatalf("expected empty CoverImage, got %s", article.CoverImage)
	}
}
