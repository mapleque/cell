import React, { Component } from 'react'
import { Typography, Form, Input, Button } from 'antd'
const { Title, Paragraph, Text } = Typography

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Typography>
        <Title level={2}>Please Confirm:</Title>
        <Paragraph>
          Your are deleting the application{' '}
          <Text strong>
            {this.props.app.name}
          </Text>
          {' '}which app id is{' '}
          <Text strong>
            {this.props.app.app_id}
          </Text>. Once you do that, it can not be recoverd.
        </Paragraph>
        <Paragraph>
          Please enter the app name:{' '}
          <Form.Item
            validateStatus={this.state.inputStatus}
            style={{ display: 'inline-block', margin: '-10px 0 0' }}
          >
            <Input style={{ width: 'auto', marginLeft: 16 }} onChange={e => {
              if (e.target.value !== this.props.app.name) {
                this.setState({ inputStatus: 'error' })
              } else {
                this.setState({ inputStatus: 'success' })
              }
            }}/>
          </Form.Item>
        </Paragraph>
        <Paragraph style={{ display: 'flex', flexDirection: 'row-reverse' }}>
          <Button style={{ margin: '0 8px' }} onClick={() => {
            this.props.onCancel()
          }}>Cancel</Button>
          <Button
            style={{ margin: '0 8px' }}
            type='danger'
            disabled={this.state.inputStatus !== 'success'}
            onClick={() => {
              this.props.onCommit()
            }}
          >Delete it!</Button>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

