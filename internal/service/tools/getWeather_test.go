package tools

import "testing"

func TestGetWeather(t *testing.T) {
	res, err := GetWeather("Murmansk", "ru")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
