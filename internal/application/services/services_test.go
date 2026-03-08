package services

import (
	"context"
	"errors"
	"markitos-it-svc-articles/internal/domain"
	"testing"
)

// fakeRepo is a happy-path fake: all operations succeed and return predictable data.
type fakeRepo struct{}

func (fakeRepo) GetAll(ctx context.Context) ([]domain.Article, error) { return nil, nil }
func (fakeRepo) GetByID(ctx context.Context, id string) (*domain.Article, error) {
	return &domain.Article{ID: id}, nil
}
func (fakeRepo) Create(ctx context.Context, doc *domain.Article) error { return nil }
func (fakeRepo) Update(ctx context.Context, doc *domain.Article) error { return nil }
func (fakeRepo) Delete(ctx context.Context, id string) error          { return nil }

// failingRepo always returns an error on every operation.
type failingRepo struct {
	err error
}

func (r failingRepo) GetAll(ctx context.Context) ([]domain.Article, error)            { return nil, r.err }
func (r failingRepo) GetByID(ctx context.Context, id string) (*domain.Article, error) { return nil, r.err }
func (r failingRepo) Create(ctx context.Context, doc *domain.Article) error            { return r.err }
func (r failingRepo) Update(ctx context.Context, doc *domain.Article) error            { return r.err }
func (r failingRepo) Delete(ctx context.Context, id string) error                    { return r.err }

// ---------------------------------------------------------------------------
// NewArticleService
// ---------------------------------------------------------------------------

func TestNewArticleService_WithNilRepo(t *testing.T) {
	svc := NewArticleService(nil)
	if svc == nil {
		t.Fatal("expected non-nil ArticleService")
	}
}

func TestNewArticleService_WithRepo(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	if svc == nil {
		t.Fatal("expected non-nil ArticleService")
	}
}

// ---------------------------------------------------------------------------
// GetAllArticles
// ---------------------------------------------------------------------------

func TestArticleService_GetAllArticles_ReturnsResults(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	got, err := svc.GetAllArticles(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil && len(got) != 0 {
		t.Fatalf("expected nil/empty slice, got %v", got)
	}
}

func TestArticleService_GetAllArticles_PropagatesError(t *testing.T) {
	want := errors.New("repo down")
	svc := NewArticleService(failingRepo{err: want})
	_, err := svc.GetAllArticles(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

// ---------------------------------------------------------------------------
// GetArticleByID
// ---------------------------------------------------------------------------

func TestArticleService_GetArticleByID_ReturnsCorrectID(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	got, err := svc.GetArticleByID(context.Background(), "my-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || got.ID != "my-id" {
		t.Fatalf("expected ID=my-id, got %+v", got)
	}
}

func TestArticleService_GetArticleByID_PropagatesError(t *testing.T) {
	want := errors.New("not found")
	svc := NewArticleService(failingRepo{err: want})
	_, err := svc.GetArticleByID(context.Background(), "any")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

// ---------------------------------------------------------------------------
// CreateArticle
// ---------------------------------------------------------------------------

func TestArticleService_CreateArticle_Success(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	if err := svc.CreateArticle(context.Background(), &domain.Article{ID: "new-id"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestArticleService_CreateArticle_PropagatesError(t *testing.T) {
	want := errors.New("create failed")
	svc := NewArticleService(failingRepo{err: want})
	err := svc.CreateArticle(context.Background(), &domain.Article{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

// ---------------------------------------------------------------------------
// UpdateArticle
// ---------------------------------------------------------------------------

func TestArticleService_UpdateArticle_Success(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	if err := svc.UpdateArticle(context.Background(), &domain.Article{ID: "existing-id"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestArticleService_UpdateArticle_PropagatesError(t *testing.T) {
	want := errors.New("update failed")
	svc := NewArticleService(failingRepo{err: want})
	err := svc.UpdateArticle(context.Background(), &domain.Article{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

// ---------------------------------------------------------------------------
// DeleteArticle
// ---------------------------------------------------------------------------

func TestArticleService_DeleteArticle_Success(t *testing.T) {
	svc := NewArticleService(fakeRepo{})
	if err := svc.DeleteArticle(context.Background(), "existing-id"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestArticleService_DeleteArticle_PropagatesError(t *testing.T) {
	want := errors.New("delete failed")
	svc := NewArticleService(failingRepo{err: want})
	err := svc.DeleteArticle(context.Background(), "any")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, want) {
		t.Fatalf("expected %v, got %v", want, err)
	}
}

