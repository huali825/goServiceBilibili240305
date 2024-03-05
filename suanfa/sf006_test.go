package suanfa

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func TestSuanfa006(t *testing.T) {
	movie1 := Movie{
		"喜剧之王", 1999, 20, []string{"zhouxingchi", "zhangbozhi"},
	}

	//生成json字符串
	jsonStr, err := json.Marshal(movie1)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", jsonStr)

	//myMovieJsonStr := "{\"title\":\"喜剧之王\",\"year\":1999,\"price\":20,\"actors\":[\"周星驰\",\"张柏芝\"]}"
	myMovieStruct := Movie{}
	err = json.Unmarshal(jsonStr, &myMovieStruct)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", myMovieStruct)
}
