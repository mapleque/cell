import React, { Component, Fragment } from 'react'
import { Layout } from 'antd'
import Header from './components/Header'
import Footer from './components/Footer'

class Web extends Component {
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
        <Layout.Content style={{ padding: '0 50px' }}>
          {React.Children.map(
            children,
            args => React.cloneElement(args, {
              route,
              request,
            })
          )}
        </Layout.Content>
        <Footer/>
      </Fragment>
    )
  }
}

export default Web

