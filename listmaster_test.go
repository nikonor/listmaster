package main

import (
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
	// "fmt"
	"strings"
)

func TestGetMaxIdx(t *testing.T) {
	DevData = []ListElement{
		{1, "Аптека"},
		{1.001, "Канефрон"},
		{1.002, "Йод"},
		{2, "Зоо магазин"},
		{2.001, "Феликс 10 пакетиков"},
		{3, "Овощи, фрукты"},
	}
	out := GetMaxIdx(DevData)
	if out != 4 {
		t.Errorf("GetMaxIdx wrong. Test 1. out=%v!", out)
	}

	DevData = []ListElement{
		{1, "Аптека"},
		{1.001, "Канефрон"},
		{1.002, "Йод"},
		{2, "Зоо магазин"},
		{2.001, "Феликс 10 пакетиков"},
	}
	out = GetMaxIdx(DevData)
	if out != 3 {
		t.Errorf("GetMaxIdx wrong. Test 2. out=%v!", out)
	}

	DevData = []ListElement{
		{1, "Аптека"},
		{2, "Овощи"},
	}
	out = GetMaxIdx(DevData)
	if out != 3 {
		t.Errorf("GetMaxIdx wrong. Test 3. out=%v!", out)
	}

	DevData = []ListElement{}
	out = GetMaxIdx(DevData)
	if out != 1 {
		t.Errorf("GetMaxIdx wrong. Test 4. out=%v!", out)
	}

}

func TestGetCommandCode(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{`/add`, 1},
		{`/list`, 2},
		{`/qwe`, 0},
	}

	for _, c := range cases {
		out, _ := GetCommandCode(c.in)
		if c.out != out {
			t.Errorf("GetCommandCode wrong: %d!=%d", c.out, out)
		}
	}

}

func TestGetListIdx(t *testing.T) {
	cases := []struct {
		in_code int
		in_word string
		out_idx float32
	}{
		{1, "", 0},
		{1, "1", 1},
		{1, "2.2", 2.2},
		{1, "Зоо магазин", 2},
		{1, "продукты", 0},
	}

	DevData := []ListElement{
		{1, "Аптека"},
		{1.1, "Канефрон"},
		{1.2, "Йод"},
		{2, "Зоо магазин"},
		{2.1, "Феликс 10 пакетиков"},
	}

	for _, c := range cases {
		out_idx, _ := GetListIdx(c.in_code, c.in_word, DevData)
		if out_idx != c.out_idx {
			t.Errorf("GetListIdx wrong: %v!=%v", out_idx, c.out_idx)
		}
	}

	// GetListIdx
}

func TestCheckWords(t *testing.T) {
	cases := []struct {
		in  string
		out []string
	}{
		{`/add продукты`, []string{"/add", "продукты"}},
		{`/list "Зоо магазин"`, []string{"/list", `Зоо магазин`}},
		{`/add "Зоо магазин" "насыпку большую"`, []string{`/add`, `Зоо магазин`, `насыпку большую`}},
	}

	for _, c := range cases {
		out := CheckWords(strings.Fields(c.in))
		if len(c.out) != len(out) {
			t.Errorf("Len of array not equal")
		}

		isRight := true
		for i := range out {
			if out[i] != c.out[i] {
				isRight = false
			}
		}
		if isRight == false {
			t.Errorf("Arrays is not equal")
		}
	}
}

func TestParseCommand(t *testing.T) {
	cases := []struct {
		in          string
		out_code    int
		out_idx     float32
		out_element string
		out_error   string
	}{
		// исходная строка                     код индес элемент            есть ошибка
		{`/add продукты`, 1, 0, `продукты`, ""},
		{`/list "Зоо магазин"`, 2, 2, `Зоо магазин`, ""},
		{`/add 1 "витамины сыну"`, 1, 1, "витамины сыну", ""},
		{`/add 1 "чай"`, 1, 1, "чай", ""},
		{`/add "Зоо магазин" "насыпку большую"`, 1, 2, "насыпку большую", ""},
		{`add "Зоо магазин" "насыпку большую"`, 1, 2, "насыпку большую", "1"},
		{`/qwe wert eeeee`, 0, 0, "", "1"},
		{`/list`, 2, 0, "", ""},
	}

	DevData := []ListElement{
		{1, "Аптека"},
		{1.1, "Канефрон"},
		{1.2, "Йод"},
		{2, "Зоо магазин"},
		{2.1, "Феликс 10 пакетиков"},
	}

	for _, c := range cases {
		got_code, got_idx, got_element, got_error := ParseCommand(c.in, DevData)
		if got_error != nil && c.out_error != "" {
			// fmt.Printf("Right error: got=!%s!,out=!%s!\n",got_error,string(c.out_error))
		} else {
			if got_code != c.out_code || got_idx != c.out_idx || got_element != c.out_element {
				t.Errorf("ParseCommand wrong: in=%s\n\tcode=!%d!<=>!%d!%t!\n\tidx=!%f!<=>!%f!%t!\n\tel=!%s!<=>!%s!%t!\n", c.in, c.out_code, got_code, (c.out_code == got_code), c.out_idx, got_idx, (c.out_idx == got_idx), c.out_element, got_element, (c.out_element == got_element))
			}
		}
	}
}
