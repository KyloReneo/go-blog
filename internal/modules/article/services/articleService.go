package services

import (
	"errors"

	ArticleRepository "github.com/kyloReneo/go-blog/internal/modules/article/repositories"
	ArticleResponse "github.com/kyloReneo/go-blog/internal/modules/article/responses"

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
