import s from "./board.module.css"
import Plate from "./plate/plate";
import React from "react"

const Board = (props) => {
  let container = []
  for (let i = 7; i >= 0; i--) {
    for (let j = 0; j < 8; j++) {
      const color = (i + j) % 2 === 0 ? "black" : "white"
      const figure = props.figures.Get(i, j)
      let onClick = () => props.onClickEmpty(i, j)
      if (figure !== undefined) {
        onClick = () => props.onClickFigure(i, j)
      }
      container.push({color, figure, onClick})
    }
  }

  return (
    <div className={s.board}>
      {container.map((element, index) => (
        <React.Fragment key={index}>
          <Plate color={element.color} figure={element.figure}
                 onClick={element.onClick}/>
        </React.Fragment>
      ))}
    </div>
  )
}

export default Board