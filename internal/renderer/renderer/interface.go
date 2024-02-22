package renderer

import (
	"net/http"
)

type FormRenderer interface {
	Load(schemaFile string) error
	Output(r *http.Request, w http.ResponseWriter)
}
