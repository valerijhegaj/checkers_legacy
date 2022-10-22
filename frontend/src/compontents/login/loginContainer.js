import Login from "./login";
import {onClick, update, back} from "../../usingRedux/redusers/login";
import Connect from "../../redux/react-redux/Connect";

const mapStateToProps = (state) => {
  return {
    state: state.login
  }
}

const LoginContainer =
  Connect(mapStateToProps,
    {update, onClick, back})(Login)

export default LoginContainer