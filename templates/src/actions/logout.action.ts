import { LOGOUT_SUCCESS, LOGOUT_FAIL } from '../constants/action-types'
import API from '../modules/api.module'
import { URLS } from '../modules/endpoint.module'
import { store } from '../store'

export const logout = () => (dispatch: Function) => {
  return API()
    .post(URLS.logout, {}, {}, { jwt: store.getState().user.jwtToken }, {})
    .then(response => {
      return dispatch({
        type: LOGOUT_SUCCESS,
        payload: { data: response.data }
      })
    })
    .catch(error => {
      return dispatch({
        type: LOGOUT_FAIL,
        payload: { data: error.response ? error.response.data : error }
      })
    })
}
