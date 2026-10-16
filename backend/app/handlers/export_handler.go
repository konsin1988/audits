package handlers

import (
    "net/http"
    "fmt"

    "github.com/go-chi/chi/v5"
    "konsin1988/go-server/services"
)

type ExportHandler struct {
    Service *services.ExportService
}

func NewExportHandler(s *services.ExportService) *ExportHandler {
    return &ExportHandler{Service: s}
}

func (h *ExportHandler) Export(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    db := chi.URLParam(r, "db")
    entity := chi.URLParam(r, "entity")
    format := chi.URLParam(r, "format")

    data, err := h.Service.GetTableData(ctx, db, entity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    file, contentType, err := h.Service.Export(format, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", contentType)
    w.Header().Set(
        "Content-Disposition",
        fmt.Sprintf(`attachment; filename="%s_%s.%s"`, db, entity, format),
    )
    w.Write(file)
}
