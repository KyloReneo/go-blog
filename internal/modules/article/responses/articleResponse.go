package responses

import (
	"fmt"

	articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	userResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"

)

// Response structure for the article responses
type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      userResponse.User
}

// Structure for return bunch of Articles
type Articles struct {
	Data []Articles
}

// Recives the model object and transforms it to response object
func ToArticle(article articleModel.Article) Article {
	return Article{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d",
			article.CreatedAt.Year(),
			article.CreatedAt.Month(),
			article.CreatedAt.Day()),
		Image: "/assets/img/demopic/10.jpg",
		User:  userResponse.ToUser(article.User),
	}
}
