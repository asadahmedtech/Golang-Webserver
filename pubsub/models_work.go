package main

import(
	"fmt"
	// "encodings/json"
	"sync"
	"log"
	"errors"
)

type Data string

type Client struct{
	Topics map[string] []Subscription
	SubscriptionList []*Subscription

}

type Publisher struct{
	ID string
	PublishingChannel chan<- Data
}

type Topic struct{
	ID string
	Subscriptions []Subscription
}

type Subscription struct{
	ID string
	Pending []Data
	Datatosubs *chan Data
	Mu sync.Mutex
}

type Subscriber struct{
	TopicID string
	Datafromsubs <-chan Data
}

func CreateNewClient() *Client{
	m := make(map[string] []Subscription)
	return &Client{
		Topics: m,
	}
}

func (c *Client) CreateTopic(topic string) error{
	_, ok := c.Topics[topic]
	temp := []Subscription{}
	if !ok{
		c.Topics[topic] = temp
		return nil
	}
	return errors.New("Topic already exist")
}

func (c *Client) CreateSubscription(topic, ID string) (*Subscription, error){
	_, ok := c.Topics[topic]
	if !ok{
		return nil, errors.New("Topic doesnot exist")
	}
	tempChan := make(chan Data)
	tempSub := Subscription{
		ID: ID,
		Datatosubs: &tempChan,
	}
	c.Topics[topic] = append(c.Topics[topic], tempSub)
	c.SubscriptionList = append(c.SubscriptionList, &tempSub)

	return &tempSub, nil
}

func (c *Client) SendtoSubscription(topic string, message Data){
	for _, subs := range c.Topics[topic]{
		subs.Mu.Lock()
		*subs.Datatosubs <- message
		subs.Mu.Unlock()
	}
}

func (s *Subscription) GetBufferMessages(){
	message := <-*s.Datatosubs
	s.Pending = append(s.Pending, message)
}

func (s *Subscription) GetMessage() Data{
	topMessage := s.Pending[0]
	s.Pending = s.Pending[1:]
	return topMessage
}

func (c *Client) PublishMessage(topic string, message Data) error{
	_, ok := c.Topics[topic]
	if !ok{
		return errors.New("Topic doesnot exist")
	}

	//Save the message and send the Acknoledgement
	go c.SendtoSubscription(topic, message)

	return nil
} 

func (c *Client) SubscriptionIterator(){
	var wg sync.WaitGroup
	for _, subs := range c.SubscriptionList{
		wg.Add(1)
		go func(){
			defer wg.Done()
			subs.GetBufferMessages()
		}()
	}
	wg.Wait()
}

func main(){
	newClient := CreateNewClient()
	err := newClient.CreateTopic("Topic1")
	if err!=nil{
		log.Println(err)
	}

	sub1, err := newClient.CreateSubscription("Topic1", "Sub1")
	if err!=nil{
		log.Println(err)
	}
	sub2, err := newClient.CreateSubscription("Topic1", "Sub2")
	if err!=nil{
		log.Println(err)
	}
	// sub1.GetBufferMessages()
	// sub2.GetBufferMessages()

	err = newClient.PublishMessage("Topic1", "A1")
	if err!=nil{
		log.Println(err)
	}

	err = newClient.PublishMessage("Topic1", "A2")
	if err!=nil{
		log.Println(err)
	}

	// TODO: Make it fetch automatically and with Ack when new data is pusher 
	// newClient.SubscriptionIterator()
	sub1.GetBufferMessages()
	sub2.GetBufferMessages()
	sub1.GetBufferMessages()
	sub2.GetBufferMessages()

	fmt.Println("SUB1 : ", sub1.GetMessage())
	fmt.Println("SUB2 : ", sub2.GetMessage())
	fmt.Println("SUB1 : ", sub1.GetMessage())
	fmt.Println("SUB2 : ", sub2.GetMessage())
}