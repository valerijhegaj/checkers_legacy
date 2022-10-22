import {switcherCondition, updateSwitcher} from "./switcher";

export const register = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.createUser))
}

export const login = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.login))
}