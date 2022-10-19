import {Route} from "react-router-dom";
import {Ways} from "./pages/ways/ways";

function App() {
  return (
    <div>
      <Route path={Ways.Root} element={<Authorization/>}/>
      <Route path={Ways.Authorization} element={<Authorization/>}/>
      <Route path={Ways.Register} element={<Register/>}/>
      <Route path={Ways.Login} element={<Login/>}/>
    </div>
  );
}

export default App;
