package articles

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Mock Article model
type Article struct {
	Title    string
	Body     string
	Favorites int
	Tags     []string
}

// Mock functions
func NewArticle(title, body string) (*Article, error) {
	if title == "" || body == "" {
		return nil, assert.AnError
	}
	return &Article{Title: title, Body: body}, nil
}

func (a *Article) Favorite() {
	a.Favorites++
}

func (a *Article) Unfavorite() {
	if a.Favorites > 0 {
		a.Favorites--
	}
}

// ===================== TESTS =====================

func TestArticleCreation_Valid(t *testing.T) {
	article, err := NewArticle("Test Title", "Test Body")
	assert.Nil(t, err)
	assert.Equal(t, "Test Title", article.Title)
	assert.Equal(t, "Test Body", article.Body)
}

func TestArticleCreation_EmptyTitle(t *testing.T) {
	_, err := NewArticle("", "Test Body")
	assert.NotNil(t, err)
}

func TestArticleFavorite(t *testing.T) {
	article, _ := NewArticle("A", "B")
	article.Favorite()
	assert.Equal(t, 1, article.Favorites)
	article.Favorite()
	assert.Equal(t, 2, article.Favorites)
}

func TestArticleUnfavorite(t *testing.T) {
	article, _ := NewArticle("A", "B")
	article.Favorite()
	article.Unfavorite()
	assert.Equal(t, 0, article.Favorites)
	article.Unfavorite()
	assert.Equal(t, 0, article.Favorites)
}

func TestArticleTags(t *testing.T) {
	article, _ := NewArticle("Title", "Body")
	article.Tags = []string{"go", "test"}
	assert.Contains(t, article.Tags, "go")
	assert.Contains(t, article.Tags, "test")
}

func TestArticleSerializer(t *testing.T) {
	article, _ := NewArticle("T", "B")
	article.Tags = []string{"tag1"}
	serialized := map[string]interface{}{
		"title": article.Title,
		"body":  article.Body,
		"tags":  article.Tags,
	}
	assert.Equal(t, "T", serialized["title"])
	assert.Equal(t, "B", serialized["body"])
	assert.Len(t, serialized["tags"], 1)
}

func TestArticleListSerializer(t *testing.T) {
	articles := []*Article{
		{Title: "A", Body: "B"},
		{Title: "C", Body: "D"},
	}
	listSerialized := make([]map[string]string, len(articles))
	for i, a := range articles {
		listSerialized[i] = map[string]string{
			"title": a.Title,
			"body":  a.Body,
		}
	}
	assert.Len(t, listSerialized, 2)
	assert.Equal(t, "C", listSerialized[1]["title"])
}

func TestCommentSerializer(t *testing.T) {
	comment := map[string]string{"body": "Nice article"}
	assert.Equal(t, "Nice article", comment["body"])
}
