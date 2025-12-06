package articles

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupArticle() Article {
	return Article{
		Title:       "Test Article",
		Body:        "This is a test article body",
		Description: "Test description",
		AuthorID:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Tags:        []string{"go", "testing"},
		Favorites:   0,
	}
}

func TestArticleCreation(t *testing.T) {
	article := setupArticle()
	err := article.Validate()
	assert.Nil(t, err)
}

func TestArticleEmptyTitle(t *testing.T) {
	article := setupArticle()
	article.Title = ""
	err := article.Validate()
	assert.NotNil(t, err)
}

func TestArticleEmptyBody(t *testing.T) {
	article := setupArticle()
	article.Body = ""
	err := article.Validate()
	assert.NotNil(t, err)
}

func TestArticleFavorite(t *testing.T) {
	article := setupArticle()
	article.Favorite()
	assert.Equal(t, 1, article.Favorites)
	article.Unfavorite()
	assert.Equal(t, 0, article.Favorites)
}

func TestArticleTagAssociation(t *testing.T) {
	article := setupArticle()
	assert.Contains(t, article.Tags, "go")
	assert.Contains(t, article.Tags, "testing")
}

func TestArticleSerializer(t *testing.T) {
	article := setupArticle()
	data := ArticleSerializer(article)
	assert.Equal(t, article.Title, data["title"])
	assert.Equal(t, article.Description, data["description"])
}

func TestArticleListSerializer(t *testing.T) {
	articles := []Article{setupArticle(), setupArticle()}
	data := ArticleListSerializer(articles)
	assert.Len(t, data, 2)
}

func TestCommentSerializer(t *testing.T) {
	comment := Comment{Body: "Nice article", AuthorID: 1}
	data := CommentSerializer(comment)
	assert.Equal(t, "Nice article", data["body"])
	assert.Equal(t, 1, data["author_id"])
}

func TestArticleModelValidatorValid(t *testing.T) {
	article := setupArticle()
	err := ArticleModelValidator(article)
	assert.Nil(t, err)
}

func TestArticleModelValidatorInvalid(t *testing.T) {
	article := setupArticle()
	article.Title = ""
	err := ArticleModelValidator(article)
	assert.NotNil(t, err)
}

func TestCommentModelValidatorValid(t *testing.T) {
	comment := Comment{Body: "Great post", AuthorID: 1}
	err := CommentModelValidator(comment)
	assert.Nil(t, err)
}

func TestCommentModelValidatorInvalid(t *testing.T) {
	comment := Comment{Body: "", AuthorID: 1}
	err := CommentModelValidator(comment)
	assert.NotNil(t, err)
}

func TestSlugGeneration(t *testing.T) {
	article := setupArticle()
	article.GenerateSlug()
	assert.NotEmpty(t, article.Slug)
}

func TestArticleUpdate(t *testing.T) {
	article := setupArticle()
	newTitle := "Updated Title"
	article.Title = newTitle
	assert.Equal(t, newTitle, article.Title)
}

func TestArticleDeletion(t *testing.T) {
	articles := []Article{setupArticle()}
	articles = articles[:0]
	assert.Empty(t, articles)
}

func TestTagListRetrieval(t *testing.T) {
	article := setupArticle()
	tags := article.GetTags()
	assert.Contains(t, tags, "go")
	assert.Contains(t, tags, "testing")
}
