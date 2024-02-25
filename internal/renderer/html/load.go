package htmlrenderer

import "github.com/bartdeboer/api-exp/internal/form"

func (f *Form) Load(schemaFile string) error {
	return form.Load(f, schemaFile)
}

func (f *Form) AfterLoad(schemaFile string) {
	f.SchemaFile = schemaFile
}
