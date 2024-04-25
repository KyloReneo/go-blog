package services

import (
	ArticleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	ArticleRepository "github.com/kyloReneo/go-blog/internal/modules/article/repositories"
)

type ArticleSercive struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleSercive {
	return &ArticleSercive{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleSercive) GetFeaturedArticles() []ArticleModel.Article {
	return articleService.articleRepository.List(4)
}

func (articleService *ArticleSercive) GetStoriesArticles() []ArticleModel.Article {
	return articleService.articleRepository.List(6)
}
