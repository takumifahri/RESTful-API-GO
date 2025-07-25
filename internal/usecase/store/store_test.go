package store

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/takumifahri/RESTful-API-GO/internal/mocks"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
	"github.com/takumifahri/RESTful-API-GO/internal/repository/order"
)

func Test_storeUsecase_GetAllCatalogList(t *testing.T) {
	type fields struct {
		menuRepo catalog.Repository
		orderRepo order.Repository
	}
	type args struct {
		ctx  context.Context
		tipe string
	}
	tests := []struct {
		name    string
		fields  fields
		s       *storeUsecase
		args    args
		want    []models.ProductClothes
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetAllCatalogList",
			fields: fields{
				menuRepo: func() catalog.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockCatalogRepository(ctrl)

					mock.EXPECT().GetAllCatalogList(gomock.Any(), string(constant.ACCESSORIES)).
						Times(1).
						Return([]models.ProductClothes{
							{
								UNIQUEID: "PRD-123",
								NamaPakaian:     "Test Product",
								Price:    1000,
								Deskripsi: "Test Description",
								Stock: 10,
								TypeClothes: constant.ACCESSORIES,
							},
						}, nil)
					return mock
				}(),
			},
			args: args{
				ctx:  context.Background(),
				tipe: string(constant.ACCESSORIES),
			},
			want: []models.ProductClothes{
				{
					UNIQUEID: "PRD-123",
					NamaPakaian:     "Test Product",
					Price:    1000,
					Deskripsi: "Test Description",
					Stock: 10,
					TypeClothes: constant.ACCESSORIES,
				},
			},
			wantErr: false,
		},
		{
			name: "Faiul Test GetAllCatalogList",
			fields: fields{
				menuRepo: func() catalog.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockCatalogRepository(ctrl)

					mock.EXPECT().GetAllCatalogList(gomock.Any(), string(constant.ACCESSORIES)).
						Times(1).
						Return(nil, errors.New("failed to get catalog list"))
					return mock
				}(),
			},
			args: args{
				ctx:  context.Background(),
				tipe: string(constant.ACCESSORIES),
			},
			want:nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storeUsecase{
				menuRepo: tt.fields.menuRepo,
				orderRepo: tt.fields.orderRepo,
			}
			got, err := s.GetAllCatalogList(tt.args.ctx, tt.args.tipe)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeUsecase.GetAllCatalogList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeUsecase.GetAllCatalogList() = %v, want %v", got, tt.want)
			}
		})
	}
}
