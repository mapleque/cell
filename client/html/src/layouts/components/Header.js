import React, { Component, Fragment } from 'react'
import DocumentTitle from 'react-document-title'
import { Layout, Menu } from 'antd'
import { Link } from 'react-router-dom'
import logo from '@/assets/32.png'
import User from './User'
import Search from './Search'


class Header extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  getPageTitle (currentRouter) {
    return `Cell-${currentRouter.name}`
  }

  render () {
    const { route, request } = this.props
    return (
      <Fragment>
        <DocumentTitle title={this.getPageTitle(route.current)}/>
        <Layout.Header style={{ background: '#fff', padding: 0 }}>
          <Link to='/' style={{
            height: 32,
            width: 176,
            margin: '16px 24px 15px',
            lineHeight: '32px',
            float: 'left',
            display: 'flex',
          }}>
            <img src={logo} alt='Mapleque Cell'/>
            <h1 style={{
              margin: '0 0 0 12px',
              fontWeight: 600,
              fontSize: '20px',
              color: 'rgba(0, 0, 0, 0.65)',
            }}>Mapleque Cell</h1>
          </Link>
          <Menu selectable={ false } mode='horizontal' style={{ height: 64, lineHeight: '64px' }}>
            <Menu.Item>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/'}>Wellcome</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/document'}>Document</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/deployment'}>Deployment</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/contribution'}>Contribution</Link>
            </Menu.Item>
            <User request={request}/>
            <Menu.Item style={{ float: 'right', border: 'none' }}>
              <Search request={request}/>
            </Menu.Item>
          </Menu>
        </Layout.Header>
      </Fragment>
    )
  }
}

export default Header

