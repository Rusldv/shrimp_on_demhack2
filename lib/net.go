package lib

import (
	"strings"
	"os"
	"net/http"
)

// ShernRequest содержит данные запроса к серверу в том числе объект Context.
type ShernRequest struct {
	Conf          *Config
	Host          string
	ServerRoot    string
	URLPath       string
	ComponentName string
	ComponentPath string
	ComponentExists bool
	IsError bool
	Code int 
	Msg string
	SystemError string
}

// NewShernRequest принимает объект запроса *http.Request и возвращает *ShernRequest
func NewShernRequest(cfg *Config, r *http.Request) *ShernRequest {
	// Определяем путь запроса, если запрошен корень ("/"), то устанавливаем как cfg.IndexPath
	var urlPath string
	if r.URL.Path == "/" {
		urlPath = "/" + cfg.IndexComponent
	} else {
		urlPath = r.URL.Path
	}
	// Название компонента - это первый элемент urlPath (должен быть существующим каталогом)
	items := strings.Split(urlPath, "/")
	// Проверим чтобы это был каталог иначе просто выплюнуть файл
	componentName := items[1]
	componentPath := cfg.RootDir + r.Host + "/" + componentName
	componentExists := false
	isError := false
	code := 0
	msg := ""
	systemError := ""
	st, err := os.Stat(componentPath)
	if err != nil {
		if os.IsNotExist(err) {
			componentExists = false
			isError = true
			systemError = err.Error()
			code = 404
			msg = "Not Found"
		} else {
			componentExists = false
			isError = true
			systemError = err.Error()
			code = 500
			msg = "Error"
		} 
	} else {
		isError = false
		code = 200
		msg = "OK"
		// Если запрошен каталог, предполагается что это компонент
		if st.IsDir() {
			componentExists = true
		} else {
			componentExists = false
		}
	}

	return &ShernRequest{
		Conf:          cfg,
		Host:          r.Host,
		ServerRoot:    cfg.RootDir + r.Host,
		URLPath:       urlPath,
		ComponentName: componentName,
		ComponentPath: componentPath,
		ComponentExists: componentExists,
		IsError: isError,
		Code: code,
		Msg: msg,
		SystemError: systemError,
	}
}
