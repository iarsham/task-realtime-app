package repository

import (
	"context"
	"fmt"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
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

func TestRoomInsertOne(t *testing.T) {
	roomRepo := NewRoomRepository(DB, Cfg)
	data := &entities.RoomRequest{
		Name: "golang",
	}
	room, err := roomRepo.Create(data)
	assert.Nilf(t, err, "Error while inserting user")
	assert.Equal(t, "golang", room.Name)
}

func TestRoomFindOneByName(t *testing.T) {
	roomRepo := NewRoomRepository(DB, Cfg)
	room, err := roomRepo.GetByName("golang")
	assert.Nilf(t, err, "Error while finding user by username")
	assert.Equal(t, "golang", room.Name)
	_, err = roomRepo.GetByName("python")
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
}

func TestRoomList(t *testing.T) {
	roomRepo := NewRoomRepository(DB, Cfg)
	rooms, err := roomRepo.List()
	assert.Nilf(t, err, "Error while listing rooms")
	assert.Equal(t, 1, len(*rooms))
	assert.NotEqual(t, 2, len(*rooms))
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
