import React, { Component } from 'react'
import { Typography } from 'antd'
const { Title, Text, Paragraph } = Typography

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Typography style={{ margin: 24 }}>
        <Title level={2}>Deployment</Title>
        For privatization deployment, you can choose{' '}
        <Text mark>Docker</Text>,{' '}
        <Text mark>Release</Text>,{' '}
        <Text mark>Source Code</Text>, ect.
        <Title level={3}>Docker image</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
        <Title level={3}>Release package</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
        <Title level={3}>Source code</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

