package app

import "to-do-app/internal/pkg/datasource"

type Container struct {
}

func NewContainer(datasource *datasource.DataSource) *Container {
	return &Container{}
}
