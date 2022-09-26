package test

//
//import (
//	"chekers/core"
//	"chekers/saveLoad"
//	"io/ioutil"
//	"testing"
//)
//
//func getTestSave() saveLoad.Save {
//	var save saveLoad.Save
//	save.Field = getTestField()
//	save.Field.Put(core.Coordinate{1, 2}, core.Checker{1})
//	save.Field.Put(core.Coordinate{2, 3}, core.King{1})
//	save.Field.Put(core.Coordinate{3, 4}, core.Checker{0})
//	save.Field.Put(core.Coordinate{4, 5}, core.King{0})
//	save.Master = saveLoad.Participants{1, 1}
//	return save
//}
//
//func getTestSaveInString() string {
//	return "{\"figures\":" +
//		"[{\"x\":1,\"y\":2,\"figure\":\"Checker\",\"gamerId\":1}," +
//		"{\"x\":2,\"y\":3,\"figure\":\"King\",\"gamerId\":1}," +
//		"{\"x\":3,\"y\":4,\"figure\":\"Checker\",\"gamerId\":0}," +
//		"{\"x\":4,\"y\":5,\"figure\":\"King\",\"gamerId\":0}]," +
//		"\"bordersRight\":{\"x\":7,\"y\":7}," +
//		"\"bordersLeft\":{\"x\":0,\"y\":0}," +
//		"\"position\":{\"gamer0\":1,\"gamer1\":1}," +
//		"\"turnGamerId\":0}"
//}
//
//func TestSave_Write(t *testing.T) {
//	save := getTestSave()
//
//	tempDir := t.TempDir()
//	err := save.Write(tempDir + "/test.json")
//	if err != nil {
//		t.Error()
//	}
//	rawSave, err := ioutil.ReadFile(tempDir + "/test.json")
//	if err != nil {
//		t.Error(err.Error())
//	}
//	if string(rawSave) != getTestSaveInString() { //incorrect sometimes (f* hashes)
//		t.Error(string(rawSave))
//	}
//}
//
//func TestSave_Read(t *testing.T) {
//	var save saveLoad.Save
//	saveToCreate := getTestSave()
//
//	tempDir := t.TempDir()
//	saveToCreate.Write(tempDir + "/test.json")
//	err := save.Read(tempDir + "/test.json")
//	if err != nil {
//		t.Error()
//	}
//	//not implemented
//
//}
