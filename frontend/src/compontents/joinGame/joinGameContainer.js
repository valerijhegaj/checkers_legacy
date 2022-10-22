import connect from "../../redux/react-redux/Connect";
import JoinGame from "./joinGame";
import {update, onClick} from "../../usingRedux/redusers/joinGame"

const mapPropsToState = (state) => {
  return {
    state: state.joinGame
  }
}

const JoinGameContainer =
  connect(mapPropsToState, {update, onClick})(JoinGame)

export default JoinGameContainer