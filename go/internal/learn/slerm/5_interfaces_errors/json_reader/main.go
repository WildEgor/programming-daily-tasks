package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

type MyData struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	cd, err := os.Getwd()
	if err != nil {
		fmt.Printf("cd file err: %s", err.Error())
		return
	}

	fPath, err := filepath.Abs(path.Join(cd, "internal", "learn", "slerm", "5_interfaces_errors", "json_reader", "data.json"))
	if err != nil {
		fmt.Printf("path file err: %s", err.Error())
		return
	}

	data, err := os.ReadFile(fPath)
	if err != nil {
		fmt.Printf("read file err: %s", err.Error())
		return
	}

	var md []MyData
	if err := json.Unmarshal(data, &md); err != nil {
		fmt.Printf("parse data err: %s", err.Error())
		return
	}

	for i := 0; i < len(md); i++ {
		defer fmt.Println(md[i])
	}
}
