package request

import (
	"go/validation-api/pkg/response"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	// выносим Decode и IsValid отдельно чтобы не замешивать работу с http с декодированием и работой со сторонней библиотекой
	// например, чтобы заменить библиотеку валидации - будем делать это в validate.go не трогая тут логику
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(w, map[string]string{"error": err.Error()}, 400)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.Json(w, map[string]string{"error": err.Error()}, 400)
		return nil, err
	}

	return &body, nil
}
