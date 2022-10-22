import {onClick, update, back} from "../../usingRedux/redusers/createUser";
import Connect from "../../redux/react-redux/Connect";
import CreateUser from "./createUser";

const mapStateToProps = (state) => {
  return {
    state: state.createUser
  }
}

const CreateUserContainer = Connect(mapStateToProps, {update, onClick, back})(CreateUser)

export default CreateUserContainer