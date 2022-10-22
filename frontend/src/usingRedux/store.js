import {CombineReducers} from "../redux/store";
import {login} from "./redusers/login";
import {createUser} from "./redusers/createUser";
import {ApplyMiddleware, CreateStoreWithMiddleware} from "../redux/middleware";
import {ThunkMiddleware} from "../redux/thunkMiddleware";
import {switcher} from "./redusers/switcher";
import {startGame} from "./redusers/startGame";
import {joinGame} from "./redusers/joinGame";
import {game} from "./redusers/game";

let reducers = CombineReducers({login, createUser, switcher, startGame, joinGame, game})

let store = CreateStoreWithMiddleware(reducers, ApplyMiddleware(ThunkMiddleware))

export default store
window.store = store