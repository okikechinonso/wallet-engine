package mongo_db

//import (
//	"context"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"wallet-engine/domain/entity"
//	db "wallet-engine/infrastructure/persistence/mysql"
//)
//
//
//const limit = 10
//
//func (db.Database) GetMovies(page int) ([]entity.Movie, error) {
//	l := int64(limit)
//	skip := int64((page * limit) - limit)
//	option := options.FindOptions{Limit: &l, Skip: &skip}
//	coll := d.Client.Database("sample_mflix").Collection("movies")
//
//	cursor, err := coll.Find(context.TODO(), "", &option)
//	if err != nil {
//		return nil, err
//	}
//	moveis := []entity.Movie{}
//	if err = cursor.All(context.TODO(), &moveis); err != nil {
//		return nil, err
//	}
//	return moveis, nil
//}
