package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
)

func increaseViewCount(conn redis.Conn) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := conn.Do("INCR", "view-counter")

		if err != nil {
			fmt.Println("error occurred while running incr command: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	port := envy.Get("PORT", "8080")
	redisURL := envy.Get("REDIS_URL", "redis://localhost:6379")
	conn, err := redis.DialURL(redisURL)

	if err != nil {
		panic(fmt.Errorf("error connecting to redis: %v", err))
	}

	router := httprouter.New()
	router.POST("/view", increaseViewCount(conn))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
