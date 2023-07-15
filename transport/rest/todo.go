package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getAll")
	defer LogRecover()
	todos := serviceComposer.ToDos.GetAll(context.Background())
	reply := fetchResponse{Success: true, Data: todos}
	json.NewEncoder(w).Encode(reply)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	todo := serviceComposer.ToDos.Get(id)
	reply := fetchResponse{Success: true, Data: todo}
	json.NewEncoder(w).Encode(reply)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	var data ItemPatchRequest
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &data)
	id := serviceComposer.ToDos.Create(data.Title, data.Description)
	reply := fetchResponse{Success: true, Data: id}
	json.NewEncoder(w).Encode(reply)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var data ItemPatchRequest
	reply := fetchResponse{Success: false}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := json.NewDecoder(r.Body).Decode(&data)
	if err == nil {
		reply.Success = serviceComposer.ToDos.Update(id, data.Title, data.Description, data.Done)
	}
	json.NewEncoder(w).Encode(reply)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := serviceComposer.ToDos.Delete(id)
	reply := fetchResponse{Success: ok}
	json.NewEncoder(w).Encode(reply)
}
