import React, { Component } from 'react'
import { Input } from 'antd'

class Search extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Input.Search placeholder='Search in site'/>
    )
  }
}

export default Search

