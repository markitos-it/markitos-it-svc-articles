package postgres

import (
	"context"
	"database/sql"
	"markitos-it-svc-articles/internal/domain"
	"reflect"
	"testing"
	"time"
)

func helperClosedDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("postgres", "host=127.0.0.1 port=1 user=test dbname=test sslmode=disable connect_timeout=1")
	if err != nil {
		t.Fatalf("failed to create db handle: %v", err)
	}
	if err := db.Close(); err != nil {
		t.Fatalf("failed to close db handle: %v", err)
	}

	return db
}

func helperRandomArticle(t *testing.T) *domain.Article {
	t.Helper()

	prefix := domain.HelperRandomAlphaPrefix(t, 8)
	return &domain.Article{
		ID:          prefix + "-article-id",
		Title:       prefix + "-article-title",
		Description: prefix + "-article-description",
		Category:    prefix + "-article-category",
		Tags:        []string{prefix + "-go", prefix + "-grpc", prefix + "-postgres"},
		UpdatedAt:   time.Date(2026, 3, 6, 12, 0, 0, 0, time.UTC),
		ContentB64:  prefix + "-Y29udGVudA==",
		CoverImage:  "https://example.com/" + prefix + "/cover.png",
	}
}

func TestNewArticleRepository(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		args   args
		wantDB *sql.DB
	}{
		{
			name:   prefix + "-build-repository-with-same-db",
			args:   args{db: db},
			wantDB: db,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewArticleRepository(tt.args.db)
			if got == nil {
				t.Fatalf("NewArticleRepository() returned nil")
			}
			if got.db != tt.wantDB {
				t.Errorf("NewArticleRepository().db = %v, want %v", got.db, tt.wantDB)
			}
		})
	}
}

func TestArticleRepository_InitSchema(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			if err := r.InitSchema(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.InitSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleRepository_SeedData(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			if err := r.SeedData(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.SeedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleRepository_GetAll(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Article
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			got, err := r.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ArticleRepository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleRepository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleRepository_GetByID(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Article
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: prefix + "-missing-id"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			got, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ArticleRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleRepository_Create(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)
	randomDoc := helperRandomArticle(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		doc *domain.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), doc: randomDoc},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			if err := r.Create(tt.args.ctx, tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleRepository_Update(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)
	randomDoc := helperRandomArticle(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		doc *domain.Article
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), doc: randomDoc},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			if err := r.Update(tt.args.ctx, tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleRepository_Delete(t *testing.T) {
	prefix := domain.HelperRandomAlphaPrefix(t, 6)
	db := helperClosedDB(t)

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    prefix + "-returns-error-on-closed-db",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: prefix + "-to-delete"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepository{db: tt.fields.db}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

