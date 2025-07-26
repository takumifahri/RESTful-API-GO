package store_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/takumifahri/RESTful-API-GO/internal/mocks"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"github.com/takumifahri/RESTful-API-GO/internal/usecase/auth"
	"github.com/takumifahri/RESTful-API-GO/internal/usecase/store"
)

var _ = Describe("GinkgoStore", func() {
	 var usecase store.Usecase
    var authusecase auth.Usecase
    var menuRepoMock *mocks.MockCatalogRepository
    var orderRepoMock *mocks.MockOrderRepository
    var authRepoMock *mocks.MockAuthRepository
    var mockCtrl *gomock.Controller

    BeforeEach(func() {
        mockCtrl = gomock.NewController(GinkgoT())
        menuRepoMock = mocks.NewMockCatalogRepository(mockCtrl)
        orderRepoMock = mocks.NewMockOrderRepository(mockCtrl)
        authRepoMock = mocks.NewMockAuthRepository(mockCtrl)
        authusecase = auth.GetUsecase(authRepoMock)
        usecase = store.GetUsecase(menuRepoMock, orderRepoMock)
    })

    AfterEach(func() {
        mockCtrl.Finish()
    })

   // ada 4 cara untuk membuat test di Ginkgo
	Describe("request catalog list", func() {
		Context("It giving a valid request", func() {
			inputs := models.GetOrderInfoRequest{
				UserUniqueID: "USR-123",
				OrderID:      "order123",
			}

			When("the order is not for the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetInfoOrder(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(models.Order{
							ID: 1,
							UNIQUEID: "ORD-123",
							UserUniqueID: "ORD-999", // Different from inputs.UserUniqueID to simulate unauthorized
							Status: constant.OrderStatusCompleted,
							ProductOrder: []models.ProductOrder{},
							ReferenceID: "REF-123",
						}, nil)
				})
				It("returns an unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("order with ID order123 does not belong to user ORD-123"))
					Expect(res).To(BeEquivalentTo(models.Order{}))
				})
			})

			When("the order is for the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetInfoOrder(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(models.Order{
							ID: 1,
							UNIQUEID: "USR-123",
							UserUniqueID: "ORD-123",
							Status: constant.OrderStatusCompleted,
                            ProductOrder: []models.ProductOrder{},
							ReferenceID: "REF-123",
						}, nil)
				})
				It("returns the correct order info", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).To(BeNil())
					Expect(res).To(BeEquivalentTo(models.Order{
						ID: 1,
						UNIQUEID: "ORD-123",
						UserUniqueID: "USR-456",
						Status: constant.OrderStatusCompleted,
						ProductOrder: []models.ProductOrder{},
						ReferenceID: "REF-123",
					}))
				})
			})
		})
	})

    Describe("auth usecase", func() {
        Context("CheckSession with valid session", func() {
            It("returns user unique id", func() {
                session := models.UserSession{JWTToken: "token123"}
				authRepoMock.EXPECT().
                    CheckSession(gomock.Any(), session).
                    Return("USR-123", nil)

                userID, err := authusecase.CheckSession(context.Background(), session)
                Expect(err).To(BeNil())
                Expect(userID).To(Equal("USR-123"))
            })
        })

        Context("CheckSession with invalid session", func() {
            It("returns error", func() {
                session := models.UserSession{JWTToken: "invalidtoken"}
                authRepoMock.EXPECT().
                    CheckSession(gomock.Any(), session).
                    Return("", errors.New("invalid session"))

                userID, err := authusecase.CheckSession(context.Background(), session)
                Expect(err).To(HaveOccurred())
                Expect(userID).To(BeEmpty())
            })
        })
    })
})
