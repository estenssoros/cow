import { LOGIN_SUCCESS, LOGIN_FAIL } from '../constants/action-types'
import API from '../modules/api.module'
import { URLS } from '../modules/endpoint.module'
import { toast } from 'react-toastify'

export const login = (username: string, password: string) => (dispatch: Function) => {
  var payload = {
    username: username,
    password: password
  }
  return API()
    .post(URLS.login, {}, {}, payload, {})
    .then(response => {
      toast.success('logged in and saved new config file')
      return dispatch({
        type: LOGIN_SUCCESS,
        payload: { data: response.data }
      })
    })
    .catch(error => {
      toast.error('failed to login')
      return dispatch({
        type: LOGIN_FAIL,
        payload: { data: error.response ? error.response.data : error }
      })
    })
}
