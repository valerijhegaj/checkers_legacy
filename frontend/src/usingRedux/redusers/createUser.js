import {authAPI} from "../../api/api";
import {switcherCondition, updateSwitcher} from "./switcher";

const ActionTypes = {
  UpdateCreateUser: "update create user"
}

const initialState = {
  password: '',
  username: ''
}

export const createUser = (state = initialState, action) => {
  switch (action.type) {
    case ActionTypes.UpdateCreateUser:
      return {username: action.username, password: action.password}
    default:
      return state
  }
}

export const update = (username: string, password: string) => {
  return {type: ActionTypes.UpdateCreateUser, username: username, password: password}
}

export const onClick = (username: string, password: string) => async (dispatch) => {
  if (username === "") {return}
  await authAPI.register(username, password).catch(error => {
    console.log(error.response.status)
  })
  let response = await authAPI.login(username, password).catch(()=>{})
  if (response !== undefined) {
    dispatch(updateSwitcher(switcherCondition.mainMenu))
  }
}

export const back = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.startLoading))
}