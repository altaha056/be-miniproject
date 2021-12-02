package business

import "antonio/features/third-party-news"

type ApiService struct {
	newsRepository news.Data
}

func NewApiService(nr news.Data) news.Business {
	return &ApiService{nr}
}

func (as *ApiService) GetNews(keyword string) ([]news.NewsCore, error) {
	data, err := as.newsRepository.GetData(keyword)
	if err != nil {
		return nil, err
	}

	return data, nil
}
