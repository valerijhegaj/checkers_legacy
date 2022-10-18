package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

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
	// tries to create user with same nick
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
}
