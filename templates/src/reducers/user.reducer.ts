import * as decoder from 'jwt-decode'

import { LOGIN_SUCCESS, LOGIN_FAIL, LOGOUT_SUCCESS, LOGOUT_FAIL } from '../constants/action-types'

const initialState = {
  user: '',
  isAuthorized: false,
  jwt: {},
  jwtToken: ''
}

export default function(state = initialState, action: any) {
  switch (action.type) {
    case LOGIN_SUCCESS: {
      return {
        ...state,
        isAuthorized: true,
        jwtToken: action.payload.data,
        jwt: (decoder as any)(action.payload.data)
      }
    }

    case LOGIN_FAIL: {
      return {
        ...state,
        isAuthorized: false,
        jwt: {}
      }
    }

    case LOGOUT_SUCCESS: {
      return {
        ...state,
        isAuthorized: false,
        jwt: {}
      }
    }

    case LOGOUT_FAIL: {
      return {
        ...state,
        isAuthorized: false,
        jwt: {}
      }
    }
  }
  return state
}
