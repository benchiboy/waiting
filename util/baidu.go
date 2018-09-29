package util

type BaiduLbs struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Source   string `json:"source"`
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		Bound             string `json:"bound"`
		FormattedAddress  string `json:"formatted_address"`
		AddressComponents struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Street   string `json:"street"`
			Level    string `json:"level"`
		} `json:"address_components"`
		Precise float64 `json:"precise"`
	} `json:"result"`
}

/*
	百度LBS的构造函数
*/
func NewBaiduLbs() *BaiduLbs {
	return &BaiduLbs{}
}

/*
	百度LBS的逆地址编码函数
*/
func (b *BaiduLbs) GetAddr(lng string, lat string) string {

	return ""
}
