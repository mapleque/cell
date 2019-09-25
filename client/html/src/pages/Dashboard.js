import React, { Component, Fragment } from 'react'
import { Row, Col, Modal } from 'antd'
import AuthList from './components/AuthList'
import AppList from './components/AppList'
import AppCreateForm from './components/AppCreateForm'
import AppUpdateForm from './components/AppUpdateForm'
import AppDeleteForm from './components/AppDeleteForm'
import AuthDeleteForm from './components/AuthDeleteForm'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {
      authList: [],
      appList: [],

      showCreateAppModal: false,
      showEditAppModal: false,
      showDeleteAppModal: false,
      showDeleteAuthModal: false,

      currentApp: null,
      currentAuth: null,
    }
  }

  componentDidMount() {
    this.loadAppList()
    this.loadAuthList()
  }

  async loadAppList() {
    const { list } = await this.props.request('/app/list', {})
    this.setState({
      appList: list,
      appDisplayList: this.filter(list, this.state.appSearch),
    })
  }

  async loadAuthList() {
    const { list } = await this.props.request('/authorization/list', {})
    this.setState({
      authList: list,
      authDisplayList: this.filter(list, this.state.authSearch),
    })
  }

  filter(list, search) {
    return list.filter(item => item.name.indexOf(search) >= 0)
  }

  render() {
    return (
      <Fragment>
        <Row gutter={16} style={{ margin: '16px 0 24px' }}>
          <Col span={12}>
            <AuthList list={this.state.authList} onDelete={item => {
              this.setState({ currentAuth: item }, () => {
                this.setState({ showDeleteAuthModal: true })
              })
            }}/>
          </Col>
          <Col span={12}>
            <AppList
              list={this.state.appList}
              onCreate={() => {
                this.setState({ showCreateAppModal: true })
              }}
              onUpdate={item => {
                this.setState({ currentApp: item }, () => {
                  this.setState({ showUpdateAppModal: true })
                })
              }}
              onDelete={item => {
                this.setState({ currentApp: item }, () => {
                  this.setState({ showDeleteAppModal: true })
                })
              }}
            />
          </Col>
        </Row>
        <Modal
          visible={ this.state.showCreateAppModal }
          footer={null}
          closable={false}
        >
          <AppCreateForm
            onSubmit={async item => {
              await this.props.request('/app/create', item)
              await this.loadAppList()
              this.setState({
                showCreateAppModal: false,
              })
            }}
            onCancel={() => {
              this.setState({
                showCreateAppModal: false,
              })
            }}
          />
        </Modal>
        <Modal
          visible={ this.state.showUpdateAppModal }
          footer={null}
          closable={false}
        >
          <AppUpdateForm
            app={this.state.currentApp}
            onSubmit={async item => {
              await this.props.request('/app/update', { app_id: this.state.currentApp.app_id, ...item })
              await this.loadAppList()
              this.setState({
                showUpdateAppModal: false,
              })
            }}
            onCancel={() => {
              this.setState({
                showUpdateAppModal: false,
              })
            }}
          />
        </Modal>
        <Modal
          visible={ this.state.showDeleteAppModal }
          footer={null}
          closable={false}
        >
          <AppDeleteForm
            app={this.state.currentApp}
            onSubmit={async () => {
              await this.props.request('/app/delete', { app_id: this.state.currentApp.app_id })
              await this.loadAppList()
              this.setState({
                showDeleteAppModal: false,
              })
            }}
            onCancel={() => {
              this.setState({
                showDeleteAppModal: false,
              })
            }}
          />
        </Modal>
        <Modal
          visible={ this.state.showDeleteAuthModal }
          footer={null}
          closable={false}
        >
          <AuthDeleteForm
            auth={this.state.currentAuth}
            onSubmit={async () => {
              await this.props.request('/authorization/delete', { id: this.state.currentAuth.id })
              await this.loadAuthList()
              this.setState({
                showDeleteAuthModal: false,
              })
            }}
            onCancel={() => {
              this.setState({
                showDeleteAuthModal: false,
              })
            }}
          />
        </Modal>
      </Fragment>
    )
  }
}

export default Index

