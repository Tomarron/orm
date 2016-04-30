package orm

import (
	"os"
	"testing"

	_ "gopkg.in/kwiscale/orm.v0/dialects/sqlite"
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
