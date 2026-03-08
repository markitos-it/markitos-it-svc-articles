package grpc

import (
	"context"
	"log"

	"markitos-it-svc-articles/internal/application/services"
	"markitos-it-svc-articles/internal/domain"
	pb "markitos-it-svc-articles/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleServer struct {
	pb.UnimplementedArticleServiceServer
	service *services.ArticleService
}

func NewArticleServer(service *services.ArticleService) *ArticleServer {
	return &ArticleServer{
		service: service,
	}
}

func (s *ArticleServer) GetAllArticles(ctx context.Context, req *pb.GetAllArticlesRequest) (*pb.GetAllArticlesResponse, error) {
	log.Println("GetAllArticles called")

	docs, err := s.service.GetAllArticles(ctx)
	if err != nil {
		log.Printf("Error getting all articles: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get articles: %v", err)
	}

	pbDocs := make([]*pb.Article, 0, len(docs))
	for _, doc := range docs {
		pbDocs = append(pbDocs, articleToProto(&doc))
	}

	return &pb.GetAllArticlesResponse{
		Articles: pbDocs,
	}, nil
}

func (s *ArticleServer) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (*pb.GetArticleByIdResponse, error) {
	log.Printf("GetArticleById called with id: %s", req.Id)

	doc, err := s.service.GetArticleByID(ctx, req.Id)
	if err != nil {
		log.Printf("Error getting article by id %s: %v", req.Id, err)
		return nil, status.Errorf(codes.NotFound, "article not found: %v", err)
	}

	return &pb.GetArticleByIdResponse{
		Article: articleToProto(doc),
	}, nil
}

func articleToProto(doc *domain.Article) *pb.Article {
	return &pb.Article{
		Id:          doc.ID,
		Title:       doc.Title,
		Description: doc.Description,
		Category:    doc.Category,
		Tags:        doc.Tags,
		UpdatedAt:   timestamppb.New(doc.UpdatedAt),
		ContentB64:  doc.ContentB64,
		CoverImage:  doc.CoverImage,
	}
}
