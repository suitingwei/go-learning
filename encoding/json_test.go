package encoding

import (
	"encoding/json"
	"strings"
	"testing"
)

var data = `{\"ret_code\":0,\"ret_msg\":\"OK\",\"city_infos\":[{\"cityid\":55000199,\"city_desc\":\"São Paulo\",\"countyid\":55002300,\"county_desc\":\"São Paulo\",\"countryid\":55,\"country_desc\":\"Brazil\",\"provinceid\":35,\"province_desc\":\"São Paulo\",\"country_code\":\"BRA\",\"district_code\":\"01199\",\"time_zone\":{\"utc_offset\":null,\"china_app_offset\":null,\"china_rt_offset\":-660,\"location_utc_offset\":-180,\"os_utc_offset\":null},\"canonical_country_code\":\"BR\"}]}`

func TestName(t *testing.T) {
	result := make(map[string]interface{})

	data = strings.ReplaceAll(data, "\\", "")

	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		t.Fatalf("failed to decode original data,err:%v", err)
		return
	}
	t.Logf("decode result: %v", result)
}
