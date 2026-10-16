package services

import (
    "context"
    "errors"

    "konsin1988/go-server/models"
    "konsin1988/go-server/repositories"
    "konsin1988/go-server/exporters"
)

type ExportService struct {
    Providers map[string]repositories.DataProvider
    Exporters map[string]exporters.Exporter
}

func (s *ExportService) GetTableData(
    ctx context.Context,
    db, entity string,
) (models.TableData, error) {
    
    key := db + "/" + entity

    provider, ok := s.Providers[key]
    if !ok {
        return models.TableData{}, errors.New("unknown db/entity")
    }

    return provider.GetTableData(ctx)
}

func (s *ExportService) Export(
    format string,
    data models.TableData,
) ([]byte, string, error) {

    exporter, ok := s.Exporters[format]
    if !ok {
        return nil, "", errors.New("unknown format")
    }

    return exporter.Export(data)
}

