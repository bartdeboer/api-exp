package renderer

import (
	"net/http"
)

type FormRenderer interface {
	Load(schemaFile string) error
	AfterLoad(schemaFile string)
	Output(r *http.Request, w http.ResponseWriter)
}
