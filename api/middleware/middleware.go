package middleware

type Middleware struct {
	Multipart MultipartMiddleware
}

func NewMiddleware() Middleware {
	return Middleware{
		Multipart: NewMultipartMiddleware(),
	}
}
