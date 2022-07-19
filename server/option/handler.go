package option

import (
	"net/http"

	"github.com/xulei131401/gox/responsex"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/handler"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// NotFoundHandler 路由不存在, 404
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.Error(w, responsex.NewDefaultCodeError("路由不存在"))
		return
	}
}

// NotAllowedHandler 不允许访问
func NotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.Error(w, responsex.NewDefaultCodeError("不允许访问"))
		return
	}
}

// UnauthorizedCallback JWT授权不通过,http code 401
func UnauthorizedCallback() handler.UnauthorizedCallback {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.Error(w, responsex.NewDefaultCodeError("JWT授权不通过"))
		return
	}
}

// SessionHandler Session全局中间件
func SessionHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("全局中间件")
		//servererCtx.SessionStore, _ = serverCtx.SessionManager.SessionStart(w, r)
		//defer serverCtx.SessionStore.SessionRelease(context.Background(), w)
		next(w, r)
	}
}

func ExampleHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "second")
		logx.Info("全局中间件")
		next(w, r)
	}
}

// ErrorHandler 全局错误处理器,统一拦截代码中产生的error,最终返回json
func ErrorHandler() {
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		//logx.Infof("err:%#v", err)
		//logx.Info(string(debug.Stack()))

		switch e := err.(type) {
		case *responsex.CodeError:
			return http.StatusOK, e.Data()
		case error:
			return http.StatusOK, responsex.NewCodeErrorResponse(responsex.FailCode, e.Error())
		default:
			return http.StatusInternalServerError, nil
		}
	})
}
