package main

import (
	"fmt"
	"sync"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
)

func main() {
	wg := sync.WaitGroup{}
	itemChan := make(chan todo_store.Todo)
	statusChan := make(chan bool)

	items := createTodoList()

	wg.Add(2)

	go printDescriptions(&wg, itemChan, statusChan, items)
	go printStatuses(&wg, itemChan, statusChan, items)

	wg.Wait()
}

func printDescriptions(wg *sync.WaitGroup, itemChan chan todo_store.Todo, statusChan chan bool, items []todo_store.Todo) {
	defer wg.Done()

	for _, item := range items {
		fmt.Println(item.Description)

		itemChan <- item
		<-statusChan
	}
}

func printStatuses(wg *sync.WaitGroup, itemChan chan todo_store.Todo, statusChan chan bool, items []todo_store.Todo) {
	defer wg.Done()

	for _, item := range items {
		<-itemChan

		if item.Complete {
			fmt.Println("Complete")
		} else {
			fmt.Println("Incomplete")
		}

		statusChan <- item.Complete
	}
}

func createTodoList() []todo_store.Todo {
	todoStore := todo_memory_store.TodoStore{}
	for i := 1; i <= 10; i++ {
		id, _ := todoStore.Create(fmt.Sprintf("Item %d", i))
		if i%2 == 0 {
			todoStore.ToggleItemStatus(id)
		}
	}

	return todoStore.GetItems()
}
