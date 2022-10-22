export const ThunkMiddleware = store => dispatch => action => {
  if (typeof action === 'function') {
    return action(store.Dispatch.bind(store))
  }
  return dispatch(action)
}
