export const StartMenu = (props) => {
  return (
    <div>
      <button onClick={() => props.register()}>register</button>
      <button onClick={() => props.login()}>login</button>
    </div>
  )
}