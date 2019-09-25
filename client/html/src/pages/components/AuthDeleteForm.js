import React, { Component } from 'react'
import { Typography, Button } from 'antd'
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
          Your are deleting the authorization to{' '}
          <Text strong>
            {this.props.auth.name}
          </Text>
          . Once you do that, you be logout from that application immediately.
        </Paragraph>
        <Paragraph style={{ display: 'flex', flexDirection: 'row-reverse' }}>
          <Button style={{ margin: '0 8px' }} onClick={() => {
            this.props.onCancel()
          }}>Cancel</Button>
          <Button
            style={{ margin: '0 8px' }}
            type='primary'
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

