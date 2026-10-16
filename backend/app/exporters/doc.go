package exporters

import (
    _ "fmt"
    _ "io/ioutil"
    _ "os"
    _ "path/filepath"
    _ "strings"
    _ "time"

    "konsin1988/go-server/models"
)

type DocExporter struct{}

func (e *DocExporter) Export(data models.TableData) ([]byte, string, error){
    bytes := []byte{}
    return bytes, "application/vnd.openxmlformats-officedocument.wordprocessingml.document", nil
}
