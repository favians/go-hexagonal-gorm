package item

import (
	"context"
	"go-hexagonal/business"
	"go-hexagonal/business/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDBRepository The implementation of user.Repository object
type MongoDBRepository struct {
	collection *mongo.Collection
}

type collection struct {
	ID         string    `bson:"_id"`
	Name       string    `bson:"name"`
	Username   string    `bson:"username"`
	Password   string    `bson:"password"`
	CreatedAt  time.Time `bson:"created_at"`
	CreatedBy  string    `bson:"created_by"`
	ModifiedAt time.Time `bson:"modified_at"`
	ModifiedBy string    `bson:"modified_by"`
	Version    int       `bson:"version"`
}

func newCollection(user user.User) *collection {

	return &collection{
		user.ID,
		user.Name,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
		user.ModifiedAt,
		user.ModifiedBy,
		user.Version,
	}

}

func (col *collection) ToUser() user.User {
	var user user.User

	user.ID = col.ID
	user.Name = col.Name
	user.Username = col.Username
	user.Password = col.Password
	user.CreatedAt = col.CreatedAt
	user.CreatedBy = col.CreatedBy
	user.ModifiedAt = col.ModifiedAt
	user.ModifiedBy = col.ModifiedBy
	user.Version = col.Version

	return user
}

//NewMongoDBRepository Generate mongo DB user repository
func NewMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("user"),
	}
}

//FindUserByID If data not found will return nil without error
func (repo *MongoDBRepository) FindUserByID(id string) (*user.User, error) {

	var col collection

	filter := bson.M{
		"_id": id,
	}

	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&col); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, business.ErrNotFound
		}

		return nil, err
	}

	user := col.ToUser()

	return &user, nil
}

//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
func (repo *MongoDBRepository) FindAllUserWithPagination(skip int, rowPerPage int) ([]user.User, error) {

	var users []user.User

	rowPageWithAddtion := rowPerPage + 1

	option := options.Find()
	option.SetSort(bson.D{{Key: "_id", Value: 1}})
	option.SetSkip(int64(skip))
	option.SetLimit(int64(rowPageWithAddtion))

	cursor, err := repo.collection.Find(context.Background(), bson.M{}, option)
	if err != nil {
		return users, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var col collection

		err := cursor.Decode(&col)
		if err != nil {
			return users, err
		}

		user := col.ToUser()

		users = append(users, user)
	}

	return users, nil
}

//InsertUser Insert new User into storage
func (repo *MongoDBRepository) InsertUser(user user.User) error {

	col := newCollection(user)

	_, err := repo.collection.InsertOne(context.Background(), col)
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *MongoDBRepository) UpdateUser(user user.User, currentVersion int) error {

	col := newCollection(user)

	filter := bson.M{
		"_id":     col.ID,
		"version": currentVersion,
	}

	updated := bson.M{
		"$set": col,
	}

	_, err := repo.collection.UpdateOne(context.TODO(), filter, updated)
	if err != nil {
		return err
	}

	return nil
}
