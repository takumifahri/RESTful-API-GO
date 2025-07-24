package store

import (
	"context"
	"reflect"
	"testing"

	"github.com/takumifahri/RESTful-API-GO/internal/models"
)

func Test_storeUsecase_GetAllCatalogList(t *testing.T) {
	type args struct {
		ctx  context.Context
		tipe string
	}
	tests := []struct {
		name    string
		s       *storeUsecase
		args    args
		want    []models.ProductClothes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllCatalogList(tt.args.ctx, tt.args.tipe)
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
