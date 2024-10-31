package queue

import (
	"fmt"
)

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (order *Order) New(
	priority int,
	quantity int,
	product string,
	customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName

}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func QueueExec() {
	var queue Queue
	queue = make(Queue, 0)
	order1 := &Order{}
	priority1 := 2
	quantity1 := 20
	product1 := "Computer"
	customerName1 := "Greg White"
	order1.New(priority1, quantity1, product1, customerName1)
	order2 := &Order{}
	priority2 := 1
	quantity2 := 10
	product2 := "Monitor"
	customerName2 := "John Smith"
	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)
	queue.Add(order2)

	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}

}
