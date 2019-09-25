import React, { Component } from 'react'
import { Layout } from 'antd'

class Footer extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Layout.Footer style={{
        textAlign: 'center',
        background: '#fff',
        padding: '0',
      }}>
        <div style={{
          borderTop: '1px solid #e8e8e8',
          padding: '12px 0 0',
        }}>
          © 2013 - {new Date().getFullYear()} @ mapleque.com
        </div>
      </Layout.Footer>
    )
  }
}

export default Footer

