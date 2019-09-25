import React, { Component } from 'react'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  async componentDidMount() {
    await this.props.request('/user/logout', {})
    localStorage.removeItem('_cell_session')
    localStorage.removeItem('_cell_authenticator')
    window.location.replace('/')
  }

  render() {
    return (
      <div>tbd</div>
    )
  }
}

export default Index

