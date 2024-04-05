package main

import (
	reqYandex "BriefRetelling/internal/request"
	"BriefRetelling/internal/scrapHTML"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type request struct {
	Url string `json:"url"`
}

type response struct {
	Header string   `json:"header"`
	Info   []string `json:"info"`
}

func main() {
	//cfg := config.MustLoad()
	//connPQ := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode)
	//
	//storage, err := pq.New(connPQ)
	//if err != nil {
	//	slog.Error("failed to create storage:", err)
	//	os.Exit(1)
	//}
	pathJsonRes := "D:/University/Технологии программирования//Курсовая/SearchForInformation/response.json"
	pathJsonReq := "D:/University/Технологии программирования//Курсовая/SearchForInformation/request.json"
	for {
		ex, _ := exists(pathJsonReq)
		if !ex {
			time.Sleep(time.Millisecond * 200)
			continue
		}
		println(2)
		f, _ := os.Open(pathJsonReq)
		datafileJson, _ := io.ReadAll(f)
		p := request{}
		fmt.Println(string(datafileJson))
		json.Unmarshal(datafileJson, &p)
		f.Close()
		res, _ := reqYandex.RequestMethod(p.Url)
		fmt.Println(res)
		var resq *scrapHTML.Data
		resq = scrapHTML.Scrap(res)
		f, _ = os.Create(pathJsonRes)
		m, _ := json.Marshal(resq)
		f.Write(m)
		os.Remove(pathJsonReq)
		f.Close()
	}
	//for {
	//
	//	res, err := request.RequestMethod(url)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	data := scrapHTML.Scrap(res)
	//	err = storage.NewData(data.Heading, strings.Join(data.TextItem, "\n"), url)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	// TODO: maybe init logger: slog

	// TODO: init handlers for HTTP requests

	// TODO: init Docker
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
