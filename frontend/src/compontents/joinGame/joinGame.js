const JoinGame = (props) => {
  return (
    <div>
      Gamename
      <input type="text" value={props.state.gamename}
             onChange={event => {
               props.update(event.target.value, props.state.password)
             }}/>
      Password
      <input type="text" value={props.state.password}
             onChange={event => {
               props.update(props.state.gamename, event.target.value)
             }}/>
      <button onClick={() => {
        props.onClick(props.state.gamename, props.state.password)
      }}>login
      </button>
    </div>
  )
}

export default JoinGame