package request

import (
	"antonio/features/articles"
)

type ArticleRequest struct {
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func toTagCoreList(requestTagsData []string) []articles.TagCore {
	convertedData := make([]articles.TagCore, 0, len(requestTagsData))
	for _, v := range requestTagsData {
		convertedData = append(convertedData, articles.TagCore{Title: v})
	}
	return convertedData
}

func (requestData *ArticleRequest) ToArticleCore() articles.Core {
	return articles.Core{
		Title:   requestData.Title,
		Image:   requestData.Image,
		Content: requestData.Content,
		Tags:    toTagCoreList(requestData.Tags),
	}
}
