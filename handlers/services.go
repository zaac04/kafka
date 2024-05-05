package handlers

import (
	"anya/db/redis"
	"anya/enums"
	"anya/handlers/helpers"
	"anya/structs"
	"anya/utilities"
	"encoding/json"
	"net/http"
)

func ServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Pong!"))
}

func RegisterService(w http.ResponseWriter, r *http.Request) {

	var Service structs.Service
	svc_hash := utilities.GetServiceDetailsKey(Service.Name)
	utilities.UnmarshalReqBody(r.Body, &Service)
	err := redis.Redis.SetJSON(&structs.SetCache{
		Ctx:    r.Context(),
		Key:    svc_hash,
		Val:    Service,
		Expiry: enums.NeverExpire,
	})

	if err != nil {
		utilities.LogError(err, enums.RedisJsonSetFailed)
		response, err := json.Marshal(&structs.Error{
			Error: err.Error(),
		})

		if err != nil {
			utilities.LogError(err, "JSON Marshalling Error")
		}

		helpers.SendError(w, response)
		return
	}

	err = redis.Redis.SetList(&structs.SetCache{
		Ctx:    r.Context(),
		Key:    utilities.GetServicesKey(),
		Val:    svc_hash,
		Expiry: enums.NeverExpire,
	})

	if err != nil {
		utilities.LogError(err, enums.RedisListSetFailed)
	}

}

func UpdateHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Pong!"))
}

func SharedSVConfig(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Pong!"))
}
