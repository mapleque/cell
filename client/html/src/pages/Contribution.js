import React, { Component } from 'react'
import { Typography } from 'antd'
import SyntaxHighlighter from 'react-syntax-highlighter'
import { docco } from 'react-syntax-highlighter/dist/esm/styles/hljs'
import Contributor from './components/Contributor'
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
        Our target is make a{' '}
        <Text mark>purer</Text>,{' '}
        <Text mark>securer</Text>, and {' '}
        <Text mark>more convenient</Text> user authorization service{' '}.
        <Title level={3}>Pull Request</Title>
        <Paragraph>
          Commit a pull request if there are any improvements.
        </Paragraph>
        <Paragraph>
          You can run in development enviroment by execute the commands:
          <ul>
            <li>
              Run server:
              <SyntaxHighlighter language='shell' style={docco} className='code'>
                {'make run-server-dev'}
              </SyntaxHighlighter>
            </li>
            <li>
              Run client:
              <SyntaxHighlighter language='shell' style={docco} className='code'>
                {'make run-html-client-dev'}
              </SyntaxHighlighter>
            </li>
          </ul>
        </Paragraph>
        <Title level={3}>Issues</Title>
        <Paragraph>
          If you have any problem, please{' '}
          <a href='https://github.com/mapleque/cell/issues/new'>create an issue</a>{' '}
          on github.
        </Paragraph>
        <Title level={3}>Contributors</Title>
        <Paragraph>
          The contributors here are grateful (listed in no particular order):
        </Paragraph>
        <Paragraph>
          <Contributor name='mapleque'/>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

