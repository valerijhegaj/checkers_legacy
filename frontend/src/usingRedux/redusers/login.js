import {authAPI} from "../../api/api";
import {switcherCondition, updateSwitcher} from "./switcher";

const ActionTypes = {
  UpdateLogin: "update login"
}

const initialState = {
  password:'',
  username:''
}

export const login = (state = initialState, action) => {
  switch (action.type) {
    case ActionTypes.UpdateLogin:
      return {username: action.username, password: action.password }
    default:
      return state
  }
}

export const update = (username: string, password: string) => {
  return {type: ActionTypes.UpdateLogin, username: username, password: password}
}

export const onClick = (username: string, password: string) => async (dispatch) => {
  if (username === "") {
    return
  }
  let response = await authAPI.login(username, password).catch(()=>{})
  if (response !== undefined) {
    dispatch(updateSwitcher(switcherCondition.mainMenu))
  }
}

export const back = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.startLoading))
}