package main

import (
    "testing"
    // . "github.com/smartystreets/goconvey/convey"
    // "fmt"
)

func TestParseCommand(t *testing.T) {
    cases := []struct {
        in string
        out_code int
        out_idx float32
        out_element string
    }{
        {`/add продукты`,1,0,`продукты`},
        {`/list "Зоо магазин"`,2,0,`Зоо магазин`},
        {`/add 1 "витамины сыну"`,1,1,"витамины сыну"},
        {`/add "Зоо магазин" "насыпку большую"`,1,2,"насыпку большую"},
    }

    for _,c := range cases {
        got_code,got_idx,got_element,got_err := ParseCommand(c.in)  
        if got_err != nil {
            t.Errorf("Error")
        } else {
            if (got_code != c.out_code || got_idx != c.out_idx || got_element != c.out_element ) {
                t.Errorf("ParseCommand wrong: in=%s\n\t!%d!<=>!%d!\n\t!%f!<=>!%f!\n\t!%s!<=>!%s!\n", c.in, c.out_code,got_code,c.out_idx,got_idx,c.out_element,got_element)
            }
        }
    }
}