import {switcherCondition, updateSwitcher} from "./switcher";

export const start = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.startScreen))
}

export const join = () => async (dispatch) => {
  dispatch(updateSwitcher(switcherCondition.joinScreen))
}