import {authAPI} from "../../api/api";
import {switcherCondition, updateSwitcher} from "./switcher";
import {createConnection, updateGame} from "./game";

const ActionTypes = {
  update: "update join game"
}

const initialState = {
  gamename: '',
  password: ''
}

export const joinGame = (state = initialState, action) => {
  switch (action.type) {
    case ActionTypes.update:
      return {gamename: action.gamename, password: action.password}
    default:
      return state
  }
}

export const update = (gamename, password) => {
  return {
    type: ActionTypes.update,
    gamename: gamename,
    password: password
  }
}

export const onClick = (gamename: string, password: string) => async (dispatch) => {
  if (gamename === "") {
    return
  }
  const response = await authAPI.loginGame(gamename, password).catch(() => 0)
  if (response !== undefined) {
    dispatch(updateGame(gamename))
    dispatch(createConnection(gamename))
    dispatch(updateSwitcher(switcherCondition.gameScreen))
  }
}