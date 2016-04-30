package orm

import (
	"os"
	"testing"

	_ "github.com/kwiscale/orm/dialects/sqlite"
)

func up() {
	Initialize("sqlite", "test.db")
}

func down() {
	os.Remove("test.db")
}

type User struct {
	Model
	Name string
	Age  int
}

func TestCreateData(t *testing.T) {
	defer down()
	up()

	db := Get()
	db.Create(&User{
		Name: "Foo",
		Age:  42,
	})

}
