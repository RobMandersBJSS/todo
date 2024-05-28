package todo_store

type Todo struct {
	ID          string `json:"ID"`
	Description string `json:"Description"`
	Complete    bool   `json:"Complete"`
}

type TodoStore interface {
	GetItems() []Todo
	Create(description string) (string, error)
	Read(id string) (Todo, error)
	UpdateItem(id, description string) error
	ToggleItemStatus(id string) error
	Delete(id string) error
}
