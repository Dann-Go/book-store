package mongodb

import (
	"context"
	"github.com/Dann-Go/book-store/internal/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	Client *mongo.Client
}

func (m mongoRepository) GetByTitle(title string) ([]domain.Book, error) {
	return nil, nil
}

func (m mongoRepository) Add(book *domain.Book) error {
	collection := m.Client.Database("book-store").Collection("books")

	_, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Error(err)
		return err
	}
	return err
}

func (m mongoRepository) GetAll() ([]domain.Book, error) {
	collection := m.Client.Database("book-store").Collection("books")

	res := make([]domain.Book, 0)

	cur, err := collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		book := domain.Book{}
		err := cur.Decode(&book)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		res = append(res, book)
	}

	if err := cur.Err(); err != nil {
		log.Error(err)
		return nil, err
	}
	cur.Close(context.TODO())

	return res, err

}

func (m mongoRepository) GetById(id int) (*domain.Book, error) {
	collection := m.Client.Database("book-store").Collection("books")
	res := &domain.Book{}

	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return res, err
}

func (m mongoRepository) Delete(id int) error {
	collection := m.Client.Database("book-store").Collection("books")

	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return err
}

func (m mongoRepository) Update(book *domain.Book, id int) error {
	collection := m.Client.Database("book-store").Collection("books")
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"id", book.ID}, {"authors", book.Authors},
		{"title", book.Title}, {"year", book.Year}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return err
}

func NewMongoRepository(Client *mongo.Client) domain.BookRepository {
	return &mongoRepository{Client}
}
