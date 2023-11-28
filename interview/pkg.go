package interview

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type chs struct {
	key   string
	value string
	trait string
}

func InterviewTask(ctx *gin.Context, ch chan map[string]interface{}, webhookURL string) {
	attributes := make([]chs, 10)
	traits := make([]chs, 10)
	m := map[string]string{
		"ev":  "event",
		"et":  "event_type",
		"id":  "app_id",
		"uid": "user_id",
		"mid": "message_id",
		"t":   "page_title",
		"p":   "page_url",
		"l":   "browser_language",
		"sc":  "screen_size",
	}

	var mp map[string]string
	err := ctx.BindJSON(&mp)
	if err != nil {
		log.Println("unable to marshal")
		return
	}
	result := map[string]interface{}{}

	for k, v := range mp {
		if len(k) >= 4 && k[:4] == "atrk" {
			i, _ := strconv.Atoi(k[4:])
			attributes[i].key = v
			continue
		}
		if len(k) >= 4 && k[:4] == "atrv" {
			i, _ := strconv.Atoi(k[4:])
			attributes[i].value = v
			continue
		}
		if len(k) >= 4 && k[:4] == "atrt" {
			i, _ := strconv.Atoi(k[4:])
			attributes[i].trait = v
			continue
		}
		if len(k) >= 5 && k[:5] == "uatrk" {
			i, _ := strconv.Atoi(k[5:])
			traits[i].key = v
			continue
		}
		if len(k) >= 5 && k[:5] == "uatrv" {
			i, _ := strconv.Atoi(k[5:])
			traits[i].value = v
			continue

		}
		if len(k) >= 5 && k[:5] == "uatrt" {
			i, _ := strconv.Atoi(k[5:])
			traits[i].trait = v
			continue
		}
		result[m[k]] = v
	}

	// Populate attributes in the desired format
	attributeMap := make(map[string]map[string]interface{})
	for _, k := range attributes {
		if k.value == "" {
			continue
		}
		attributeMap[k.key] = map[string]interface{}{
			"value": k.value,
			"type":  k.trait,
		}
	}

	// Populate traits in the desired format
	traitMap := make(map[string]map[string]interface{})
	for _, k := range traits {
		if k.value == "" {
			continue
		}
		traitMap[k.key] = map[string]interface{}{
			"value": k.value,
			"type":  k.trait,
		}
	}

	result["attributes"] = attributeMap
	result["traits"] = traitMap

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Println("Error marshaling result to JSON:", err)
		return
	}

	// Make an HTTP POST request to the webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error making POST request to webhook:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Webhook response Status:", resp.Status)
	ch <- result
}
