package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
)

func UnmarshalReqBody(body io.ReadCloser, output interface{}) (err error) {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(output)
	return err
}

func GetServiceDetailsKey(service string) string {
	hash := md5.Sum([]byte(service))
	return hex.EncodeToString(hash[:])
}

func GetServicesKey() string {
	hash := md5.Sum([]byte("All_Services"))
	return hex.EncodeToString(hash[:])
}
