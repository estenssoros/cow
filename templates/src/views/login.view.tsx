import * as React from 'react'
import { Button, Form, Grid, Header, Segment } from 'semantic-ui-react'
import 'semantic-ui-css/semantic.min.css'
import { connect } from 'react-redux'
import { login } from '../actions/login.action'
import { Redirect } from 'react-router'

class Login extends React.Component<any, any> {
  state = {
    username: '',
    password: ''
  }
  handleSubmit() {
    this.props.login(this.state.username, this.state.password)
  }
  render() {
    if (this.props.isAuthorized) {
      return (
        <Redirect
          to={{
            pathname: '/dashboard'
          }}
        />
      )
    }

    return (
      <div>
        <Grid textAlign="center" style={{ height: '100%' }} verticalAlign="middle">
          <Grid.Column style={{ maxWidth: 450 }}>
            <Header as="h2" color="teal" textAlign="center" />
            <Form size="large">
              <Segment stacked>
                <Form.Input
                  fluid
                  icon="user"
                  iconPosition="left"
                  placeholder="username"
                  onChange={event => {
                    this.setState({ username: event.target.value })
                  }}
                />
                <Form.Input
                  fluid
                  icon="lock"
                  iconPosition="left"
                  placeholder="Password"
                  type="password"
                  onChange={event => {
                    this.setState({ password: event.target.value })
                  }}
                />
                <Button
                  color="teal"
                  fluid
                  size="large"
                  onClick={event => {
                    event.preventDefault()
                    this.handleSubmit()
                  }}
                >
                  Login
                </Button>
              </Segment>
            </Form>
          </Grid.Column>
        </Grid>
      </div>
    )
  }
}

const mapStateToProps = (state: any) => ({
  isAuthorized: state.user.isAuthorized,
  jwt: state.user.jwt
})

const mapDispatchToProps = {
  login
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Login)
