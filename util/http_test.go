package util

import (
	"testing"
)

func TestHttp_get(t *testing.T) {

	url := "http://api.map.baidu.com/cloudrgc/v1?geotable_id=194731&id=2560371989&ak=QjOpH9XPrX9Ak5qhFQYXNE2hP3KMQPms&"

	Http_get(url, map[string]string{"location": "40.055,114.308"})
}
