package services

import (
	"errors"

	ArticleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	ArticleRepository "github.com/kyloReneo/go-blog/internal/modules/article/repositories"
	"github.com/kyloReneo/go-blog/internal/modules/article/requests/articles"
	ArticleResponse "github.com/kyloReneo/go-blog/internal/modules/article/responses"
	UserResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"
)

type ArticleSercive struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleSercive {
	return &ArticleSercive{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleSercive) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleSercive) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleSercive) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return ArticleResponse.ToArticle(article), nil

}

func (articleService *ArticleSercive) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	var response ArticleResponse.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)
	if newArticle.ID == 0 {
		return response, errors.New("somthing went wrong when creating the article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}
