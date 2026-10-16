package main

import (
    "net/http"
    "path/filepath"
    "log"
    "strings"
    "github.com/go-chi/chi/v5"

    "konsin1988/go-server/config"
    "konsin1988/go-server/handlers"
    "konsin1988/go-server/repositories"
    "konsin1988/go-server/exporters"
    "konsin1988/go-server/services"

)

func FileServer(r chi.Router, path string, root http.FileSystem) {
    if strings.ContainsAny(path, "{}*"){
        panic("FileServer does not permit URL parameters")
    }

    fs := http.StripPrefix(path, http.FileServer(root))

    if path != "/" && path[len(path)-1] != '/' {
        r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
        path += "/"
    }

    r.Get(path+"*", fs.ServeHTTP)
}


func main() {

    dbConfigs := config.LoadDBConfigs()
    dbs, err := config.InitDatabases(dbConfigs)
    if err != nil {
        log.Fatal(err)
    }
    defer dbs.Close()

    _ = dbs.AuditCkk
    _ = dbs.AuditInternal
    _ = dbs.Certificate

    providers := map[string]repositories.DataProvider{
        "ckk/checkreestr": repositories.NewCkkCheckReestrRepo(dbs.AuditCkk),
    }

    exporters := map[string]exporters.Exporter{
        "pdf": &exporters.PdfExporter{},
        "xlsx": &exporters.XlsxExporter{},
        "doc": &exporters.DocExporter{},
    }

    exportService := &services.ExportService{
        Providers: providers,
        Exporters: exporters,
    }

    exportHandler := handlers.NewExportHandler(exportService)


    r := chi.NewRouter()

    // HealthCheck
    r.Get("/auditso/api/health/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"stutus":"ok"}`))
    })


    // API
    r.Route("/auditso/api", func(r chi.Router) {
        r.Get("/{db}/{entity}/{format}/", exportHandler.Export)
    })

    // Static assets
    FileServer(r, "/auditso/assets", http.Dir("./ui/assets"))

    // SPA fallback
    r.Get("/auditso/*", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("ui", "index.html"))
    })

    


    //fs := http.FileServer(http.Dir("./ui"))
    //mux.Handle("/auditso/assets/", http.StripPrefix("/auditso/", fs))

    //mux.HandleFunc("/auditso/", func(w http.ResponseWriter, r *http.Request) {
    //    http.ServeFile(w, r, filepath.Join("ui", "index.html"))
    //})

    log.Println("Server started on :8014")
    http.ListenAndServe(":8014", r)

}
