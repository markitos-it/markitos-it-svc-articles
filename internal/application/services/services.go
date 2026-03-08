package services

import (
	"context"
	"markitos-it-svc-articles/internal/domain"
)

type ArticleService struct {
	repo domain.Repository
}

func NewArticleService(repo domain.Repository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetAllArticles(ctx context.Context) ([]domain.Article, error) {
	return s.repo.GetAll(ctx)
}

func (s *ArticleService) GetArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ArticleService) CreateArticle(ctx context.Context, doc *domain.Article) error {
	return s.repo.Create(ctx, doc)
}

func (s *ArticleService) UpdateArticle(ctx context.Context, doc *domain.Article) error {
	return s.repo.Update(ctx, doc)
}

func (s *ArticleService) DeleteArticle(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
