import {CreateStore, Store} from "./store";

export function CreateStoreWithMiddleware(reducer, enhancer) {
  return enhancer(CreateStore)(reducer)
}

class WrappedStore extends Store {
  Dispatch(action) {
    this._middleware(this)(super.Dispatch.bind(this))(action)
  }

  constructor(father, middleware) {
    super()
    for (let key in father) {
      this[key] = father[key]
    }
    this._middleware = middleware
  }
}

export const ApplyMiddleware = (middleware) => (createStore) => (reducer) => {
  return new WrappedStore(createStore(reducer), middleware)
}
