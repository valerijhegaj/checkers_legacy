package saveLoad

import (
	"chekers/core"
	"encoding/json"
	"io/ioutil"
	"reflect"
)

const (
	Man = iota
	Bot
)

type Master struct {
	Gamer0 int `json:"gamer0"`
	Gamer1 int `json:"gamer1"`
}

func GetSaveList(path string) ([]string, error) {
	var saveList []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return saveList, err
	}
	for _, file := range files {
		if !file.IsDir() {
			saveList = append(saveList, file.Name())
		}
	}

	return saveList, err
}

type Save struct {
	Field       core.Field
	Master      Master
	TurnGamerId int
}

func (c *Save) putFiguresOnField(figures []figureInfo) {
	for _, i := range figures {
		if i.Figure == "Checker" {
			c.Field.Figures[core.Coordinate{i.X, i.Y}] = core.Checker{i.GamerId}
		} else if i.Figure == "King" {
			c.Field.Figures[core.Coordinate{i.X, i.Y}] = core.King{i.GamerId}
		}
	}
}

func (c *Save) InitFromJsonSave(jsonSave *JsonSave) {
	c.Field.Figures = make(map[core.Coordinate]core.Figure)
	c.putFiguresOnField(jsonSave.Figures)
	c.Field.BordersRight = jsonSave.BordersRight
	c.Field.BordersLeft = jsonSave.BordersLeft
	c.Master = jsonSave.Position
	c.TurnGamerId = jsonSave.TurnGamerId
}

func (c *Save) Read(path string) error {
	var helper JsonSave
	err := helper.Read(path)
	c.InitFromJsonSave(&helper)
	return err
}

func (c *Save) Write(path string) error {
	var helper JsonSave
	helper.initFromSave(c)
	return helper.write(path)
}

type figureInfo struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Figure  string `json:"figure"`
	GamerId int    `json:"gamerId"`
}

type JsonSave struct {
	Figures      []figureInfo    `json:"figures"`
	BordersRight core.Coordinate `json:"bordersRight"`
	BordersLeft  core.Coordinate `json:"bordersLeft"`
	Position     Master          `json:"position"`
	TurnGamerId  int             `json:"turnGamerId"`
}

// warning reflect.TypeOf(figure).String()[5:] can don't work with other names and struct of project
func (c *JsonSave) takeFiguresFromField(field core.Field) {
	c.Figures = make([]figureInfo, len(field.Figures))
	i := 0
	for coordinate, figure := range field.Figures {
		c.Figures[i].X = coordinate.X
		c.Figures[i].Y = coordinate.Y
		c.Figures[i].Figure = reflect.TypeOf(figure).String()[5:]
		c.TurnGamerId = figure.GetOwnerId()
		i++
	}
}

func (c *JsonSave) initFromSave(save *Save) {
	c.Position = save.Master
	c.TurnGamerId = save.TurnGamerId
	c.BordersRight = save.Field.BordersRight
	c.BordersLeft = save.Field.BordersLeft
	c.takeFiguresFromField(save.Field)
}

func (c *JsonSave) write(path string) error {
	rawSave, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, rawSave, 0644)
	return err
}

func (c *JsonSave) Read(path string) error {
	rawSave, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawSave, c)
	return err
}
