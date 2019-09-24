import React, { Component } from 'react'
import { Menu, Icon } from 'antd'
import { Link } from 'react-router-dom'

class User extends Component {
  constructor (props) {
    super(props)
    this.state = {
      profile: null,
    }
  }

  componentDidMount () {
    const authenticator = localStorage.getItem('_cell_authenticator')
    if (authenticator) {
      try {
        this.setState({ profile: JSON.parse(authenticator) })
      } catch (e) {
        console.error(e)
      }
    }
  }

  render () {
    const { profile } = this.state
    return profile ? (
      <Menu selectable={ false } mode='horizontal' style={{ height: 63, lineHeight: '64px', float: 'right', boxShadow: 'none' }}>
        <Menu.Item title='Logout' style={{ float: 'right', padding: '0 10px' }}>
          <Link style={{ display: 'inline' }} to='/logout'>
            <Icon type='logout'/>
          </Link>
        </Menu.Item>
        <Menu.Item title='Profile' style={{ float: 'right', padding: '0 10px' }}>
          <Link style={{ display: 'inline' }} to='/profile'>
            <Icon type='user'/>
          </Link>
        </Menu.Item>
        <Menu.Item title='Dashboard' style={{ float: 'right', padding: '0 10px' }}>
          <Link style={{ display: 'inline' }} to='/admin'>
            <Icon type='appstore'/>
          </Link>
        </Menu.Item>
      </Menu>
    ): (
      <Menu selectable={ false } mode='horizontal' style={{ height: 63, lineHeight: '64px', float: 'right', boxShadow: 'none' }}>
        <Menu.Item title='Login' style={{ float: 'right' }}>
          <Link to={'/login'}>
            <Icon type='login'/>
          </Link>
        </Menu.Item>
      </Menu>
    )
  }
}

export default User

