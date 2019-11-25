package routers

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/HFO4/cloudreve/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"testing"
)

var mock sqlmock.Sqlmock
var memDB *gorm.DB

// TestMain 初始化数据库Mock
func TestMain(m *testing.M) {
	// 设置gin为测试模式
	gin.SetMode(gin.TestMode)

	// 初始化sqlmock
	var db *sql.DB
	var err error
	db, mock, err = sqlmock.New()
	if err != nil {
		panic("An error was not expected when opening a stub database connection")
	}

	// 初始话内存数据库
	model.Init()
	memDB = model.DB

	model.DB, _ = gorm.Open("mysql", db)
	defer db.Close()

	m.Run()
}

func switchToMemDB() {
	model.DB = memDB
}