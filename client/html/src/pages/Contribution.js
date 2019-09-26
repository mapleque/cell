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
        <Title level={2}>Contribution</Title>
        <Paragraph style={{ fontSize: 20 }}>
        We wish your contribution.
        </Paragraph>
        Our target is make a
        <Text mark>purer</Text>,{' '}
        <Text mark>securer</Text>, and {' '}
        <Text mark>more convenient</Text> user authorization service{' '}.
        <Title level={3}>Pull Request</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
        <Title level={3}>Issues</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
        <Title level={3}>Contributors</Title>
        <Paragraph style={{ fontSize: 20 }}>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

