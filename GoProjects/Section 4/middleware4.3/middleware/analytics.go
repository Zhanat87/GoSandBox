package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	ErrInvalidID    = errors.New("Invalid ID")
	ErrInvalidEmail = errors.New("Invalid email")
)

// CreateLogger creates a new logger that writes to the given filename.
func CreateLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

// Time runs the next function in the chain
func Time(logger *log.Logger, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)
		logger.Println(elapsed)
	})
}

// Recover recovers from any panicking goroutine
func Recover(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				switch err {
				case ErrInvalidEmail:
					http.Error(w, ErrInvalidEmail.Error(), http.StatusUnauthorized)
				case ErrInvalidID:
					http.Error(w, ErrInvalidID.Error(), http.StatusUnauthorized)
				default:
					http.Error(w, "Unknown error, recovered from panic", http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
