package admin

import (
	"reflect"

	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
)

// SelectManyConfig meta configuration used for select many
type SelectManyConfig struct {
	Collection interface{}
	SelectOneConfig
}

// ConfigureQorMeta configure select many meta
func (selectManyConfig *SelectManyConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*Meta); ok {
		selectManyConfig.SelectOneConfig.Collection = selectManyConfig.Collection
		selectManyConfig.SelectOneConfig.ConfigureQorMeta(meta)
		meta.Type = "select_many"

		// Set FormattedValuer
		if meta.FormattedValuer == nil {
			meta.SetFormattedValuer(func(record interface{}, context *qor.Context) interface{} {
				reflectValues := reflect.Indirect(reflect.ValueOf(meta.GetValuer()(record, context)))
				var results []string
				for i := 0; i < reflectValues.Len(); i++ {
					results = append(results, utils.Stringify(reflectValues.Index(i).Interface()))
				}
				return results
			})
		}
	}
}
