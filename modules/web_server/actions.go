package web_server

import "fmt"

type Action struct {
	Action      string `json:"action"`
	ID          string `json:"id"`
	Description string `json:"description"`
}

func (s *Server) createItem(messageJSON Action) {
	fmt.Println("Create item")

	_, err := s.todoStore.Create(messageJSON.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to create item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}

	s.sendTodosToClient()
}

func (s *Server) updateItem(messageJSON Action) {
	fmt.Printf("Update item %s\n", messageJSON.ID)

	err := s.todoStore.UpdateItem(messageJSON.ID, messageJSON.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to update item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}

	s.sendTodosToClient()
}

func (s *Server) toggleItemStatus(messageJSON Action) {
	fmt.Printf("Change status of item %s\n", messageJSON.ID)

	err := s.todoStore.ToggleItemStatus(messageJSON.ID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to change item status - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}

	s.sendTodosToClient()
}

func (s *Server) deleteItem(messageJSON Action) {
	fmt.Printf("Delete item %s\n", messageJSON.ID)

	err := s.todoStore.Delete(messageJSON.ID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to delete item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}

	s.sendTodosToClient()
}
