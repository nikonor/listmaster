package main

import (
    "testing"
    // . "github.com/smartystreets/goconvey/convey"
    "fmt"
    "strings"
)

func TestCheckWords (t *testing.T) {
    cases := []struct {
        in string
        out []string
    }{
        {`/add продукты`,[]string{"/add","продукты"}},
        {`/list "Зоо магазин"`,[]string{"/list",`"Зоо магазин"`}},
        {`/add "Зоо магазин" "насыпку большую"`,[]string{`/add`,`"Зоо магазин"`,`"насыпку большую"`}},
    }

    for _,c := range cases {
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
        in string
        out_code int
        out_idx float32
        out_element string
        out_error string
    }{
        {`/add продукты`,1,0,`продукты`,""},
        {`/list "Зоо магазин"`,2,0,`Зоо магазин`,""},
        {`/add 1 "витамины сыну"`,1,1,"витамины сыну",""},
        {`/add "Зоо магазин" "насыпку большую"`,1,2,"насыпку большую",""},
        {`add "Зоо магазин" "насыпку большую"`,1,2,"насыпку большую","1"},
        {`/qwe wert eeeee`,0,0,"","1"},
    }

    for _,c := range cases {
        got_code,got_idx,got_element,got_error := ParseCommand(c.in)  
        if ( got_error != nil && c.out_error != "") {
            fmt.Printf("Right error: got=!%s!,out=!%s!\n",got_error,string(c.out_error))
        } else {
            if (got_code != c.out_code || got_idx != c.out_idx || got_element != c.out_element) {
                t.Errorf("ParseCommand wrong: in=%s\n\t!%d!<=>!%d!%t!\n\t!%f!<=>!%f!%t!\n\t!%s!<=>!%s!%t!\n", c.in, c.out_code,got_code,(c.out_code==got_code),c.out_idx,got_idx,(c.out_idx==got_idx),c.out_element,got_element,(c.out_element==got_element))
            }
        }
    }
}

