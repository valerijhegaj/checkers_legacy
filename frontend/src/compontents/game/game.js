import Board from "./board/board";

export const Game = (props) => {
  const onClickEmpty = (i, j) => {
    props.onClickEmpty(
      i, j,
      props.state.gamename,
      props.state.from,
      props.state.to
    )
  }
  const onClickFigure = (i, j) => {
    props.onClickFigure(i, j)
  }
  return (
    <Board figures={props.state.figures}
           onClickEmpty={onClickEmpty}
           onClickFigure={onClickFigure}/>
  )
}