package read_file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"todo/modules/store"
)

func ReadTodosFromFile(fileName string) ([]store.Todo, error) {
	file, readFileError := os.Open(fileName)
	if readFileError != nil {
		message := fmt.Sprintf("error when reading file from %q: %v", fileName, readFileError)
		return nil, errors.New(message)
	}
	defer file.Close()

	bytes, readBytesError := io.ReadAll(file)
	if readBytesError != nil {
		message := fmt.Sprintf("error reading bytes from file %q: %v", fileName, readBytesError)
		return nil, errors.New(message)
	}

	var todos []store.Todo

	unmarshalError := json.Unmarshal(bytes, &todos)
	if unmarshalError != nil {
		message := fmt.Sprintf("error unmarchaling json from file %q: %v", fileName, unmarshalError)
		return nil, errors.New(message)
	}

	return todos, nil
}
