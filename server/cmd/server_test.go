package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"checkers/core"
	"checkers/saveLoad"
	"checkers/server/pkg/defines"
	"checkers/server/test/format"
	apiParser "checkers/test/api"
)

func Test_server(t *testing.T) {
	os.Chdir("..")
	defer os.Chdir("cmd")
	log.SetOutput(ioutil.Discard)
	go main()

	valerijhegaj := &apiParser.User{
		Username: "valerijhegaj", Password: "123", PORT: 4444,
	}

	//----------------------test1---------------------------------------
	// create user, log in
	{
		code, err := valerijhegaj.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = valerijhegaj.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		if valerijhegaj.IsEmptyCookies() {
			t.Error(format.ErrorString("cookies", "no cookies"))
		}
	}

	//----------------------test2---------------------------------------
	// try to create user with same nick
	{
		hacker := &apiParser.User{
			Username: valerijhegaj.Username, Password: "wrong", PORT: 4444,
		}
		code, err := hacker.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = hacker.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
	}

	gameName1, password1 := "fitstField", "1"
	firstField := core.NewStandard8x8Field()

	//----------------------test3---------------------------------------
	// create game, log in, get, move and get
	{
		code, err := valerijhegaj.CreateGame(
			gameName1, password1, defines.Settings{},
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = valerijhegaj.LogInGame(
			gameName1, password1,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, rawSave, err := valerijhegaj.GetGame(gameName1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		save := saveLoad.NewSaveFromRawSave(rawSave)

		if !core.IsEqual(&save.Field, &firstField) {
			t.Error(format.ErrorField(&firstField, &save.Field))
		}
		if save.TurnGamerId != 0 {
			t.Error(format.ErrorInt(0, save.TurnGamerId))
		}
		if save.Winner != -1 {
			t.Error(format.ErrorInt(-1, save.Winner))
		}

		from := core.Coordinate{2, 0}
		to := []core.Coordinate{{3, 1}}

		code, err = valerijhegaj.Move(gameName1, from, to)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, rawSave, err = valerijhegaj.GetGame(gameName1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		save = saveLoad.NewSaveFromRawSave(rawSave)

		figure := firstField.At(from)
		figure.Move(&firstField, from, to)

		if !core.IsEqual(&save.Field, &firstField) {
			t.Error(format.ErrorField(&firstField, &save.Field))
		}
		if save.TurnGamerId != 1 {
			t.Error(format.ErrorInt(0, save.TurnGamerId))
		}
		if save.Winner != -1 {
			t.Error(format.ErrorInt(-1, save.Winner))
		}
	}

	aboba := &apiParser.User{
		Username: "aboba", Password: "abob", PORT: 4444,
	}

	{
		code, err := aboba.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = aboba.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		if aboba.IsEmptyCookies() {
			t.Error(format.ErrorString("cookies", "no cookies"))
		}
	}

	//----------------------test4---------------------------------------
	// try to log in game with wrong password
	// try without log in to get field, move
	// try to create game with such name
	{
		code, err := aboba.LogInGame(gameName1, password1+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = aboba.LogInGame(gameName1+"evil", password1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusNotFound {
			t.Error(format.ErrorInt(http.StatusNotFound, code))
		}

		code, _, err = aboba.GetGame(gameName1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = aboba.Move(
			gameName1, core.Coordinate{5, 1}, []core.Coordinate{{4, 0}},
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = aboba.CreateGame(
			gameName1, password1, defines.Settings{},
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

	}

	//----------------------test5---------------------------------------
	// log in game, get, move and get
	move := func(
		isCorrect bool, from core.Coordinate, to []core.Coordinate,
		user *apiParser.User,
	) {
		code, err := user.Move(gameName1, from, to)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if isCorrect {
			if code != http.StatusCreated {
				t.Error(format.ErrorInt(http.StatusCreated, code))
			}
			figure := firstField.At(from)
			figure.Move(&firstField, from, to)
		} else {
			if code != http.StatusMethodNotAllowed {
				t.Error(format.ErrorInt(http.StatusMethodNotAllowed, code))
			}
		}

		code, rawSave, err := user.GetGame(gameName1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		save := saveLoad.NewSaveFromRawSave(rawSave)
		if !core.IsEqual(&firstField, &save.Field) {
			t.Error(format.ErrorField(&firstField, &save.Field))
		}
	}
	generateFromTo := func(data []int) (
		core.Coordinate, []core.Coordinate,
	) {
		from := core.Coordinate{data[0], data[1]}
		var to []core.Coordinate
		for i := 2; i < len(data); i += 2 {
			to = append(to, core.Coordinate{data[i], data[i+1]})
		}
		return from, to
	}

	{
		code, err := aboba.LogInGame(gameName1, password1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, rawSave, err := aboba.GetGame(gameName1)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		save := saveLoad.NewSaveFromRawSave(rawSave)
		if !core.IsEqual(&firstField, &save.Field) {
			t.Error(format.ErrorField(&firstField, &save.Field))
		}

		from, to := generateFromTo([]int{5, 1, 4, 0})
		move(true, from, to, aboba)
	}

	//----------------------test6---------------------------------------
	// 0 try to make wrong moves
	// 0 move to make 1 move to eat
	// 1 try move wrong
	// 1 eat two
	// 0 move to make 1 move to eat
	// 1 eat
	{
		from, to := generateFromTo([]int{3, 1, 2, 0})
		move(false, from, to, valerijhegaj)

		from, to = generateFromTo([]int{5, 3, 4, 2})
		move(false, from, to, valerijhegaj)
		move(false, from, to, aboba)

		from, to = generateFromTo([]int{2, 2, 3, 3})
		move(false, from, to, aboba)
		move(true, from, to, valerijhegaj)

		from, to = generateFromTo([]int{5, 3, 4, 2})
		move(false, from, to, valerijhegaj)

		from, to = generateFromTo([]int{4, 0, 2, 2, 4, 4})
		move(true, from, to, aboba)

		from, to = generateFromTo([]int{2, 4, 3, 3})
		move(true, from, to, valerijhegaj)

		from, to = generateFromTo([]int{4, 4, 2, 2})
		move(true, from, to, aboba)
	}

	//----------------------test7---------------------------------------
	//is authorized
	{
		isAuth, err := valerijhegaj.IsAuthorized()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if !isAuth {
			t.Error(format.ErrorInt(1, 0))
		}

		byba := apiParser.User{
			Username: "byba",
			PORT:     4444,
			Password: password1,
		}

		isAuth, err = byba.IsAuthorized()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if isAuth {
			t.Error(format.ErrorInt(0, 1))
		}
	}
}
