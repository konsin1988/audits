package exporters

import  "konsin1988/go-server/models"

type Exporter interface {
    Export(models.TableData) ([]byte, string, error)
}
