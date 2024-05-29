package write_file

import (
	"bytes"
	"os"
	"todo/modules/print"
	"todo/modules/todo_store"
)

func WriteFile(filepath string, items ...todo_store.Todo) {
	file, _ := os.Create(filepath)

	buffer := bytes.Buffer{}
	print.PrintTodosJSON(&buffer, items...)

	file.Write(buffer.Bytes())
}
