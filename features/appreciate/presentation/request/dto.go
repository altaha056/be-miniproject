package request

import (
	"antonio/features/contents"
)

type ArticleRequest struct {
	Title   string   `json:"title"`
	Image   string   `json:"image"`
	Content string   `json:"content"`
	Tools   []string `json:"tools"`
}

func toTagCoreList(requestTagsData []string) []contents.ToolCore {
	convertedData := make([]contents.ToolCore, 0, len(requestTagsData))
	for _, v := range requestTagsData {
		convertedData = append(convertedData, contents.ToolCore{Title: v})
	}
	return convertedData
}

func (requestData *ArticleRequest) ToArticleCore() contents.Core {
	return contents.Core{
		Title:   requestData.Title,
		Image:   requestData.Image,
		Content: requestData.Content,
		Tools:   toTagCoreList(requestData.Tools),
	}
}
