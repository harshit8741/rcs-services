package controllers

import (
	"net/http"
)

const (
	rcsApi = "https://rcsapi-uat.jiocx.com/api/v1/sendMessage"
)

func SimpleTextMessage(w http.ResponseWriter, r *http.Request) {
	var payload SimpleTextMessagePayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, payload)
}

func SimpleImgWithLinkMessage(w http.ResponseWriter, r *http.Request) {
	var payload SimpleImgMessagePayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, payload)
}

func SuggestedReplyMessage(w http.ResponseWriter, r *http.Request) {
	var payload SuggestedReplyPayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, payload)
}

func SuggestedCalenderEventMessage(w http.ResponseWriter, r *http.Request) {
	var payload SuggestedCalenderPayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, payload)
}

func SuggestedLocationMessage(w http.ResponseWriter, r *http.Request) {
	var payload SuggestedLocationPayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, payload)
}

func SuggestedDailerMessage(w http.ResponseWriter, r *http.Request) {
	var payload SuggestedDailerPayload
	if !decodeJSONPayload(r, &payload, w) {
		return
	}
	sendRequestToExternalAPI(w, rcsApi, &payload)
}
