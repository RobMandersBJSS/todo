package web_server

import "fmt"

type Action struct {
	Action      string `json:"action"`
	ID          string `json:"id"`
	Description string `json:"description"`
}

func (s *Server) createItem(messageJSON Action) {
	_, err := s.todoStore.Create(messageJSON.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to create item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}

func (s *Server) updateItem(messageJSON Action) {
	err := s.todoStore.UpdateItem(messageJSON.ID, messageJSON.Description)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to update item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}

func (s *Server) toggleItemStatus(messageJSON Action) {
	err := s.todoStore.ToggleItemStatus(messageJSON.ID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to change item status - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}

func (s *Server) deleteItem(messageJSON Action) {
	err := s.todoStore.Delete(messageJSON.ID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error attempting to delete item - %s", err)
		s.sendWebsocketErrorMessage(errorMessage)
	}
}
