package models 

import (
    _ "time"
)

type TableData struct {
    Title       string
    Headers     []string
    Rows        [][]string
}
