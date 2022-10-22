import LoginContainer from "../login/loginContainer";
import CreateUserContainer from "../creteUser/createUserContainer";
import {switcherCondition} from "../../usingRedux/redusers/switcher";
import StartMenuContainer from "../startMenu/startMenuContainer";
import MainMenuContainer from "../mainMenu/mainMenuContainer";
import StartGameContainer from "../startGame/startGameContainer";
import JoinGameContainer from "../joinGame/joinGameContainer";
import GameContainer from "../game/gameContainer";

export const Switcher = (props) => {
  switch (props.state.condition) {
    case switcherCondition.startLoading:
      props.startLoad()
      return <div>loading</div>
    case switcherCondition.login:
      return <LoginContainer />
    case switcherCondition.createUser:
      return <CreateUserContainer />
    case switcherCondition.mainMenu:
      return <MainMenuContainer />
    case switcherCondition.startMenu:
      return <StartMenuContainer />
    case switcherCondition.startScreen:
      return <StartGameContainer />
    case switcherCondition.joinScreen:
      return <JoinGameContainer />
    case switcherCondition.gameScreen:
      return <GameContainer />
    default:
      return <div>loading</div>
  }
}
