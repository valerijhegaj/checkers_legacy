import {authAPI} from "../../api/api";
import {switcherCondition, updateSwitcher} from "./switcher";
import {createConnection, updateGame} from "./game";

const ActionTypes = {
  Update: "update game name"
}

const initialState = {
  gamename: '',
  password: ''
}

export const startGame = (state = initialState, action) => {
  switch (action.type) {
    case ActionTypes.Update:
      return {gamename: action.gamename, password: action.password}
    default:
      return state
  }
}

export const update = (gamename: string, password: string) => {
  return {
    type: ActionTypes.Update,
    gamename: gamename,
    password: password
  }
}

export const createFor2 = (gamename: string, password: string) => async (dispatch) => {
  let response = await authAPI.createGame(gamename, password, {
    gamer0: 0,
    gamer1: 0
  }).catch(() => 1)
  if (response !== undefined) {
    await authAPI.loginGame(gamename, password)
    dispatch(updateGame(gamename))
    dispatch(createConnection(gamename))
    dispatch(updateSwitcher(switcherCondition.gameScreen))
  }
}

export const createFor1 = (gamename: string, password: string) => async (dispatch) => {
  let response = await authAPI.createGame(gamename, password, {
    gamer0: 0,
    gamer1: 1,
    level1: 3
  }).catch(() => 1)
  if (response !== undefined) {
    await authAPI.loginGame(gamename, password)
    dispatch(updateGame(gamename))
    dispatch(createConnection(gamename))
    dispatch(updateSwitcher(switcherCondition.gameScreen))
  }
}