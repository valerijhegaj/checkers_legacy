import Connect from "../../redux/react-redux/Connect";
import {StartMenu} from "./startMenu";
import {login, register} from "../../usingRedux/redusers/startMenu";

const StartMenuContainer = Connect(()=>{}, {login,register})(StartMenu)

export default StartMenuContainer