import Connect from "../../redux/react-redux/Connect";
import {startLoad} from "../../usingRedux/redusers/switcher";
import {Switcher} from "./switcher";

const mapStateToProps = (state) => {
  return {
    state: state.switcher
  }
}

const SwitcherContainer = Connect(mapStateToProps, {startLoad})(Switcher)

export default SwitcherContainer