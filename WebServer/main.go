package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	connectToDatabase()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":1521", router))
}
