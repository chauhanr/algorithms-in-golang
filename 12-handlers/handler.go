package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func RoutePatternAHandler(res http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("token")
	result := ""

	if token == "" {
		err := errors.New("no token present cannot process further")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	} else {
		err := processToken(token)
		if err != nil {
			result = "error processing request :" + err.Error()
			res.WriteHeader(http.StatusInternalServerError)
		} else {
			result = "processing successful - for token : " + token
			res.WriteHeader(http.StatusOK)
		}
	}
	res.Write([]byte(result))
	return
}

func RoutePatternBHandler(res http.ResponseWriter, req *http.Request) {
	var (
		code   int
		err    error
		result string
	)
	token := req.Header.Get("token")
	code = 200

	if token == "" {
		log.Println("no token so should return 500")
		err = errors.New("no token present cannot process further")
		code = http.StatusInternalServerError
	} else {
		err = processToken(token)
		if err != nil {
			result = "error processing request :" + err.Error()
			code = http.StatusInternalServerError
		} else {
			result = "processing successful - for token : " + token
			code = http.StatusOK
		}
	}
	if err == nil {
		res.Write([]byte(result))
	} else {
		res.WriteHeader(code)
		data := ErrorInfoFunc(err, code)
		err = json.NewEncoder(res).Encode(data)
		if err != nil {
			log.Printf("error encoding response error data")
		}
	}
	return
}

func processToken(token string) error {
	if token == "cerror" {
		return errors.New("code gen error occurred")
	} else if token == "werror" {
		return errors.New("write error occured")
	} else if token == "derror" {
		return errors.New("directory error occurred")
	} else {
		return nil
	}
}

type ErrorInfo struct {
	Message string
	Code    int
}

func ErrorInfoFunc(err error, code int) *ErrorInfo {
	i := ErrorInfo{}
	i.Message = err.Error()
	i.Code = code
	return &i
}
