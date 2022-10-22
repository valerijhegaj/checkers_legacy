import Connect from "../../redux/react-redux/Connect"
import {MainMenu} from "./mainMenu";
import {join, start} from "../../usingRedux/redusers/mainMenu";

const MainMenuContainer = Connect(() => {}, {join, start})(MainMenu)

export default MainMenuContainer