package vatsim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetDatafeed retrieves network data from server
func GetDatafeed() (jsonData APIResponse) {
	res, err := http.Get("https://data.vatsim.net/v3/vatsim-data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	res.Body.Close()

	jsonData = APIResponse{}

	err = json.Unmarshal(body, &jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
