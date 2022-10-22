export class Store {
  GetState() {
    return this._state
  }

  Subscribe(observer) {
    this._observers.add(observer)
    return () => this._observers.delete(observer)
  }

  Dispatch(action) {
    this._state = this._reducer(this._state, action)
    this._observers.forEach(observer => observer())
  }

  constructor() {
    this._state = {}
    this._observers = new Set()
    this._reducer = {}
  }

  _addReducer(reducer) {
    this._reducer = reducer
    this._state = reducer(this._state, {type: undefined})
  }
}

export const CombineReducers = (reducers) => (state, action) => {
  let isChanged = false
  let nextState = {}
  for (let key in reducers) {
    nextState[key] = reducers[key](state[key], action)
    isChanged = isChanged || (nextState[key] !== state[key])
  }
  if (isChanged) {
    return nextState
  }
  return state
}

export function CreateStore(reducer) {
  let store = new Store()
  store._addReducer(reducer)
  return store
}

