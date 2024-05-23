package application

import (
	"net/http"
)

//func (app *Application) middlewareAuth(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//		if req.Method == http.MethodOptions {
//			next.ServeHTTP(rw, req)
//			return
//		}
//
//		authToken := req.Header.Get("authorization")
//		if authToken == "" {
//			authToken = req.URL.Query().Get("token")
//		}
//		if authToken == "" {
//			rw.Header().Set("content-type", "application/json")
//			bff.Forbidden(rw)
//			return
//		}
//
//		u, errU := app.store.GetAuthUserByToken(req.Context(), authToken)
//		if errU != nil {
//			if !errors.Is(errU, pgx.ErrNoRows) {
//				logger.Error("error get user by token", zap.Error(errU))
//			}
//			rw.Header().Set("content-type", "application/json")
//			bff.ErrorResponse(rw, http.StatusUnauthorized, "unauthorized")
//			return
//		}
//
//		ctx := req.Context()
//		ctx = context.WithValue(ctx, util.ContextKeyUser, u)
//		req = req.WithContext(ctx)
//
//		next.ServeHTTP(rw, req)
//	})
//}

func (app *Application) middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Headers", "*")
		rw.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		rw.Header().Add("Access-Control-Allow-Headers", "Authorization")

		if req.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(rw, req)
	})
}

func (app *Application) middlewareResponseContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, req)
	})
}
