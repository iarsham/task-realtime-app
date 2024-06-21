package repository

import (
	"context"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/domain"
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/iarsham/task-realtime-app/user-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type usersRepositoryImpl struct {
	db  *mongo.Database
	cfg *configs.Config
}

func NewUsersRepository(db *mongo.Database, cfg *configs.Config) domain.UserRepository {
	return &usersRepositoryImpl{
		db:  db,
		cfg: cfg,
	}
}

func (u *usersRepositoryImpl) GetUserById(id primitive.ObjectID) (*models.Users, error) {
	filter := bson.M{"_id": id}
	return collectUsersDocument(filter, u.db, u.cfg)

}

func (u *usersRepositoryImpl) GetUserByEmail(email string) (*models.Users, error) {
	filter := bson.M{"email": email}
	return collectUsersDocument(filter, u.db, u.cfg)

}

func (u *usersRepositoryImpl) GetUserByUsername(username string) (*models.Users, error) {
	filter := bson.M{"username": username}
	return collectUsersDocument(filter, u.db, u.cfg)
}

func (u *usersRepositoryImpl) CreateUser(user *entities.SignupRequest) (*models.Users, error) {
	user.CreatedAt = time.Now().UTC()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := u.db.Collection(u.cfg.Mongo.Collection).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return collectUsersDocument(bson.M{"email": user.Email}, u.db, u.cfg)
}

func collectUsersDocument(filter bson.M, db *mongo.Database, cfg *configs.Config) (*models.Users, error) {
	var user models.Users
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := db.Collection(cfg.Mongo.Collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
