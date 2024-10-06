package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO
type ToDo struct {
	UserID    int    `json:"userId" xml:"userId"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type IData interface {
	GetData() (*ToDo, error)
}

type RemoteService struct {
	Remote IData
}

func (rs *RemoteService) CallRemoteService() (*ToDo, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct {
}

func (jb *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo ToDo

	err = json.Unmarshal(body, &todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `
    <?xml version="1.0" encoding="UTF-8 ?>
    <root>
        <userId>1</userId>
        <id>1</id>
        <title>delectur aut autem</title>
        <completed>false</completed>
    </root>
    `

	var todo ToDo

	err := xml.Unmarshal([]byte(xmlFile), &todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	// No adapter
	todo := getRemoteData()
	fmt.Println("Todo without adapter:\t", todo.ID, todo.Title)

	// With Adapter, using JSON
	jsonBackend := &JSONBackend{}

	jsonAdapter := &RemoteService{
		Remote: jsonBackend,
	}

	tdFromJSON, err := jsonAdapter.CallRemoteService()

	if err != nil {
		log.Print(err)
	}

	fmt.Println("From JSON Adapter:\t", tdFromJSON.ID, tdFromJSON.Title)

	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{
		Remote: xmlBackend,
	}

	tdFromXML, _ := xmlAdapter.CallRemoteService()

	fmt.Println("From XML Adapter:\t", tdFromXML.ID, tdFromXML.Title)

}

func getRemoteData() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo ToDo

	err = json.Unmarshal(body, &todo)

	if err != nil {
		log.Fatal(err)
	}

	return &todo
}
