package catalog

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/takumifahri/RESTful-API-GO/internal/models"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_menuRepo_GetAllCatalogList(t *testing.T) {

	type args struct {
		ctx  context.Context
		tipe string
	}
	tests := []struct {
		name    string
		m       *menuRepo
		args    args
		want    []models.ProductClothes
		wantErr bool
		initMock func() (*sql.DB, sqlmock.Sqlmock, error)
	}{
		// TODO: Add test cases.
		{
			name: "Test GetAllCatalogList with valid type",
			args: args{
				ctx:  context.Background(),
				tipe: "",
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "product_clothes"`),
				).WillReturnRows(sqlmock.NewRows([]string{
					"UNIQUEID", 
					"NamaPakaian",
					"Price",
					"Deskripsi",
					"Stock",
					"TypeClothes",
				}).AddRow("PRD-123", "Test Product", 1000, "Test Description", 10, constant.ACCESSORIES))
			
				return db, mock, err
			},
			want: []models.ProductClothes{
				{
					UNIQUEID:     "PRD-123",
					NamaPakaian:  "Test Product",
					Price:        1000,
					Deskripsi:    "Test Description",
					Stock:        10,
					TypeClothes:  constant.ACCESSORIES,
				},
			},

		},
		{
			name: "Fail Test GetAllCatalogList",
			args: args{
				ctx:  context.Background(),
				tipe: "",
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "product_clothes"`),
				).WillReturnError(errors.New("failed mock to get catalog list"))
			
				return db, mock, err
			},
			want: nil,
			wantErr: true,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, err := tt.initMock()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				DSN: "sqlock_db_0",
				DriverName: "postgres",
				Conn: db,
				PreferSimpleProtocol: true,
			}))
			if err != nil {
				t.Error(err)
			}
			m := &menuRepo{
				db: gormDB,
			}
			got, err := m.GetAllCatalogList(tt.args.ctx, tt.args.tipe)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuRepo.GetAllCatalogList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuRepo.GetAllCatalogList() = %v, want %v", got, tt.want)
			}
			if err := dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("dbMock.ExpectationsWereMet() = %s", err)
			}
		})
	}
}
