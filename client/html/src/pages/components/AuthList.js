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
        title='My Authorizations'
        bordered={true}
        bodyStyle={{ padding: '0 10px', minHeight: 400 }}
        extra={
          <div style={{ display: 'flex', justifyContent: 'space-between' }}>
            <Input.Search onChange={ e => {
              this.setState({
                search: e.target.value,
                displayList: this.state.list.filter(item => item.name.indexOf(this.state.search) >= 0),
              })
            }}/>
          </div>
        }
      >
        <List
          dataSource={ this.state.displayList }
          renderItem={ item => (
            <List.Item
              actions={[
                <Tooltip title='delete'>
                  <Button icon='delete' size='small' onClick={ () => {
                    this.props.onDelete(item)
                  }}/>
                </Tooltip>,
              ]}
            >
              <List.Item.Meta
                title={
                  <span>
                    <span style={{ fontWeight: 'bold' }}>{item.name}</span>
                    <span style={{ marginLeft: 16, color: 'rgba(0, 0, 0, 0.45)' }}>
                          granted at: {moment(item.create_at).format('YYYY-MM-DD HH:mm:ss')}
                    </span>
                  </span>
                }
                description={item.description}
              />
            </List.Item>
          )}
        />
      </Card>
    )
  }
}

export default Index

