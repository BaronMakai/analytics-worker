package helpers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/kafka-go"
)

// Error represents an error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Validate validates a struct
func Validate(s interface{}) error {
	return validator.New().Struct(s)
}

// GetRandomString returns a random string of a given length
func GetRandomString(length int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// WaitGroupWrapper is a wrapper for a WaitGroup
type WaitGroupWrapper struct {
	sync.WaitGroup
}

// WaitGroupWrapperFunc is a function that takes a WaitGroupWrapper
type WaitGroupWrapperFunc func(*WaitGroupWrapper)

// WaitGroupWrapperFuncs is a slice of WaitGroupWrapperFuncs
type WaitGroupWrapperFuncs []WaitGroupWrapperFunc

// NewWaitGroupWrapper returns a new WaitGroupWrapper
func NewWaitGroupWrapper() *WaitGroupWrapper {
	return &WaitGroupWrapper{}
}

// WaitGroupWrapperFuncsFunc is a function that takes a WaitGroupWrapperFuncs
type WaitGroupWrapperFuncsFunc func(WaitGroupWrapperFuncs)

// NewWaitGroupWrapperFuncs returns a new WaitGroupWrapperFuncs
func NewWaitGroupWrapperFuncs() WaitGroupWrapperFuncs {
	return []WaitGroupWrapperFunc{}
}

// NewHTTPClient returns a new HTTP client
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// NewKafkaWriter returns a new Kafka writer
func NewKafkaWriter(topic string) (*kafka.Writer, error) {
	w := &kafka.Writer{
		Topic: topic,
		KafkaBrokers: []string{
			"localhost:9092",
		},
	}
	return w, nil
}

// NewSQLxDB returns a new SQLxDB
func NewSQLxDB(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewSQLxDBContext returns a new SQLxDB context
func NewSQLxDBContext(db *sqlx.DB) func(context.Context) (*sqlx.DB, error) {
	return func(ctx context.Context) (*sqlx.DB, error) {
		return db, nil
	}
}

// NewMuxRouter returns a new Mux router
func NewMuxRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}).Methods("GET")
	return r
}

// NewLogger returns a new logger
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
}