package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/inigofu/shippy-user-service/proto/auth"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
)

var (
	srv micro.Service
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})
	db.AutoMigrate(&pb.Role{})
	db.AutoMigrate(&pb.Menu{})
	db.AutoMigrate(&pb.Badge{})
	db.AutoMigrate(&pb.Wrapper{})
	db.AutoMigrate(&pb.Atributes{})
	db.AutoMigrate(&pb.Form{})
	db.AutoMigrate(&pb.FormSchema{})
	db.AutoMigrate(&pb.Buttons{})
	db.AutoMigrate(&pb.Class{})
	db.AutoMigrate(&pb.Values{})
	db.AutoMigrate(&pb.SelectOptions{})

	dblog, _ := strconv.ParseBool(os.Getenv("DB_LOG"))
	db.LogMode(dblog)

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv = micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.auth"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Will comment this out now to save having to run this locally
	// publisher := micro.NewPublisher("user.created", srv.Client())

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
