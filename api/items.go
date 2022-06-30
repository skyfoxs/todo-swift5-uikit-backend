package main

import (
	"encoding/json"
	"net/http"
)

var t = todo{
	items: []todoItem{},
}

func (app *application) itemsHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Printf("Get items")
	response, err := json.Marshal(t.items)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	app.logger.Printf("Response %s", response)
}

func (app *application) updateItemsHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Printf("Update items")
	var i []todoItem
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.items = i
	app.logger.Printf("Items %#v", i)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
