export const StartGame = (props) => {
  return (
    <div>
      Gamename
      <input type="text" value={props.state.gamename}
             onChange={event => props.update(event.target.value, props.state.password)} />
      Password
      <input type="text" value={props.state.password}
                        onChange={event => props.update(props.state.gamename, event.target.value)} />
      <button onClick={() => props.createFor2(props.state.gamename, props.state.password)}>create for 2 players</button>
      <button onClick={() => props.createFor1(props.state.gamename, props.state.password)}>create for 1 player</button>
    </div>
  )
}