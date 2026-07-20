package main

import (
    "encoding/json"
    "net/http"
   	"strconv"
)

type Menu struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Type  string  `json:"type"`
}

var menus = []Menu{
    {ID: 1, Name: "ต้มยำกุ้ง", Price: 120, Type: "soup"},
    {ID: 2, Name: "พิซซ่าฮาวายเอี้ยน", Price: 199, Type: "pizza"},
    {ID: 3, Name: "พิซซ่าเห็ด", Price: 179, Type: "pizza"},
}

var nextID = 4

func listMenu(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(menus)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /menu", listMenu)
    http.ListenAndServe(":8080", mux)
	mux.HandleFunc("GET /menu/{id}", getMenu)
}

func getMenu(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "เลขจานต้องเป็นตัวเลข", http.StatusBadRequest)
        return
    }
    for _, m := range menus {
        if m.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(m)
            return
        }
    }
    http.Error(w, "ไม่พบเมนูหมายเลขนี้", http.StatusNotFound)
}