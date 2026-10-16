package repositories

import (
    "context"

    "konsin1988/go-server/models"
)

type DataProvider interface {
    GetTableData(ctx context.Context) (models.TableData, error)
}
