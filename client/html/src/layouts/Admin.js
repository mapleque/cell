import React, { Component, Fragment } from 'react'
import { Layout } from 'antd'
import Header from '@/layouts/components/Header'
import Footer from '@/layouts/components/Footer'
import Menu from '@/layouts/components/Menu'

class Index extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  getPageTitle(currentRouter) {
    return `Cell-${currentRouter.name}`
  }

  render() {
    const { children, route, request } = this.props
    return (
      <Fragment>
        <Header route={route} request={request}/>
        <Layout.Content>
          <Layout style={{ background: '#fff' }}>
            <Layout.Sider theme='light' style={{ marginRight: 24 }}>
              <Menu route={route} request={request}/>
            </Layout.Sider>
            <Layout.Content>
              {React.Children.map(
                children,
                args => React.cloneElement(args, {
                  route,
                  request,
                })
              )}
            </Layout.Content>
          </Layout>
        </Layout.Content>
        <Footer/>
      </Fragment>
    )
  }
}

export default Index

