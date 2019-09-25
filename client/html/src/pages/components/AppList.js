import React, { Component } from 'react'
import { Card, List, Input, Tooltip, Button } from 'antd'
import moment from 'moment'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {
      search: '',
      list: [],
      displayList: [],
      showSecret: false,
    }
  }

  componentDidMount() {
    this.setState({
      list: this.props.list,
      displayList: this.props.list,
    })
  }

  componentWillReceiveProps(props) {
    this.setState({
      list: props.list,
      displayList: props.list.filter(item => item.name.indexOf(this.state.search) >= 0),
    })
  }

  render() {
    return (
      <Card
        title='My Applications'
        bordered={true}
        bodyStyle={{ padding: '0 10px' , minHeight: 400 }}
        extra={
          <div style={{ display: 'flex' }}>
            <Input.Search onChange={ e => {
              this.setState({
                search: e.target.value,
                displayList: this.state.list.filter(item => item.name.indexOf(this.state.search) >= 0),
              })
            }}/>
            <Tooltip title='Create New Application'>
              <Button
                icon='plus'
                shape='circle'
                type='primary'
                style={{ marginLeft: 16 }}
                onClick={() => {
                  this.props.onCreate()
                }}
              />
            </Tooltip>
          </div>
        }
      >
        <List
          dataSource={ this.state.displayList }
          renderItem={ item => (
            <List.Item
              actions={[
                <Tooltip title='edit'>
                  <Button icon='edit' size='small' onClick={ () => {
                    this.props.onUpdate(item)
                  }}/>
                </Tooltip>,
                <Tooltip title='delete'>
                  <Button icon='delete' size='small' onClick={ () => {
                    this.props.onDelete(item)
                  }}/>
                </Tooltip>,
              ]}
            >
              <List.Item.Meta
                title={item.name}
                description={
                  <div>
                    <div>{item.description}</div>
                    <div>App ID: {item.app_id}</div>
                    <div>
                      App Secret: {this.state.showSecret ? item.secret : '<This secret has been hidden>'}
                      <Button
                        icon={ this.state.showSecret ? 'eye-invisible' : 'eye' }
                        size='small'
                        style={{ border: 'none', marginLeft: 6 }}
                        onClick={ () => {
                          this.setState({ showSecret: !this.state.showSecret })
                        }}
                      />
                    </div>
                    <div>OIDC redirect URI: {item.oidc_redirect_uri || '<not set yet>'}</div>
                    <div>
                      update at: {moment(item.create_at).format('YYYY-MM-DD HH:mm:ss')}
                    </div>
                  </div>
                }
              />
            </List.Item>
          )}
        />
      </Card>
    )
  }
}

export default Index

