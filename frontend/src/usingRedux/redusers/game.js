import {gameAPI} from "../../api/api";

class boardStorage {
  constructor() {
    this._data = Array(64)
  }

  Get(i, j) {
    return this._data[i * 8 + j]
  }

  Insert(i, j, element) {
    this._data[i * 8 + j] = element
  }
}

const ActionTypes = {
  update: "update game",
  updateFigures: "update game board",
  updateFrom: "update from game",
  updateTo: "update to game"
}

const initialState = {
  figures: new boardStorage(),
  gamename: "",
  from: undefined,
  to: []
}

export const game = (state = initialState, action) => {
  switch (action.type) {
    case ActionTypes.update:
      return {
        ...state,
        gamename: action.gamename
      }
    case ActionTypes.updateFigures:
      return {
        ...state,
        figures: action.figures
      }
    case ActionTypes.updateFrom:
      return {
        ...state,
        from: action.from,
        to: []
      }
    case ActionTypes.updateTo:
      return {
        ...state,
        to: [...state.to, action.to]
      }
    default:
      return state
  }
}

export const updateGame = (gamename) => {
  return {type: ActionTypes.update, gamename}
}

export const update = (figures) => {
  return {type: ActionTypes.updateFigures, figures}
}

export const updateTo = (to) => {
  return {type: ActionTypes.updateTo, to}
}

export const updateFrom = (from) => {
  return {type: ActionTypes.updateFrom, from}
}

export const getBoard = (gamename) => async (dispatch) => {
  const response = await gameAPI.getGame(gamename).catch(()=>0)
  if (response === undefined) {
    return
  }
  let figures = new boardStorage()
  response.data.figures.forEach((elem) => {
    figures.Insert(elem.x, elem.y, elem.figure + elem.gamerId.toString())
  })
  dispatch(update(figures))
  console.log("got")
}

export const createConnection = (gamename) => async (dispatch) => {
  await getBoard(gamename)(dispatch)
  setTimeout(()=>{dispatch(createConnection(gamename))}, 1000)
}

export const onClickFigure = (i, j) => {
  return updateFrom({x: i, y: j})
}

export const onClickEmpty = (i, j, gamename, from, to) =>
  async (dispatch) => {
  if (to.length === 0) {
    dispatch(updateTo({x: i, y: j}))
    return
  }

  const lastTo = to[to.length - 1]
  if (i !== lastTo.x || j !== lastTo.y) {
    dispatch(updateTo({x: i, y: j}))
    return
  }
  let response = await gameAPI.move(gamename, from, to).catch(()=>0)
  if (response !== undefined) {
    dispatch(getBoard(gamename))
  }
}
