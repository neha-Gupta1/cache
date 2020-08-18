package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Node single element of linklist
type Node struct {
	// mux   sync.mutex
	Value string
	ID    int
	Next  *Node
}

// Data found from db
// swagger:model
type Data struct {
	ID    int    `json:"id" gorm:"primary_key;auto_increment"`
	Value string `json:"value"`
}

//Server struct
type Server struct {
	DBServer *gorm.DB
}

// Initialize will intialize DB
func (Server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	Server.DBServer, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	Server.DBServer.Debug().AutoMigrate(&Data{})
}

// Server getter
func (Server *Server) Server() *gorm.DB {
	return Server.DBServer
}
