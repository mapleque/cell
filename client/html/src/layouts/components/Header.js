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
        <Layout.Header style={{ background: '#fff' }}>
          <Link to='/' style={{
            height: 32,
            margin: '16px 24px 15px 0',
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
              <Link to={'/'}>Home</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/product'}>Product</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/doc'}>Document</Link>
            </Menu.Item>
            <Menu.Item>
              <Link to={'/about'}>About</Link>
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

