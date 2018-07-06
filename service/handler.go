package service

import (
	"net/http"
	"github.com/leandroandrade/go-mongodb/database"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"time"
	"runtime"
	"fmt"
	"bytes"
	"io"
)

type Handler struct {
	mongo *database.MongoDatabase
}

func NewHandler(m *database.MongoDatabase) *Handler {
	return &Handler{
		mongo: m,
	}
}

func (h Handler) Home(writer http.ResponseWriter, request *http.Request) {
	// first block
	/*var user User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		log.Printf("ERR: %v", err.Error())
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// really need this? NO, because the Server will close the request body.
	// The ServeHTTP Handler does not need to.
	defer request.Body.Close()*/

	// second block
	/*body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()

	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Printf("ERR: %v", err.Error())
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}*/

	// third block
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, request.Body); err != nil {
		log.Printf("ERR: %v", err.Error())
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// really need this? NO, because the Server will close the request body.
	// The ServeHTTP Handler does not need to.
	defer request.Body.Close()

	var user User
	if err := json.Unmarshal(buf.Bytes(), &user); err != nil {
		log.Printf("ERR: %v", err.Error())
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// common code
	user.ID = uuid.New().String()
	user.DateCreated = time.Now()

	// Types to get Mongo session
	collection := h.mongo.Get().DB("golang").C("people")
	//collection := h.mongo.Clone().DB("golang").C("people")
	//collection := h.mongo.Copy().DB("golang").C("people")

	if err := collection.Insert(user); err != nil {
		log.Printf("ERR: %v", err.Error())
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h Handler) PrintMemory(writer http.ResponseWriter, request *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memoryUsage := MemoryUsage{
		Alloc:        fmt.Sprintf("%v MB", bytesToMegabytes(m.Alloc)),
		TotalAlloc:   fmt.Sprintf("%v MB", bytesToMegabytes(m.TotalAlloc)),
		Sys:          fmt.Sprintf("%v MB", bytesToMegabytes(m.Sys)),
		Numgc:        fmt.Sprintf("%v", m.NumGC),
		HeapAlloc:    fmt.Sprintf("%v MB", bytesToMegabytes(m.HeapAlloc)),
		HeapSys:      fmt.Sprintf("%v MB", bytesToMegabytes(m.HeapSys)),
		HeapReleased: fmt.Sprintf("%v MB", bytesToMegabytes(m.HeapReleased)),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(memoryUsage)
}

func bytesToMegabytes(b uint64) uint64 {
	return b / 1024 / 1024
}

type MemoryUsage struct {
	Alloc        string `json:"alloc,omitempty"`
	TotalAlloc   string `json:"total_alloc,omitempty"`
	Sys          string `json:"sys,omitempty"`
	Numgc        string `json:"numgc,omitempty"`
	HeapAlloc    string `json:"heap_alloc,omitempty"`
	HeapSys      string `json:"heap_sys,omitempty"`
	HeapReleased string `json:"heap_released,omitempty"`
}

type User struct {
	ID          string
	Name        string    `json:"name,omitempty"`
	Age         int       `json:"age,omitempty"`
	Password    string    `json:"password,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
}
