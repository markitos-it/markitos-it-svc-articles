package grpc

import (
	"context"
	"errors"
	"markitos-it-svc-articles/internal/application/services"
	"markitos-it-svc-articles/internal/domain"
	pb "markitos-it-svc-articles/proto"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type stubRepo struct {
	docs []domain.Article
	doc  *domain.Article
	err  error
}

func (r *stubRepo) GetAll(ctx context.Context) ([]domain.Article, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.docs, nil
}

func (r *stubRepo) GetByID(ctx context.Context, id string) (*domain.Article, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.doc != nil {
		return r.doc, nil
	}
	return &domain.Article{ID: id, UpdatedAt: time.Unix(0, 0).UTC()}, nil
}

func (r *stubRepo) Create(ctx context.Context, doc *domain.Article) error { return nil }
func (r *stubRepo) Update(ctx context.Context, doc *domain.Article) error { return nil }
func (r *stubRepo) Delete(ctx context.Context, id string) error          { return nil }

func TestNewArticleServer(t *testing.T) {
	svc := services.NewArticleService(&stubRepo{})
	got := NewArticleServer(svc)

	if got == nil {
		t.Fatal("expected non-nil server")
	}
	if got.service != svc {
		t.Fatal("expected same service instance")
	}
}

func TestArticleServer_GetAllArticles_Success(t *testing.T) {
	now := time.Date(2026, 3, 6, 12, 0, 0, 0, time.UTC)
	repo := &stubRepo{
		docs: []domain.Article{
			{
				ID:          "id-1",
				Title:       "title-1",
				Description: "desc-1",
				Category:    "cat-1",
				Tags:        []string{"a", "b"},
				UpdatedAt:   now,
				ContentB64:  "Y29udGVudA==",
				CoverImage:  "https://example.com/cover.png",
			},
		},
	}
	s := NewArticleServer(services.NewArticleService(repo))

	got, err := s.GetAllArticles(context.Background(), &pb.GetAllArticlesRequest{})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got == nil || len(got.Articles) != 1 {
		t.Fatalf("unexpected response: %+v", got)
	}
	if got.Articles[0].Id != "id-1" {
		t.Fatalf("expected id-1, got %q", got.Articles[0].Id)
	}
}

func TestArticleServer_GetAllArticles_Error(t *testing.T) {
	s := NewArticleServer(services.NewArticleService(&stubRepo{err: errors.New("db down")}))

	got, err := s.GetAllArticles(context.Background(), &pb.GetAllArticlesRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != nil {
		t.Fatalf("expected nil response, got %+v", got)
	}
	if status.Code(err) != codes.Internal {
		t.Fatalf("expected Internal, got %v", status.Code(err))
	}
}

func TestArticleServer_GetArticleById_Success(t *testing.T) {
	now := time.Date(2026, 3, 6, 12, 0, 0, 0, time.UTC)
	repo := &stubRepo{
		doc: &domain.Article{
			ID:          "id-42",
			Title:       "title-42",
			Description: "desc-42",
			Category:    "cat-42",
			Tags:        []string{"x"},
			UpdatedAt:   now,
			ContentB64:  "YQ==",
			CoverImage:  "https://example.com/42.png",
		},
	}
	s := NewArticleServer(services.NewArticleService(repo))

	got, err := s.GetArticleById(context.Background(), &pb.GetArticleByIdRequest{Id: "id-42"})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got == nil || got.Article == nil {
		t.Fatalf("unexpected response: %+v", got)
	}
	if got.Article.Id != "id-42" {
		t.Fatalf("expected id-42, got %q", got.Article.Id)
	}
}

func TestArticleServer_GetArticleById_Error(t *testing.T) {
	s := NewArticleServer(services.NewArticleService(&stubRepo{err: errors.New("not found")}))

	got, err := s.GetArticleById(context.Background(), &pb.GetArticleByIdRequest{Id: "missing"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != nil {
		t.Fatalf("expected nil response, got %+v", got)
	}
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", status.Code(err))
	}
}

// TestArticleServer_GetArticleById_AllFieldsMapped verifies that every field of the
// domain.Article is correctly mapped to the corresponding proto field.
func TestArticleServer_GetArticleById_AllFieldsMapped(t *testing.T) {
	now := time.Date(2026, 3, 6, 12, 0, 0, 0, time.UTC)
	doc := &domain.Article{
		ID:          "id-fields",
		Title:       "Test Title",
		Description: "Test Description",
		Category:    "TestCat",
		Tags:        []string{"tag1", "tag2"},
		UpdatedAt:   now,
		ContentB64:  "Y29udGVudA==",
		CoverImage:  "https://example.com/cover.png",
	}
	s := NewArticleServer(services.NewArticleService(&stubRepo{doc: doc}))

	resp, err := s.GetArticleById(context.Background(), &pb.GetArticleByIdRequest{Id: doc.ID})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	g := resp.Article

	if g.Id != doc.ID {
		t.Errorf("Id: want %q, got %q", doc.ID, g.Id)
	}
	if g.Title != doc.Title {
		t.Errorf("Title: want %q, got %q", doc.Title, g.Title)
	}
	if g.Description != doc.Description {
		t.Errorf("Description: want %q, got %q", doc.Description, g.Description)
	}
	if g.Category != doc.Category {
		t.Errorf("Category: want %q, got %q", doc.Category, g.Category)
	}
	if len(g.Tags) != len(doc.Tags) {
		t.Fatalf("Tags length: want %d, got %d", len(doc.Tags), len(g.Tags))
	}
	for i, tag := range doc.Tags {
		if g.Tags[i] != tag {
			t.Errorf("Tags[%d]: want %q, got %q", i, tag, g.Tags[i])
		}
	}
	if !g.UpdatedAt.AsTime().Equal(doc.UpdatedAt) {
		t.Errorf("UpdatedAt: want %v, got %v", doc.UpdatedAt, g.UpdatedAt.AsTime())
	}
	if g.ContentB64 != doc.ContentB64 {
		t.Errorf("ContentB64: want %q, got %q", doc.ContentB64, g.ContentB64)
	}
	if g.CoverImage != doc.CoverImage {
		t.Errorf("CoverImage: want %q, got %q", doc.CoverImage, g.CoverImage)
	}
}
