import { combineReducers } from 'redux'
import user from './user.reducer'

const allReducers = combineReducers({
  user
})

interface Iaction {
  type: string
  payload: object
}

function rootReducer(state: any, action: Iaction) {
  return allReducers(state, action)
}

export default rootReducer
