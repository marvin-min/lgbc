package main

import (
	"encoding/json"
	"fmt"
)

type DayLive struct {
	TotalCount int
	LogDay string
}

type Rs struct {
	DayLives []DayLive
}

func main() {
	var rs Rs
	str := `{"DayLives":[{"TotalCount":189393,"LogDay":"2016-10-21"},{"TotalCount":222218,"LogDay":"2016-10-22"},{"TotalCount":225080,"LogDay":"2016-10-23"},{"TotalCount":213959,"LogDay":"2016-10-24"},{"TotalCount":220441,"LogDay":"2016-10-25"},{"TotalCount":235729,"LogDay":"2016-10-26"},{"TotalCount":246066,"LogDay":"2016-10-27"},{"TotalCount":178436,"LogDay":"2016-10-28"}]}`
	err := json.Unmarshal([]byte(str), &rs)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v",len(rs.DayLives))
}

