package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"enigmacamp.com/be-enigma-laundry/mock/middleware_mock"
	"enigmacamp.com/be-enigma-laundry/mock/usecase_mock"
	"enigmacamp.com/be-enigma-laundry/model"
	"enigmacamp.com/be-enigma-laundry/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BillControllerTestSuite struct {
	suite.Suite
	bum *usecase_mock.BillUseCaseMock
	rg  *gin.RouterGroup
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *BillControllerTestSuite) SetupTest() {
	suite.bum = new(usecase_mock.BillUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middleware_mock.AuthMiddlewareMock)
}

func TestBillControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BillControllerTestSuite))
}

var dummyBill = model.Bill{
	Id:       "1",
	BillDate: time.Now(),
	Customer: model.Customer{
		Id:   "1",
		Name: "Jojo",
	},
	User: model.User{
		Id:   "1",
		Name: "Shinta",
	},
	BillDetails: []model.BillDetail{
		{
			Id:     "1",
			BillId: "1",
			Product: model.Product{
				Id:    "1",
				Name:  "Cuci + Setrika",
				Price: 10000,
				Type:  "Kg",
			},
			Qty:   1,
			Price: 10000,
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayload = dto.BillRequestDto{
	CustomerId: "1",
	UserId:     "1",
	BillDetails: []model.BillDetail{
		{
			Product: model.Product{Id: "1"},
			Qty:     1,
		},
	},
}

var mockTokenJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDExNzk1ODEsImlhdCI6MTcwMTAxNDU3MCwiaXNzIjoidXNlciIsInJvbGUiOiJlbXBsb3llZSIsInNlcnZpY2VzIjpudWxsLCJ1c2VySWQiOiI3ZTgwOTdkOC1mZWYxLTQ5ZWQtOWY3Yi1jZWNjOTUyMGNjYTgifQ.7jKIMS_WvP6i25HPBuITtvBhtmpw-RC-pueqkZpzcqs"

func (suite *BillControllerTestSuite) TestCreateBill_Success() {
	suite.bum.On("RegisterNewBill", mockPayload).Return(dummyBill, nil)
	BillController := NewBillController(suite.bum, suite.rg, suite.amm)
	BillController.Route()
	record := httptest.NewRecorder()

	// simulasi mengirim sebuah payload dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	// simulasi membuat request ke path `/api/v1/bills`
	// authorization
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", dummyBill.User.Id)
	BillController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *BillControllerTestSuite) TestCreateBill_BindingFailed() {
	// persiapan
	billController := NewBillController(suite.bum, suite.rg, suite.amm)
	billController.Route()

	// ekspektasi
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/bills", nil)
	assert.NoError(suite.T(), err)

	// eksekusi
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("user", dummyBill.User.Id)
	billController.createHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}
