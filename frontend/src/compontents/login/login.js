const Login = (props) => {
  return (
    <div>
      <input type="text" value={props.state.username}
             onChange={event => {
               props.update(event.target.value, props.state.password)
             }}/>
      <input value={props.state.password}
             onChange={event => {
               props.update(props.state.username, event.target.value)
             }}/>
      <button
        onClick={() => {
          props.onClick(props.state.username, props.state.password)
        }}>login
      </button>
      <button onClick={() => props.back()}>return</button>
    </div>
  )
}

export default Login