package interview

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type chs struct {
	key   string
	value string
	trait string
}

func InterviewTask(ctx *gin.Context, ch chan map[string]interface{}) {
	ls := make([]chs, 10)
	ls2 := make([]chs, 10)
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
	result := map[string]interface{}{
		"attributes": "",
		"traits":     "",
	}
	var i int
	for k, v := range mp {
		if len(k) >= 4 && k[:4] == "atrk" {
			i, _ = strconv.Atoi(k[4:])
			ls[i].key = v
			continue
		}
		if len(k) >= 4 && k[:4] == "atrv" {
			i, _ = strconv.Atoi(k[4:])
			ls[i].value = v
			continue
		}
		if len(k) >= 4 && k[:4] == "atrt" {
			i, _ = strconv.Atoi(k[4:])
			ls[i].trait = v
			continue
		}
		if len(k) >= 5 && k[:5] == "uatrk" {
			i, _ = strconv.Atoi(k[5:])
			ls2[i].key = v
			continue
		}
		if len(k) >= 5 && k[:5] == "uatrv" {
			i, _ = strconv.Atoi(k[5:])
			ls2[i].value = v
			continue

		}
		if len(k) >= 5 && k[:5] == "uatrt" {
			i, _ = strconv.Atoi(k[5:])
			ls2[i].trait = v
			continue
		}
		result[m[k]] = v

	}

	var l []map[string]interface{}

	for _, k := range ls {
		if k.value == "" {
			continue
		}
		l = append(l, map[string]interface{}{
			k.key: map[string]string{
				"value": k.value,
				"type":  k.trait,
			},
		})
	}

	var l2 []map[string]interface{}

	for _, k := range ls2 {
		if k.value == "" {
			continue
		}
		l2 = append(l2, map[string]interface{}{
			k.key: map[string]string{
				"value": k.value,
				"type":  k.trait,
			},
		})
	}
	result["attributes"] = l
	result["traits"] = l2
	ch <- result
}
