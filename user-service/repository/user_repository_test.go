package repository

import (
	"context"
	"fmt"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
)

var (
	DB  *mongo.Database
	Cfg *configs.Config
)

func TestUserInsertOne(t *testing.T) {
	userRepo := NewUsersRepository(DB, Cfg)
	data := &entities.SignupRequest{
		Username: "test_username",
		Email:    "test@gmail.com",
		Password: "test_pass",
	}
	user, err := userRepo.CreateUser(data)
	assert.Nilf(t, err, "Error while inserting user")
	assert.Equal(t, "test_username", user.Username)
}

func TestUserFindOneByUsername(t *testing.T) {
	userRepo := NewUsersRepository(DB, Cfg)
	user, err := userRepo.GetUserByUsername("test_username")
	assert.Nilf(t, err, "Error while finding user by username")
	assert.Equal(t, "test_username", user.Username)
}

func TestUserFindOneByEmail(t *testing.T) {
	userRepo := NewUsersRepository(DB, Cfg)
	user, err := userRepo.GetUserByUsername("test_username")
	assert.Nilf(t, err, "Error while finding user by email")
	assert.NotEqual(t, "email", user.Email)
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Env:        []string{"MONGO_INITDB_ROOT_USERNAME=" + Cfg.Mongo.MongoUser, "MONGO_INITDB_ROOT_PASSWORD=" + Cfg.Mongo.MongoPass},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	defer func() {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()
	err = pool.Retry(func() error {
		var err error
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
			fmt.Sprintf("mongodb://%s:%s@127.0.0.1:%s",
				Cfg.Mongo.MongoUser, Cfg.Mongo.MongoPass,
				resource.GetPort("27017/tcp"))),
		)
		if err != nil {
			return err
		}
		if err := client.Ping(context.TODO(), nil); err != nil {
			return err
		}
		DB = client.Database(Cfg.Mongo.MongoDB)
		return nil
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	exitCode := m.Run()
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(exitCode)
}
func init() {
	var err error
	Cfg, err = configs.NewConfig()
	if err != nil {
		log.Fatalf("configs.NewConfig error: %v", err)
	}
}
