import React, { Component, Fragment } from 'react'
import axios from 'axios'

const tryToCall = (f, params) => typeof f === 'function' ? f(params) : params

class Server extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  async request(path, params) {
    const server = this.props.servers.find(server => {
      const r = server.path_rule
      const t = typeof r
      switch (t) {
        case 'function':
          return tryToCall(r, path)
        case 'string':
          return path.indexOf(r) >= 0
        case 'boolean':
          return r
        case 'object':
          if (r instanceof RegExp) {
            return r.test(path)
          }
          console.error('unsupport object of path_rule', r)
          break
        default:
          console.error('unsupport type of path_rule', t)
      }
      return false
    })
    return await this._request(server, path ,params)
  }

  async _request(server, path, params) {
    let request = {path, params, headers:{}, resp: null, err: null }
    try {
      if (server) {
        // do something with server settings
        request = tryToCall(server.before, request) || request
      }
      request.resp = await axios.post(request.path, request.params, {
        withCredentials: true,
        headers: request.headers,
      })
    }catch (err) {
      request.err = err
    }
    return tryToCall(server.after, request) || request
  }

  render() {
    const request = this.request.bind(this)
    return (
      <Fragment>
        {
          React.Children.map(
            this.props.children,
            args => React.cloneElement( args, { request })
          )
        }
      </Fragment>
    )
  }
}

Server.defaultProps = {
  // There could define multiple servers,
  // which could be find with request path.
  servers: [
    {
      // Find server by this rule with path.
      // This rule support:
      //   - boolean -- use directly
      //   - fuction -- call with param path
      //   - string -- deal as prefix of path
      //   - RegExp -- test regexp with path
      //   - Other -- return false and console log error
      path_rule: true,

      // Do something before request.
      // The return values will be used in request.
      // This must be a function.
      before: request => {},

      // Do something after request.
      // The return values will be return to app.
      // This must be a function
      after: response => {},
    },
  ],
}

export default Server

