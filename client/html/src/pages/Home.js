import React, { Component } from 'react'
import { Typography } from 'antd'
const { Title, Text, Paragraph } = Typography

class Home extends Component {
  constructor (props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Typography style={{ margin: 24 }}>
        <Title level={2}>Welcome to <Text >Mapleque Cell</Text>! </Title>
        <Paragraph style={{ fontSize: 20 }}>
          This is a pure user authorization solution. We support{' '}
          <Text mark>Kerberos</Text>
          ,{' '}
          <Text mark>Oauth2.0</Text>
          ,{' '}
          <Text mark>OIDC</Text>
          , etc.
        </Paragraph>
        <Title level={3}>Why?</Title>
        <Paragraph style={{ fontSize: 20 }}>
          <ul>
            <li>Most wellknow user authorization services are bind with there busyness.</li>
            <li>They are not open source, we use it just build on trust.</li>
            <li>There are usually many complex scene to deal with.</li>
          </ul>
        </Paragraph>
        <Title level={3}>What we do?</Title>
        <Paragraph style={{ fontSize: 20 }}>
          <ul>
            <li>Providing an online user authorization service.</li>
            <li>Release for privatiztion deployment.</li>
            <li>Open sorce on Github</li>
          </ul>
        </Paragraph>
        <Title level={3}>What's more?</Title>
        <Paragraph style={{ fontSize: 20 }}>
          <ul>
            <li><a href='/document'>Using our public service</a></li>
            <li><a href='/deployment'>Privatization Deployment</a></li>
            <li><a href='/contribution'>Contribution</a></li>
          </ul>
        </Paragraph>
      </Typography>
    )
  }
}

export default Home
