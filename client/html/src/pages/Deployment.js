import React, { Component } from 'react'
import { Typography } from 'antd'
import SyntaxHighlighter from 'react-syntax-highlighter'
import { docco } from 'react-syntax-highlighter/dist/esm/styles/hljs'
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
        <Title level={3}>Database</Title>
        <Paragraph>
          First of all, you should prepare a database which can connect with mysql driver.
          <ul>
            <li>
              Create a database for cell, usually execute the sql:
              <SyntaxHighlighter language='sql' style={docco} className='code'>
                {'CREATE DATABASE cell;'}
              </SyntaxHighlighter>
            </li>
            <li>
              Create tables with those sql files: https://github.com/mapleque/cell/tree/master/sql.
            </li>
            <li>
              Create use and grant to database cell:
              <SyntaxHighlighter language='sql' style={docco} className='code'>
                {`CREATE USER 'cell'@'cell.service.com' IDENTIFIED BY 'password';
GRANT ALL ON cell.* TO 'cell'@'cell.service.com';`}
              </SyntaxHighlighter>
            </li>
          </ul>
        </Paragraph>
        <Title level={3}>Docker image</Title>
        <Paragraph>
          <SyntaxHighlighter language='shell' style={docco} className='code'>
            {`docker run -d \\
  -p 0.0.0.0:80:80 \\
  --env DB_DSN='cell:password@tpc(db.service.com:3306)/cell?charset=utf8mb4&parseTime=true&loc=Local' \\
  --env KERBEROS_TGS_SECRET_KEY=a_random_token_with_32_charactor \\
  --env KERBEROS_APP_SECRET_KEY=b_random_token_with_32_charactor \\
  --env REACT_APP_PASSWORD_SALT= a_password_salt \\
  --env MAIL_USERNAME= \\
  --env MAIL_PASSWORD= \\
  --env MAIL_HOST= \\
  --env MAIL_ADDRESS= \\
  --env MAIL_FROM= \\
  mapleque/cell`}
          </SyntaxHighlighter>
        </Paragraph>
        <Title level={3}>Release package</Title>
        <Paragraph>
          Deploy with a release package on your application server:
          <ul>
            <li>
              Download the release package on https://github.com/mapleque/cell/releases.
            </li>
            <li>
              Unzip the package and edit .env or use enviroment valuables while removing .env file.
            </li>
            <li>
              Execute <Text code>./cell-x</Text> (x is which paltform you selected) in the release path.
            </li>
          </ul>
        </Paragraph>
        <Title level={3}>Source code</Title>
        <Paragraph>
          Usually we use the source code for developing (or you can fork in your own namespace):
          <ul>
            <li>
              Clone the repo on github:
              <SyntaxHighlighter language='shell' style={docco} className='code'>
                {'git clone https://github.com/mapleque/cell.git'}
              </SyntaxHighlighter>
            </li>
            <li>
              Compile and install on current server:
              <SyntaxHighlighter language='shell' style={docco} className='code'>
                {'cd cell && make install'}
              </SyntaxHighlighter>
            </li>
            <li>
              Edit bin/.env or use enviroment valuables while removing bin/.env file.
            </li>
            <li>
              Run on server:
              <SyntaxHighlighter language='shell' style={docco} className='code'>
                {'make run'}
              </SyntaxHighlighter>
            </li>
          </ul>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

