import * as React from 'react'
import { Menu, Image } from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import { connect } from 'react-redux'
import { logout } from '../../actions/logout.action'

const menuStyle = {
  // border: 'none',
  borderRadius: 0
  // boxShadow: 'none',
  // marginBottom: '1em',
  // transition: 'box-shadow 0.5s ease, padding 0.5s ease'
}

class NavBarComponent extends React.Component<any, any> {
  stickTopMenu = () => this.setState({ menuFixed: true })
  unStickTopMenu = () => this.setState({ menuFixed: false })

  render() {
    return (
      <div>
        <Menu borderless style={menuStyle} inverted size="massive" attached="top">
          <Menu.Item>
            <Link to="/">
              <Image size="small" src="" />
            </Link>
          </Menu.Item>
          <Menu.Item as="a">
            <Link to="/">Option 1</Link>
          </Menu.Item>
          <Menu.Item as="a">
            <Link to="/">Option 2</Link>
          </Menu.Item>
          <Menu.Menu position="right">
            <Menu.Item
              name="logout"
              onClick={_ => {
                this.props.logout()
              }}
            >
              Logout
            </Menu.Item>
          </Menu.Menu>
        </Menu>
      </div>
    )
  }
}

const mapStateToProps = (state: any) => ({})

const mapDispatchToProps = {
  logout
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(NavBarComponent)
