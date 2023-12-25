package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type HotDryNoodles struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Brand *Brand `json:"brand"`
}

var hotDryNoodles []HotDryNoodles

func fakeData() {
	// 假数据制造

	// 标准热干面 7 蔡林记
	// 麻辣热干面 9 蔡林记
	// 螺蛳粉干面 10 螺鼎记
	clj := &Brand{ID: 1, Name: "蔡林记"}
	ldj := &Brand{ID: 3, Name: "螺鼎记"}
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{ID: 1, Name: "标准热干面", Price: 7, Brand: clj})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{ID: 1, Name: "麻辣热干面", Price: 9, Brand: clj})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{ID: 1, Name: "螺蛳粉干面", Price: 10, Brand: ldj})
}

func main() {
	fakeData()

	r := mux.NewRouter()

	r.HandleFunc("/api/noodles", getAllNoodles).Methods("GET")
	r.HandleFunc("/api/noodles/{id}", getNoodles).Methods("GET")
	r.HandleFunc("/api/noodles", createNoodles).Methods("POST")
	r.HandleFunc("/api/Noodles/{id}", updateNoodles).Methods("PUT")
	r.HandleFunc("/api/Noodles/{id}", deleteNoodles).Methods("DELETE")

	log.Println("Server running on http://127.0.0.1:8080/ .")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func deleteNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application")
	param := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, err := strconv.Atoi(param["id"])
		if err != nil {
			writer.WriteHeader(400)
			_ = json.NewEncoder(writer).Encode(map[string]interface{}{
				"status": -1,
				"msg":    "func deleteNoodles 删除失败 请输入正确的 ID",
			})
			return
		}
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)
			_ = json.NewEncoder(writer).Encode(hotDryNoodles)
			return
		}
	}
}

func updateNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	param := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, err := strconv.Atoi(param["id"])
		if err != nil {
			writer.WriteHeader(400)
			_ = json.NewEncoder(writer).Encode(map[string]interface{}{
				"status": -1,
				"msg":    "func updateNoodles 输入的ID错误或不存在 请重试",
			})
			return
		}
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)
			var noodle HotDryNoodles
			err = json.NewDecoder(request.Body).Decode(&noodle)
			if err != nil {
				writer.WriteHeader(400)
				_ = json.NewEncoder(writer).Encode(map[string]interface{}{
					"status": -1,
					"msg":    "func updateNoodles 传入的数据有错 请重试",
				})
				return
			}
			_ = json.NewEncoder(writer).Encode(noodle)
			return
		}
	}
}

func createNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var noodle HotDryNoodles
	err := json.NewDecoder(request.Body).Decode(&noodle)
	if err != nil {
		writer.WriteHeader(400)
		_ = json.NewEncoder(writer).Encode(map[string]interface{}{
			"status": -1,
			"msg":    "POST的数据不对 请重试",
		})
		return
	}
	hotDryNoodles = append(hotDryNoodles, noodle)
	_ = json.NewEncoder(writer).Encode(noodle)
}

func getNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range hotDryNoodles {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			_ = json.NewEncoder(writer).Encode(map[string]interface{}{
				"status": -1,
				"msg":    "func getNoodles 输入的ID错误或不存在 请重试",
			})
			return
		}
		if item.ID == id {
			err = json.NewEncoder(writer).Encode(item)
			if err != nil {
				writer.WriteHeader(400)
				_ = json.NewEncoder(writer).Encode(map[string]interface{}{
					"status": -1,
					"msg":    "func getNoodles 服务器内部错误 解析 item 出错",
				})
			}
			return

		}
	}
}

func getAllNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(hotDryNoodles)
	if err != nil {
		writer.WriteHeader(400)
		_ = json.NewEncoder(writer).Encode(map[string]interface{}{
			"status": -1,
			"msg":    "func getAllNoodles 服务器内部错误 解析hotDryNoodles出错",
		})
	}
}
