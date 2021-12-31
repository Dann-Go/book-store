package elastic_search

import (
	"context"
	"encoding/json"
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"strconv"
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
	books := make([]domain.Book, 0)
	query := elastic.MatchAllQuery{}
	searchResult, err := e.Client.Search().Index("books").Query(&query).Do(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res := &domain.Book{}
	for _, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &res)
		if err != nil {
			return nil, err
		}
		books = append(books, *res)
	}
	return books, err
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
	bq := elastic.NewBoolQuery()
	bq.Must(elastic.NewTermQuery("id", id))

	_, err := elastic.NewDeleteByQueryService(e.Client).Index("books").Query(bq).Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	e.Client.Flush().Index("books").Do(context.Background())
	return err
}

// Ask about it ???
func (e elasticRepository) Update(book *domain.Book, id int) error {

	_, err := e.Client.Update().Index("books").Id(strconv.Itoa(id)).Doc(book).Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	e.Client.Flush().Index("books").Do(context.Background())
	return err
}
