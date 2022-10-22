import s from "./plate.module.css"

const Plate = (props) => {
  let Figure
  switch (props.figure) {
    case "Checker0":
      Figure = () => <div className={`${s.checker} ${s.gamerID0}`}/>
      break
    case "Checker1":
      Figure = () => <div className={`${s.checker} ${s.gamerID1}`}/>
      break
    case "King0":
      Figure = () => (
        <div className={`${s.king} ${s.gamerID0}`}>
          K
        </div>
      )
      break
    case "King1":
      Figure = () => (
        <div className={`${s.king} ${s.gamerID1}`}>
          K
        </div>
      )
      break
    default:
      Figure = () => <div/>
      break
  }
  switch (props.color) {
    case "black":
      return <div className={`${s.plate} ${s.black}`} onClick={props.onClick}><Figure/></div>
    case "white":
      return <div className={`${s.plate} ${s.white}`}  onClick={props.onClick}><Figure/></div>
    default:
      return <div className={s.plate} onClick={props.onClick}/>
  }
}

export default Plate