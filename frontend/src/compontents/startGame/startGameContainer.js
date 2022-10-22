import Connect from "../../redux/react-redux/Connect";
import {StartGame} from "./startGame";
import {update, createFor2, createFor1} from "../../usingRedux/redusers/startGame";

const mapStateToProps = (state) => {
  return {
    state: state.startGame
  }
}

const StartGameContainer = Connect(mapStateToProps, {update, createFor2, createFor1})(StartGame)

export default StartGameContainer