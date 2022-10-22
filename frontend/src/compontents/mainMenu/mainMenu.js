export const MainMenu = (props) => {
  return (
    <div>
      <button onClick={() => {props.start()}}>start</button>
      <button onClick={() => {props.join()}}>join</button>
    </div>
  )
}