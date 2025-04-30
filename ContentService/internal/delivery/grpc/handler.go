package grpc

import (
	"context"
	"github.com/google/uuid"
	"time"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/content"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
)

type ContentHandler struct {
	pb.UnimplementedContentServiceServer
	postUsecase    _interface.PostUsecase
	commentUsecase _interface.CommentUsecase
	likeUsecase    _interface.LikeUsecase
}

func NewContentHandler(
	postUC _interface.PostUsecase,
	commentUC _interface.CommentUsecase,
	likeUC _interface.LikeUsecase,
) *ContentHandler {
	return &ContentHandler{
		postUsecase:    postUC,
		commentUsecase: commentUC,
		likeUsecase:    likeUC,
	}
}

func (h *ContentHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	post := &entity.Post{
		ID:        req.Id,
		Title:     req.Title,
		Content:   req.Content,
		Type:      enum.PostType(req.Type.String()),
		AuthorID:  req.AuthorId,
		Published: req.Published,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := h.postUsecase.CreatePost(ctx, post); err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: convertPostToPB(post)}, nil
}

func (h *ContentHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	post := &entity.Post{
		ID:        req.Id,
		Title:     req.Title,
		Content:   req.Content,
		Published: req.Published,
		UpdatedAt: time.Now(),
	}
	if err := h.postUsecase.UpdatePost(ctx, post); err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: convertPostToPB(post)}, nil
}

func (h *ContentHandler) DeletePost(ctx context.Context, req *pb.PostIdRequest) (*pb.DeleteResponse, error) {
	if err := h.postUsecase.DeletePost(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (h *ContentHandler) GetPost(ctx context.Context, req *pb.PostIdRequest) (*pb.PostResponse, error) {
	post, err := h.postUsecase.GetPost(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PostResponse{Post: convertPostToPB(post)}, nil
}

func (h *ContentHandler) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := h.postUsecase.ListPosts(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, convertPostToPB(p))
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (h *ContentHandler) SearchPosts(ctx context.Context, req *pb.SearchPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := h.postUsecase.SearchPosts(ctx, req.Query, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, convertPostToPB(p))
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (h *ContentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	comment := &entity.Comment{
		ID:        req.Id,
		PostID:    req.PostId,
		AuthorID:  req.AuthorId,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := h.commentUsecase.CreateComment(ctx, comment); err != nil {
		return nil, err
	}
	return &pb.CommentResponse{Comment: convertCommentToPB(comment)}, nil
}

func (h *ContentHandler) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	comment := &entity.Comment{
		ID:        req.Id,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}
	if err := h.commentUsecase.UpdateComment(ctx, comment); err != nil {
		return nil, err
	}
	return &pb.CommentResponse{Comment: convertCommentToPB(comment)}, nil
}

func (h *ContentHandler) DeleteComment(ctx context.Context, req *pb.CommentIdRequest) (*pb.DeleteResponse, error) {
	if err := h.commentUsecase.DeleteComment(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (h *ContentHandler) ListComments(ctx context.Context, req *pb.PostIdRequest) (*pb.ListCommentsResponse, error) {
	comments, err := h.commentUsecase.ListCommentsByPostID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	var pbComments []*pb.Comment
	for _, c := range comments {
		pbComments = append(pbComments, convertCommentToPB(c))
	}
	return &pb.ListCommentsResponse{Comments: pbComments}, nil
}

func (h *ContentHandler) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	like := &entity.Like{
		ID:        uuid.NewString(),
		PostID:    req.PostId,
		UserID:    req.UserId,
		IsLike:    true,
		CreatedAt: time.Now(),
	}
	if err := h.likeUsecase.LikePost(ctx, like); err != nil {
		return nil, err
	}
	return &pb.LikePostResponse{LikesCount: 1}, nil
}

func (h *ContentHandler) DislikePost(ctx context.Context, req *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	dislike := &entity.Like{
		ID:        uuid.NewString(),
		PostID:    req.PostId,
		UserID:    req.UserId,
		IsLike:    false,
		CreatedAt: time.Now(),
	}
	if err := h.likeUsecase.DislikePost(ctx, dislike); err != nil {
		return nil, err
	}
	return &pb.DislikePostResponse{DislikesCount: 1}, nil
}

// mappers for post and comment
func convertPostToPB(p *entity.Post) *pb.Post {
	pbComments := make([]*pb.Comment, 0, len(p.Comments))
	for _, c := range p.Comments {
		pbComments = append(pbComments, convertCommentToPB(&c))
	}
	return &pb.Post{
		Id:            p.ID,
		Title:         p.Title,
		Content:       p.Content,
		Type:          pb.PostType(pb.PostType_value[string(p.Type)]),
		AuthorId:      p.AuthorID,
		Published:     p.Published,
		LikesCount:    p.LikesCount,
		DislikesCount: p.DislikesCount,
		CreatedAt:     p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     p.UpdatedAt.Format(time.RFC3339),
		CommentsCount: p.CommentsCount,
		Comments:      pbComments,
	}
}

func convertCommentToPB(c *entity.Comment) *pb.Comment {
	return &pb.Comment{
		Id:        c.ID,
		PostId:    c.PostID,
		AuthorId:  c.AuthorID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
	}
}
