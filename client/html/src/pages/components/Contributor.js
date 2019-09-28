import React, { Component } from 'react'
import { Tag, Icon } from 'antd'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <a href={`https://github.com/${this.props.name}`}>
        <Tag>
          <Icon type='github'/>
          <span style={{ marginLeft: 4 }}>{this.props.name}</span>
        </Tag>
      </a>
    )
  }
}

export default Index

