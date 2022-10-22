import {Game} from "./game";
import Connect from "../../redux/react-redux/Connect";
import {
  onClickEmpty,
  onClickFigure
} from "../../usingRedux/redusers/game";

const mapStateToProps = (state) => {
  return {
    state: state.game
  }
}

const GameContainer = Connect(mapStateToProps, {onClickFigure, onClickEmpty})(Game)

export default GameContainer