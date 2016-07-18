package admin

import (
	"github.com/qor/qor/resource"
)

type RichEditorConfig struct {
	AssetManager *Resource
}

// ConfigureQorMeta configure rich editor meta
func (richEditorConfig *RichEditorConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*Meta); ok {
		meta.Type = "rich_editor"

		// Compatible with old rich editor setting
		if meta.Resource != nil {
			richEditorConfig.AssetManager = meta.Resource
			meta.Resource = nil
		}
	}
}
