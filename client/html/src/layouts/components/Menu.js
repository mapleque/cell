import React, { Component } from 'react'
import { Menu } from 'antd'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {
      collapsed: false,
    }
  }

  render() {
    return (
      <Menu
        mode='inline'
        theme='linght'
        inlineCollapsed={ this.state.collapsed }
      >
      </Menu>
    )
  }
}

export default Index

