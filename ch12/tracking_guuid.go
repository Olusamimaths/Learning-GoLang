package tracker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type guidKey int
const key guidKey = 1

func contextWithGuid(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	guid, ok := ctx.Value(key).(string)
	return guid, ok
}

func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if guid := r.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGuid(ctx, guid)
		} else {
			ctx = contextWithGuid(ctx, uuid.New().String())
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(rw, r)
	})
}
type Logger struct {}

func (Logger) Log(ctx context.Context, message string) {
	if guild, ok := guidFromContext(ctx) ; ok {
		message = fmt.Sprintf("GUID: %s - %s", guild, message)
	}
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guild, ok := guidFromContext(ctx); ok {
		req.Header.Add("X-GUID", guild)
	}
	return req
}

// ClIENT
type CLogger interface {
	Log(context.Context, string)
}
type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger	CLogger
	Remote string
}
func (bl BusinessLogic) businessLogic(ctx context.Context, user string, data string) (string,error) {
	bl.Logger.Log(ctx, "starting business logic for " + user + " with " + data)
	req, err :=http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)

	if err != nil {
		bl.Logger.Log(ctx, "error building remote request: " + err.Error())
		return "", err
	}

	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	return string(resp.ContentLength), err
}

func main() {
	bl := BusinessLogic {
		RequestDecorator: Request,
		Logger: Logger{},
		Remote: "http://www.example.com/query",
	}
}