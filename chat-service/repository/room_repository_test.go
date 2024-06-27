package repository

import (
	"context"
	"fmt"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
)

var (
	DB    *mongo.Database
	Cfg   *configs.Config
	REDIS *redis.Client
)

func TestRoomInsertOne(t *testing.T) {
	redisRepo := NewRedisRepository(REDIS)
	roomRepo := NewRoomRepository(DB, redisRepo, Cfg)
	data := &entities.RoomRequest{
		Name: "golang",
	}
	room, err := roomRepo.Create(data)
	assert.Nilf(t, err, "Error while inserting user")
	assert.Equal(t, "golang", room.Name)
}

func TestRoomFindOneByName(t *testing.T) {
	redisRepo := NewRedisRepository(REDIS)
	roomRepo := NewRoomRepository(DB, redisRepo, Cfg)
	room, err := roomRepo.GetByName("golang")
	assert.Nilf(t, err, "Error while finding user by username")
	assert.Equal(t, "golang", room.Name)
	_, err = roomRepo.GetByName("python")
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
}

func TestRoomList(t *testing.T) {
	redisRepo := NewRedisRepository(REDIS)
	roomRepo := NewRoomRepository(DB, redisRepo, Cfg)
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
	redisResource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "redis",
		Tag:        "7.2",
		Env:        []string{"REDIS_PASSWORD=" + Cfg.Redis.Password},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start Redis resource: %s", err)
	}
	defer func() {
		if err = pool.Purge(redisResource); err != nil {
			log.Fatalf("Could not purge Redis resource: %s", err)
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
		log.Fatalf("Could not connect to docker and mongo: %s", err)
	}

	err = pool.Retry(func() error {
		REDIS = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("127.0.0.1:%s", redisResource.GetPort("6379/tcp")),
			Password: Cfg.Redis.Password,
			DB:       0,
		})
		if err := REDIS.Ping(context.TODO()).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Redis Could not connect to docker and redis: %s", err)
	}
	exitCode := m.Run()
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge MongoDB resource: %s", err)
	}
	if err = pool.Purge(redisResource); err != nil {
		log.Fatalf("Could not purge Redis resource: %s", err)
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
