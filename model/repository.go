package model

import (
	"log"

	"github.com/Eldius/things-manager-go/config"
	"github.com/jinzhu/gorm"

	// I need this
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	if db == nil {
		db = initDB()
	}
	return &Repository{
		db: db,
	}
}

func NewRepositoryCustom(_db *gorm.DB) *Repository {
	db = _db
	return &Repository{
		db: db,
	}
}

// GetDB gets a database connection
func (r *Repository) GetDB() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}

func (r *Repository) ListThings() []Thing {
	var results []Thing
	r.db.Find(&results)

	return results
}

func (r *Repository) SaveThing(t *Thing) *Thing {
	r.db.Save(t)
	return t
}

func (r *Repository) GetThing(id int) *Thing {
	var t Thing
	r.db.First(&t, id)
	return &t
}

func (r *Repository) GetThingByName(name string) *Thing {
	var t Thing
	log.Println(name)
	r.db.Where(&Thing{Name: name}).First(&t)
	log.Println(t)
	return &t
}

func initDB() *gorm.DB {
	db, err := gorm.Open(config.GetDBEngine(), config.GetDBURL())
	if err != nil {
		log.Printf("failed to connect database to app database\n- engine: %s\n- url: %s\n", config.GetDBEngine(), config.GetDBURL())
		panic(err.Error())
	}
	if config.GetDBLogQueries() {
		db.LogMode(true)
	}
	db.AutoMigrate(&Thing{})
	return db
}
