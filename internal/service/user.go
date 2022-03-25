package service

import (
	"context"
	v1 "knsh/api/realworld/v1"
)

func (realworld *RealworldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {

	return realworld.uc.Login(ctx, req)
}

func (realworld *RealworldService) AddCommentsToAnArticle(ctx context.Context, req *v1.AddCommentsToAnArticleRequest) (*v1.SingleComment, error) {

	return &v1.SingleComment{}, nil
}
func (realworld *RealworldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.SingleArticle, error) {
	return &v1.SingleArticle{}, nil
}
func (realworld *RealworldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.DeleteArticleResponse, error) {
	return &v1.DeleteArticleResponse{}, nil
}
func (realworld *RealworldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	return &v1.DeleteCommentResponse{}, nil
}

func (realworld *RealworldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (*v1.SingleArticle, error) {
	return &v1.SingleArticle{}, nil
}
func (realworld *RealworldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleArticles, error) {
	return &v1.MultipleArticles{}, nil
}
func FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.Profile, error) {
	return &v1.Profile{}, nil
}
func (realworld *RealworldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.SingleArticle, error) {
	return &v1.SingleArticle{}, nil
}
func (realworld *RealworldService) GetCommentsToAnArticle(ctx context.Context, req *v1.GetCommentsToAnArticleRequest) (*v1.MultipleComments, error) {
	return &v1.MultipleComments{}, nil
}
func (realworld *RealworldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserInfoReply, error) {
	return &v1.UserInfoReply{}, nil
}
func (realworld *RealworldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.Profile, error) {
	return &v1.Profile{}, nil
}
func (realworld *RealworldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.Tags, error) {
	return &v1.Tags{}, nil
}
func (realworld *RealworldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleArticles, error) {
	return &v1.MultipleArticles{}, nil
}
func (realworld *RealworldService) Registration(ctx context.Context, req *v1.RegistraRequest) (*v1.UserInfoReply, error) {
	return &v1.UserInfoReply{}, nil
}
func (realworld *RealworldService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (*v1.SingleArticle, error) {
	return &v1.SingleArticle{}, nil
}
func (realworld *RealworldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (*v1.Profile, error) {
	return &v1.Profile{}, nil
}
func (realworld *RealworldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.SingleArticle, error) {
	return &v1.SingleArticle{}, nil
}
func (realworld *RealworldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserInfoReply, error) {
	return &v1.UserInfoReply{}, nil
}
