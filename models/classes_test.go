package models

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"strings"
	"testing"
)

func TestGetClasses(t *testing.T) {
	// Data needed to test
	columns := []string{"id", "name"}
	testID := 1
	testName := "test class"

	//Start the mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s occured when opening the stub database connection", err.Error())
	}
	defer db.Close() // nolint: errcheck

	mock.ExpectQuery("SELECT id, name FROM class WHERE deleted = false").WillReturnRows(sqlmock.NewRows(columns).AddRow(testID, testName))
	mock.ExpectClose()

	classes, err := GetClasses(db)
	if err != nil {
		t.Fatalf("TestGetClasses: err -- %s", err)
	}

	// Check the results
	if classes[0].ID != testID {
		t.Fatalf("TestGetClasses: %d is not equal to %d", testID, classes[0].ID)
	}
	if !strings.EqualFold(testName, classes[0].Name) {
		t.Fatalf("TestGetClasses: %s is not equal to %s", testName, classes[0].Name)
	}
}

func TestUpdateClassName(t *testing.T) {
	columns1 := []string{"name"}
	testName := "Testing"
	testID := "1"

	//Start the mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error %s occurred when opening the stub database connection", err.Error())
	}
	defer db.Close() // nolint: errcheck

	mock.ExpectQuery("SELECT (.+) FROM class WHERE (.+)").WillReturnRows(sqlmock.NewRows(columns1).AddRow("hello"))
	mock.ExpectQuery("UPDATE class SET (.+) WHERE (.+)").WillReturnRows(sqlmock.NewRows(columns1).AddRow(""))
	mock.ExpectClose()

	err = UpdateClassName(db, testID, testName)
	if err != nil {
		t.Fatalf("TestUpdateClassName: err -- %s", err.Error())
	}
}