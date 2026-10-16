package handlers

//import (
//    _ "encoding/json"
//    "net/http"
//
//    "konsin1988/go-server/repositories"
//)


//Health
//func (h *AuditHandler) Health(w http.ResponseWriter, r *http.Request) {
//    if err := h.Repo.Ping(); err != nil {
//	http.Error(w, "database not reachable", http.StatusServiceUnavailable)
//	return 
//    }
//    w.Header().Set("Content-Type", "application/json")
//    w.WriteHeader(http.StatusOK)
//    w.Write([]byte(`{"status": "ok"}`))
//}
