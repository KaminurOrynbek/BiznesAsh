package grpc

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/KaminurOrynbek/BiznesAsh/auto-proto/content"
	"github.com/KaminurOrynbek/BiznesAsh/internal/delivery/mapper"
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
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
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
	return &pb.PostResponse{Post: mapper.ConvertPostToPB(post)}, nil
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
	return &pb.PostResponse{Post: mapper.ConvertPostToPB(post)}, nil
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
	return &pb.PostResponse{Post: mapper.ConvertPostToPB(post)}, nil
}

func (h *ContentHandler) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	offset := (int(req.Page) - 1) * int(req.Limit)
	posts, err := h.postUsecase.ListPosts(ctx, offset, int(req.Limit))
	if err != nil {
		return nil, err
	}
	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, mapper.ConvertPostToPB(p))
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (h *ContentHandler) SearchPosts(ctx context.Context, req *pb.SearchPostsRequest) (*pb.ListPostsResponse, error) {
	offset := (int(req.Page) - 1) * int(req.Limit)
	posts, err := h.postUsecase.SearchPosts(ctx, req.Query, offset, int(req.Limit))
	if err != nil {
		return nil, err
	}
	var pbPosts []*pb.Post
	for _, p := range posts {
		pbPosts = append(pbPosts, mapper.ConvertPostToPB(p))
	}
	return &pb.ListPostsResponse{Posts: pbPosts}, nil
}

func (h *ContentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
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
	return &pb.CommentResponse{Comment: mapper.ConvertCommentToPB(comment)}, nil
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
	return &pb.CommentResponse{Comment: mapper.ConvertCommentToPB(comment)}, nil
}

func (h *ContentHandler) DeleteComment(ctx context.Context, req *pb.CommentIdRequest) (*pb.DeleteResponse, error) {
	if err := h.commentUsecase.DeleteComment(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (h *ContentHandler) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	comments, err := h.commentUsecase.ListCommentsByPostID(ctx, req.GetPostId())
	if err != nil {
		return nil, err
	}

	var pbComments []*pb.Comment
	for _, c := range comments {
		pbComments = append(pbComments, &pb.Comment{
			Id:        c.ID,
			PostId:    c.PostID,
			AuthorId:  c.AuthorID,
			Content:   c.Content,
			CreatedAt: timestamppb.New(c.CreatedAt),
			UpdatedAt: timestamppb.New(c.UpdatedAt),
		})
	}

	return &pb.ListCommentsResponse{Comments: pbComments}, nil
}

func (h *ContentHandler) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	like := &entity.Like{
		PostID: req.PostId,
		UserID: req.UserId,
	}
	count, err := h.likeUsecase.LikePost(ctx, like)
	if err != nil {
		return nil, err
	}
	return &pb.LikePostResponse{LikesCount: count}, nil
}

func (h *ContentHandler) DislikePost(ctx context.Context, req *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	dislike := &entity.Like{
		PostID: req.PostId,
		UserID: req.UserId,
	}
	count, err := h.likeUsecase.DislikePost(ctx, dislike)
	if err != nil {
		return nil, err
	}
	return &pb.DislikePostResponse{DislikesCount: count}, nil
}

func (h *ContentHandler) LikeComment(ctx context.Context, req *pb.LikeCommentRequest) (*pb.LikeCommentResponse, error) {
	count, err := h.likeUsecase.LikeComment(ctx, &entity.Like{
		UserID:    req.UserId,
		CommentID: req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LikeCommentResponse{LikesCount: count}, nil
}
