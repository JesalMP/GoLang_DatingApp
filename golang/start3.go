package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DEFINED DATA STRUCTURES

type users struct {
	Id       int
	Name     string
	Location int64
	Gender   string
	Email    string
}
type likes struct {
	Id           int
	Who_likes    int64
	Who_is_liked int64
}
type pair struct {
	Id1 int
	Id2 int
}
type matchstruct struct {
	User1 users
	User2 users
}

// POSTGRE SQL DATABASE CONNECTION URL

var dsn string = "host=localhost user=postgres password=jes123 dbname=goLangjesalMPtask port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// DEFINE FUNCTIONS

func match(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MATCH users HIT")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	var likesTemp []likes
	var map_1 = make(map[int]int)
	db.Find(&likesTemp)

	for _, likesTemp := range likesTemp {
		map_1[int(likesTemp.Who_likes)] = int(likesTemp.Who_is_liked)
	}
	var pairids []pair
	for k := range map_1 {
		var x = map_1[k]
		if map_1[x] == k {
			var pairTemp pair
			pairTemp.Id1 = k
			pairTemp.Id2 = x
			pairids = append(pairids, pairTemp)
			delete(map_1, k)
			delete(map_1, x)
		}
	}
	var matches []matchstruct
	for _, pairids := range pairids {
		var matchTemp matchstruct
		var userTemp1 users
		var userTemp2 users
		db.First(&userTemp1, pairids.Id1)
		matchTemp.User1 = userTemp1
		db.First(&userTemp2, pairids.Id2)
		matchTemp.User2 = userTemp2
		matches = append(matches, matchTemp)
	}
	json.NewEncoder(w).Encode(matches)
	sqlDB.Close()

}

func usersindistance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Users within distance HIT")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	vars := mux.Vars(r)
	user1id := vars["userid"]
	disti := vars["distance"]
	//var dist int64
	dists, err := strconv.Atoi(disti)
	var user1 users
	db.First(&user1, user1id)

	user1posi := user1.Location
	var usersTemp []users
	db.Find(&usersTemp)
	var usersindistance []users

	for _, usersTemp := range usersTemp {
		if usersTemp.Id != user1.Id {
			user2dist := usersTemp.Location
			var k = user2dist - user1posi
			if k < 0 {
				k = k * -1
			}
			if int(k) <= dists {
				usersindistance = append(usersindistance, usersTemp)
			}
		}
	}

	json.NewEncoder(w).Encode(usersindistance)
	sqlDB.Close()

}

func userswithname(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Users With Name HIT")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	vars := mux.Vars(r)
	nametagi := vars["nametag"]
	nametag := strings.ToLower(nametagi)
	var usersTemp []users
	db.Find(&usersTemp)
	var userswithnameinit []users
	for _, usersTemp := range usersTemp {
		nami := usersTemp.Name
		nameTemp := strings.ToLower(nami)
		if strings.Contains(nameTemp, nametag) {
			userswithnameinit = append(userswithnameinit, usersTemp)
		}
	}
	json.NewEncoder(w).Encode(userswithnameinit)
	sqlDB.Close()
}

//DEFINE ENDPOINTS TO RESPECTIVE FUNCTIONS

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/availablematches", match).Methods("GET")
	myRouter.HandleFunc("/usersindistance/{userid}/{distance}", usersindistance).Methods("GET")
	myRouter.HandleFunc("/usernamesearch/{nametag}", userswithname).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// DEFINE SCHEMA MAPPING

func initialMigration() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	db.AutoMigrate(&likes{})
	db.AutoMigrate(&users{})
	sqlDB.Close()
}

// MAIN FUNCTION

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	handleRequests()
}
