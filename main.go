package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if len(os.Args) < 3 {
		log.Fatal("Need source and destination files as arguments")
	}
}

func main() {
	src := os.Args[1]
	dest := os.Args[2]

	fmt.Printf("Source: %s\n", src)
	fmt.Printf("Destination: %s\n", dest)

	var envs map[string]string
	envs, err := godotenv.Read(src)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for key, val := range envs {
		var env string

		if len(getEnvVar(key)) > 0 {
			env = getEnvVar(key)
		} else {
			env = val
		}

		fmt.Printf("Key: %s, Value: %s\n", key, env)
		fmt.Fprintf(file, "%s=\"%s\"\n", key, env)
	}

	fmt.Println("====================================")
	fmt.Printf("%s created successfully\n", dest)
	fmt.Println("====================================")
}

func getEnvVar(key string) string {
	return os.Getenv(key)
}
