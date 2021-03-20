package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/rusldv/shern/lib"
)

// Определение флагов
var config = flag.String("config", "./config.json", "Initial configuration file.")
var version = flag.Bool("v", false, "Version for Shern server.")

func main() {
	flag.Parse()
	// Если указан в параметрах флаг -v, то просто выводим версию и выходим из программы
	if *version {
		fmt.Println(lib.ServerVersion)
		return
	}
	// Выводим информацию о сервере
	log.Printf("%s v%s\n", lib.ServerName, lib.ServerVersion)
	// Парсим конфиг и инициализируем структуры данных
	log.Println("Initializing Shern server...")
	cfg, err := lib.ParseConfig(*config)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Обработка запроса
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shr := lib.NewShernRequest(cfg, r)
		fmt.Println(shr)
		if shr.IsError {
			fmt.Println("Есть ошибка, код", shr.Code)
			switch shr.Code {
			case 404:
				fmt.Println("Ошибка 404")
				http.NotFound(w, r)
			default:
				fmt.Println("Другая ошибка")
				http.Error(w, shr.Msg, shr.Code)
			}
			return
		} 
		fmt.Println("Ошибок нет, код", shr.Code)
		//fmt.Println(shr.Conf)
		/*
			StartAccept(w, ShernContext) (map[string][]string, error)
			Внутри мы подключаем функции работы с ctx, w, r которые извлекают и записывают данные
			они будут в lib/interpfuncs.go
		*/
	})
	// Вывод сообщения о запуске сервера
	log.Println("Server listening...")
	// Запуск слушателя порта, который вводит главный поток программы в бесконечный цикл ожидания запросов
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
