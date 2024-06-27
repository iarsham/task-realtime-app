package domain

type BrokerRepository interface {
	Publish(topic string, message []byte) error
}

type BrokerUsecase interface {
	PublishQueue(topic string, message []byte) error
}
