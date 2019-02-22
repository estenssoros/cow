import * as React from 'react'
import './App.css'
import { Route, Switch, withRouter } from 'react-router-dom'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'

import PrivateRoute from './components/navigation/privateRoute.component'
import Login from './views/login.view'
import Dashboard from './views/dashboard.view'

class App extends React.Component<any, any> {
  render() {
    return (
      <div>
        <ToastContainer />
        {
          <Switch>
            <Route path="/login" component={Login} />
            <Route exact={true} path="/" component={Login} />
            <div>
              <Switch>
                <PrivateRoute path="/dashboard" exact component={Dashboard} />
              </Switch>
            </div>
          </Switch>
        }
      </div>
    )
  }
}
export default withRouter(App)
