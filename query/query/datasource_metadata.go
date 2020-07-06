package query

import (
	"github.com/grafadruid/go-druid/query"
	"github.com/grafadruid/go-druid/query/types"
)

type DataSourceMetadata struct {
	*Base
}

func NewDataSourceMetadata() *DataSourceMetadata {
	d := &DataSourceMetadata{}
	d.SetQueryType("dataSourceMetadata")
	return d
}

func (d *DataSourceMetadata) SetDataSource(dataSource query.DataSource) *DataSourceMetadata {
	d.Base.SetDataSource(dataSource)
	return d
}

func (d *DataSourceMetadata) SetIntervals(intervals []types.Interval) *DataSourceMetadata {
	d.Base.SetIntervals(intervals)
	return d
}

func (d *DataSourceMetadata) SetContext(context map[string]interface{}) *DataSourceMetadata {
	d.Base.SetContext(context)
	return d
}
