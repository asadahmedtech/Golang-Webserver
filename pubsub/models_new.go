package main

import(
	"fmt"
	// "encodings/json"
	"sync"
	"log"
	"errors"
	"time"

	guuid "github.com/google/uuid"

)

type PublishMessageSource struct{
	ID string
	Payload Data
	Ack bool
}
type Data struct{
	ID string
	Payload string
}

type Client struct{
	Topics map[string] []Subscription
	SubscriptionList []*Subscription
	Storage map[string][]PublishMessageSource
}

// type Publisher struct{
// 	ID string
// 	PublishingChannel chan<- Data
// }

// type Topic struct{
// 	ID string
// 	Subscriptions []Subscription
// }

type Subscription struct{
	ID string
	Pending []Data
	Datatosubs *chan Data
	Mu sync.Mutex
}

// type Subscriber struct{
// 	TopicID string
// 	Datafromsubs <-chan Data
// }

func CreateNewClient() *Client{
	m := make(map[string] []Subscription)
	m1 := make(map[string][]PublishMessageSource)

	return &Client{
		Topics: m,
		Storage: m1,
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
	tempChan := make(chan Data, 10)
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
		log.Printf("%s : Sending Data to Subscription Buffer", subs.ID)
		subs.Mu.Lock()
		*subs.Datatosubs <- message
		subs.Mu.Unlock()
	}
}

func (s *Subscription) GetBufferMessages()(bool){
	log.Printf("%s : Fetching Buffer Messages", s.ID)
	select{
	case message := <-*s.Datatosubs:

		s.Pending = append(s.Pending, message)
		return true
	default:
		log.Printf("%s recieved no message ", s.ID)
		return false
	}
}

func (s *Subscription) GetMessage() (Data, error){
	log.Printf("%s : Sending Messages to Subscriber", s.ID)

	if len(s.Pending) == 0{
		ok := s.GetBufferMessages()
		if !ok{
			return Data{}, errors.New("No message recieved")
		}
	}
	topMessage := s.Pending[0]
	s.Pending = s.Pending[1:]
	return topMessage, nil
}

// func (s *Subscription) Acknoledge(){

// }
func (c *Client) PublishMessage(topic string, message Data) (bool, error){
	_, ok := c.Topics[topic]
	if !ok{
		return false, errors.New("Topic doesnot exist")
	}
	message.ID = guuid.New().String()
	c.Storage[topic] = append(c.Storage[topic], PublishMessageSource{ID: message.ID, Payload: message,})
	//Save the message and send the Acknoledgement
	go c.SendtoSubscription(topic, message)

	return true, nil
} 

func (c *Client) SubscriptionIterator(){
	var wg sync.WaitGroup
	for i := range c.SubscriptionList{
		wg.Add(1)
		index := i
		// subs := c.SubscriptionList[index]
		
		// log.Printf("%s : SubscriptionIterator", subs.ID)
		go func(){
			sub := c.SubscriptionList[index]
			// log.Printf("%s : SubscriptionIterator inside func", sub.ID)

			defer wg.Done()
			sub.GetBufferMessages()
		}()
	}
	wg.Wait()
}

func main(){
	// var wg sync.WaitGroup

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
	sub3, err := newClient.CreateSubscription("Topic1", "Sub3")
	if err!=nil{
		log.Println(err)
	}
	// sub1.GetBufferMessages()
	// sub2.GetBufferMessages()

	// newClient.SendtoSubscription("Topic1", "A1")
	ack, err := newClient.PublishMessage("Topic1", Data{Payload: "A1",})
	if err!=nil{
		log.Println(err)
	}
	log.Println("Topic1", ack)

	 // newClient.SendtoSubscription("Topic1", "A2")
	ack, err = newClient.PublishMessage("Topic1", Data{Payload: "A2",})
	if err!=nil{
		log.Println(err)
	}
	log.Println("Topic1", ack)

	// log.Println(newClient.SubscriptionList)

	// TODO: Make it fetch automatically and with Ack when new data is pushed 
	// go func(){
	// 	for{
	// 		newClient.SubscriptionIterator()
	// 		time.Sleep(1*time.Second)
	// 	}
	// }()

		time.Sleep(100000*time.Nanosecond)

	// sub1.GetBufferMessages()
	// sub2.GetBufferMessages()
	// sub1.GetBufferMessages()
	// sub2.GetBufferMessages()
	for i:=0;i<2;i++{
		m, err := sub1.GetMessage()
		fmt.Println("SUB1 : ", m, err)
		
		m, err = sub2.GetMessage()
		fmt.Println("SUB2 : ", m, err)

		m, err = sub3.GetMessage()
		fmt.Println("SUB3 : ", m, err)
	}
}