package db

import (
	chatModel "Chat/chats/models"
	conversationModel "Chat/conversations/models"
	fileModel "Chat/fileUpload/models"
	"Chat/users/models"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type SqlType string

const (
	SQLITE   SqlType = "sqlite"   //Use to initiate a SQLITE3 db connection
	POSTGRES SqlType = "postgres" //Use to initiate a postgresSql connection
	DEFAULT  SqlType = "DEFAULT"  //pass to return already initiated db engine
)

/*
This function returns DB object with connections to provided type of DB's.

the default connections are made to predefined DB's.

	params: SQL_TYPE from helpers_sql package

	returns: DB objects with connections established
*/
func getSqlConnection(sqlType SqlType) *gorm.DB {
	var dsn string
	var dialect gorm.Dialector

	switch sqlType {
	case SQLITE:
		dsn = "data/test.db"
		createFolderIfNotExists("data")
		dialect = sqlite.Open(dsn)
	case POSTGRES:
		dsn = ""
		dialect = postgres.Open(dsn)
	}

	if dialect == nil {
		fmt.Println("Could not Find mentioned SQLSERVER type")
		return nil
	}

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		fmt.Println("connection error")
		return nil
	}
	return db
}

func createFolderIfNotExists(folderName string) {
	_, err := os.Stat(folderName)
	if err == nil {
		return
	}
	err = os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		print("could not create folder")
	}
}

// NewDb
//
// This function returns the required to be connected instance of DB object
func NewDb(sqlType SqlType) *gorm.DB {
	db := getSqlConnection(sqlType)
	return db
}

func SetUpDB() *gorm.DB {
	db := NewDb(SQLITE)

	if db == nil {
		panic("connection cannot be made")
	} else {
		fmt.Println("connection to db established")
	}

	err := db.AutoMigrate(&models.User{}, &chatModel.Chat{}, &conversationModel.Conversation{}, &fileModel.Attachment{})
	if err != nil {
		return nil
	}
	return db
}
