import * as React from 'react'
import { connect } from 'react-redux'
import NavBarComponent from '../components/navigation/navbar.component'

class Dashboard extends React.Component<any, any> {
  render() {
    return (
      <div>
        <NavBarComponent />
      </div>
    )
  }
}

const mapStateToProps = (state: any) => ({})

const mapDispatchToProps = {}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard)
