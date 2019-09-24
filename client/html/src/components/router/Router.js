import React, { Component } from 'react'
import { BrowserRouter, Route, Redirect, Switch } from 'react-router-dom'

function unwrapper (arr) {
  if (typeof arr === 'object') {
    let ret = []
    arr.forEach((e) => {
      if (typeof e === 'object') {
        e.forEach((v) => {
          ret.push(v)
        })
      }
    })
    return ret
  }
  return arr
}

class Router extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  renderRouters(routes, wrapper) {
    return unwrapper(routes.map(route => this.renderRouter(route, wrapper)))
  }

  renderRouter(route, wrapper) {
    const innerWrapper = wrapper ? [ ...wrapper ] : []
    if (route.routes && route.routes.length > 0) {
      innerWrapper.push(route)
      return this.renderRouters(route.routes, innerWrapper)
    }
    return [
      <Route
        key={route.path}
        exact={route.path === '/'}
        path={route.path}
        render={this.renderComponent(this.props.config, route, innerWrapper)}
      />
    ]
  }

  renderComponent(rootRoutes, currentRoute, wrapper) {
    const { request } = this.props
    return props => {
      if (currentRoute.redirect) {
        return <Redirect to={currentRoute.redirect}/>
      } else {
        const pathRoutes = [...wrapper, currentRoute]
        let tmpComponent = <currentRoute.component route={{
          config: rootRoutes,
          current: currentRoute,
          paths: pathRoutes,
        }} request={request}/>
        wrapper.reverse().forEach(wrapperRoute => {
          if (wrapperRoute.component) {
            tmpComponent = <wrapperRoute.component route={{
              config: rootRoutes,
              current: currentRoute,
              paths: pathRoutes,
            }} request={request}>{tmpComponent}</wrapperRoute.component>
          }
        })
        return tmpComponent
      }
    }
  }

  render() {
    return (
      <BrowserRouter>
        <Switch>
          {
            this.renderRouters(this.props.config)
          }
          <Redirect to="/404" />
        </Switch>
      </BrowserRouter>
    )
  }
}

export default Router

