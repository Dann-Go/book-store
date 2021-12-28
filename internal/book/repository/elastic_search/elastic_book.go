package elastic_search

import (
	"context"
	"encoding/json"
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/olivere/elastic/v7"
)

func NewElasticRepository(Client *elastic.Client) domain.BookRepository {
	return &elasticRepository{Client}
}

type elasticRepository struct {
	Client *elastic.Client
}

func (e elasticRepository) Add(book *domain.Book) error {
	client := e.Client
	_, err := client.Index().Index("books").
		BodyJson(book).
		Do(context.Background())
	if err != nil {
		return err
	}
	return err
}

func (e elasticRepository) GetAll() ([]domain.Book, error) {
	return nil, nil
}

func (e elasticRepository) GetById(id int) (*domain.Book, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("id", id))

	res := &domain.Book{}

	searchService := e.Client.Search().Index("books").SearchSource(searchSource)

	searchResult, err := searchService.Do(context.Background())
	if err != nil {
		return nil, err
	}
	for _, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &res)
		if err != nil {
			return nil, err
		}

	}

	return res, err

}

func (e elasticRepository) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func (e elasticRepository) Update(book *domain.Book, id int) error {
	//TODO implement me
	panic("implement me")
}
