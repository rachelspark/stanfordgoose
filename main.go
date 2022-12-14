package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"stanford-goose/datasource"
	"stanford-goose/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected to be called with a subcommand")
	}

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverData := serverCmd.String("data", "", "path or url for the data file")
	serverStatic := serverCmd.String("static", "", "path to static website files")
	serverLocal := serverCmd.Bool("local", false, "set to use local mode")

	switch os.Args[1] {
	case "download":
		var courses []datasource.Course
		courses = datasource.DownloadCourses()
		fmt.Printf("Downloaded %v course", len(courses))
		coursesText, _ := json.Marshal(courses)
		if err := os.WriteFile("data/courses.json", coursesText, 0644); err != nil {
			log.Fatalf("failed to write courses.json: %v", err)
		}
	case "server":
		serverCmd.Parse(os.Args[2:])
		if *serverData == "" {
			log.Fatal("server requires a -data file")
		}
		server.Run(*serverData, *serverStatic, *serverLocal)
	default:
		log.Fatal("unexpected subcommand")
	}
}
