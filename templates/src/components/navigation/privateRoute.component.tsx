import * as React from 'react'
import { Route, Redirect } from 'react-router-dom'
import { connect } from 'react-redux'

const PrivateRoute = ({ component: Component, isAuthorized, ...rest }) => {
  if (isAuthorized) {
    return (
      <Route
        {...rest}
        render={props =>
          isAuthorized ? (
            <Component {...props} />
          ) : (
            <Redirect
              to={{
                pathname: '/login',
                state: { from: props.location }
              }}
            />
          )
        }
      />
    )
  }
  return (
    <Redirect
      to={{
        pathname: '/login',
        state: { from: '' }
      }}
    />
  )
}

const mapStateToProps = (state: any) => ({
  isAuthorized: state.user.isAuthorized
})

export default connect(
  mapStateToProps,
  null
)(PrivateRoute)
