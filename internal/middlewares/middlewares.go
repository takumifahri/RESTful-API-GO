package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/models/constant"
	"github.com/takumifahri/RESTful-API-GO/internal/usecase/auth"
	"github.com/takumifahri/RESTful-API-GO/internal/utils"
	"github.com/takumifahri/RESTful-API-GO/internal/tracing"
	"golang.org/x/net/context"
)


type AuthMiddleware struct {
	authUsecase auth.Usecase
}

func GetAuthMiddleware(authUsecase auth.Usecase) *AuthMiddleware {
	return &AuthMiddleware{
		authUsecase: authUsecase,
	} 
}

func (am *AuthMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	ctx, span := tracing.CreateSpanWrapper(context.Background(), "CheckAuth")
	defer span.End() // Pastikan span diakhiri
	return func(c echo.Context) error {
		sessionData, err := utils.GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized: " + err.Error(), // ini untuk client
				Internal: err, // ini untuk dev
			} 
		}

		userUniqueID, err := am.authUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized: " + err.Error(), // ini untuk client
				Internal: err, // ini untuk dev
			}
		}
		
		authContext := context.WithValue(c.Request().Context(), constant.AuthcontextKey, userUniqueID)
		c.SetRequest(c.Request().WithContext(authContext))

		if err := next(c); err != nil {
			return err
		}
		return nil
	}


}
