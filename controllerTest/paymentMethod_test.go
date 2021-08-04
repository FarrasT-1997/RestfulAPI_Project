package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"restfulAPI/config"
	"restfulAPI/modelsTest"
	"testing"

	"github.com/labstack/echo"
)

func testGetPaymentController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	fmt.Println(c)
	bookController(c)
	fmt.Println(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	fmt.Println(rec.Body)
	fmt.Println(body)
	var payments []modelsTest.Payment
	if err := json.Unmarshal(body, &payments); err != nil {
		t.Error(err)
	}
	if len(payments) != 1 {
		t.Errorf("expected one book, got: %#v", payments)
		return
	}
	if payments[0].Name != "Bank BCA" {
		t.Errorf("expected Bank BCA, got: %#v", payments[0].Name)
	}
}

func TestDBGetPaymentController(t *testing.T) {
	// bikin db
	db, err := config.CreateDBTest()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&modelsTest.Payment{})
	db.Delete(&modelsTest.Payment{}, "1=1")
	m := modelsTest.NewGormPaymentModel(db)
	// bikin controller
	paymentController := CreateGetPaymentController(m)
	if err != nil {
		t.Error(err)
	}
	// insert data baru
	m.Insert(modelsTest.Payment{Name: "Bank BCA"})
	// test controller
	testGetPaymentController(t, paymentController)
	db.Delete(&modelsTest.Payment{}, "1=1")
}
