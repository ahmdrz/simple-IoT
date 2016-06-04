package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	output := JSON{ErrorCode: 0, Success: true, Result: nil}
	token := r.Header.Get("Authorization")
	if !isUserExists(token) {
		output.ErrorCode = 1
		output.Success = false
	}

	json.NewEncoder(w).Encode(output)
}

func LightsShow(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	output := JSON{}
	if isUserExists(token) {
		vars := mux.Vars(r)
		lightIDString := vars["lightID"]
		lightID, err := strconv.Atoi(lightIDString)
		if err != nil {
			output.ErrorCode = 3
			output.Success = false
			output.Result = nil
		} else {
			light := getLight(token, lightID)
			if light.HostID > 0 {
				output.Result = light
				output.Success = true
				output.ErrorCode = 0
			} else {
				output.Result = nil
				output.Success = false
				output.ErrorCode = 2
			}
		}
	} else {
		output.Result = nil
		output.Success = false
		output.ErrorCode = 1
	}
	json.NewEncoder(w).Encode(output)
}

func LightsList(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	output := JSON{ErrorCode: 1, Success: false, Result: nil}
	if isUserExists(token) {
		lights := getList(token)
		output.Result = lights
		output.ErrorCode = 0
		output.Success = true
	}

	json.NewEncoder(w).Encode(output)
}

func LightsSet(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	output := JSON{Result: nil, Success: false, ErrorCode: 1}
	if isUserExists(token) {
		if err != nil {
			output.ErrorCode = 3
		} else {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				output.ErrorCode = 4
				output.Success = false
				output.Result = nil
			}

			var j Light
			json.Unmarshal(b, &j)

			if isIDExists(j.ID) && getHostID(token) == j.HostID {
				output.ErrorCode = 0
				output.Success = true
				output.Result = j
				updateLight(token, j)
			} else {
				output.ErrorCode = 5
				output.Success = false
				output.Result = nil
			}
		}
	}
	json.NewEncoder(w).Encode(output)
}
