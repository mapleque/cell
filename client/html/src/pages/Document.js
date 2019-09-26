import React, { Component } from 'react'
import { Typography } from 'antd'
import SyntaxHighlighter from 'react-syntax-highlighter'
import { docco } from 'react-syntax-highlighter/dist/esm/styles/hljs'
import './Document.css'

const { Title, Text, Paragraph } = Typography

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  render() {
    return (
      <Typography style={{ margin: 24 }}>
        <Title level={2}>Document</Title>
        <Paragraph style={{ fontSize: 20 }}>
          Using our public service, is a fastest and simplest way to develop your busyness system.
        </Paragraph>
        <Title level={3}>What's the steps</Title>
        <Paragraph>
          <ol>
            <li><Text mark><a href='/register'>Create an account</a></Text> and <Text mark><a href='/login'>login</a></Text>.</li>
            <li>Enter the <Text mark><a href='/dashboard'>Dashboard</a></Text> and <Text mark> create an application</Text>.</li>
            <li>Find the application settings in <Text mark>My Application</Text> list.</li>
          </ol>
        </Paragraph>
        <Title level={3}>Kerberos</Title>
        <Paragraph>
          Here we show the client processing with javascript snippets.
          <ol>
            <li>
              When user login, request to <Text mark>https://cell.mapleque.com/kerberos/auth</Text> with the param <Text mark>username</Text>.{' '}
              You will get <Text mark>Encrypt Client/TGS Session Key</Text> and <Text>Encrypt TGT</Text> from response.

              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {'const { ctsk, tgt } = await request(\'/kerberos/auth\', { username })'}
              </SyntaxHighlighter>
            </li>
            <li>
              Decrypt the <Text mark>Encrypt Client/TGS Session Key</Text> with <Text mark>User Key</Text>,{' '}
              which is an md5 hash code of user <Text mark>password</Text> combined with the salt <Text code>cell_salt</Text> in format of <Text code>cell_saltpassowrd</Text>.
              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {`const userKey = md5(password)
const ctKey = decrypt(ctsk, userKey)`}
              </SyntaxHighlighter>
            </li>
            <li>
              Build an <Text mark>Authenticator</Text> with <Text mark>username</Text> and <Text mark>timestamp</Text>
              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {'const authenticator = { username, timestamp: timestamp() }'}
              </SyntaxHighlighter>
            </li>
            <li>
              Request to <Text mark>https://cell.mapleque.com/kerberos/grant</Text> with the params:
              <ul>
                <li>
                  <Text mark>Encrypt TGT</Text>, which is from the step 1 response.
                </li>
                <li>
                  <Text mark>Encrypt Authenticator</Text>, which is the <Text mark>Authenticator</Text> encrypt with <Text mark>Client/TGS Session Key</Text>.
                </li>
                <li>
                  <Text mark>App ID</Text>, which could be found in <Text>My Applications</Text> list on the page of <a href='/dashboard'>Dashboard</a>.
                </li>
              </ul>
              You will get <Text mark>Encrypt Service Ticket</Text> and <Text mark>Encrypt Client/Server Session Key</Text> from response.
              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {`const { cssk, st } = await request('/kerberos/grant', {
  tgt,
  app_id: 'found App ID in My Applications list',
  authenticator: encrypt(authenticator, ctKey),
})`}
              </SyntaxHighlighter>
            </li>
            <li>
              Decrypt the <Text mark>Encrypt Client/Server Session Key</Text> with <Text mark>Client/TGS Session Key</Text>.
              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {'const csKey = decrypt(cssk, ctKey)'}
              </SyntaxHighlighter>
            </li>
            <li>
              Finally, request to your own server for login with the params:
              <ul>
                <li>
                  <Text mark>Encrypt Service Ticket</Text>, which is from step 4 response.
                </li>
                <li>
                  <Text mark>Encrypt Authenticator</Text>, which is the <Text mark>Aunthenticator</Text>, built in step 3,  encrypt with <Text mark>Client/Server Session Key</Text>.
                </li>
              </ul>
              <SyntaxHighlighter language='javascript' style={docco} className='code'>
                {`const token = await request('/user/login', {
  st,
  authenticator: encrypt(authenticator, csKey),
})`}
              </SyntaxHighlighter>
              To deal with the login request, your own server should validate with following steps:
              <ul>
                <li>
                  Decrypt the <Text mark>Encrypt Service Ticket</Text> with <Text mark>App Secret</Text>, which could be found in <Text>My Applications</Text> list on the page of <a href='/dashboard'>Dashboard</a>.
                </li>
                <li>
                  Check the <Text mark>expired</Text> of <Text mark>Service Ticket</Text> is valid.
                </li>
                <li>
                  Decrypt the <Text mark>Encrypt Authenticator</Text> with <Text mark>Client/Server Session Key</Text>, which comes from <Text mark>Service Ticket</Text>.
                </li>
                <li>
                  Check the <Text mark>username</Text> in <Text mark>Authenticator</Text> is same with the one in <Text mark>Service Ticket</Text>
                </li>
                <li>
                  You can also check the <Text mark>timestamp</Text> in <Text mark>Authenticator</Text> if you want.
                </li>
              </ul>
            </li>
          </ol>
          The javascript common method snippet:
          <SyntaxHighlighter language='javascript' style={docco} className='code'>
            {`import { MD5, AES, enc, pad, mode } from 'crypto-js'
const encrypt = (data, key) => AES.encrypt(
  JSON.stringify(data),
  enc.Utf8.parse(key),
  { mode: mode.ECB, padding: pad.Pkcs7 }
).toString()

const decrypt = (data, key) => JSON.parse(
  enc.Utf8.stringify(
    AES.decrypt(
      data,
      enc.Utf8.parse(key),
      { mode: mode.ECB, padding: pad.Pkcs7 }
    )
  )
)
const md5 = data => MD5(\`\${process.env.REACT_APP_SALT}\${data}\`).toString()
const timestamp = () => parseInt(new Date().valueOf() / 1000, 10)
`}
          </SyntaxHighlighter>
        </Paragraph>
        <Title level={3}>OAuth2.0</Title>
        <Paragraph>
        </Paragraph>
        <Title level={3}>OIDC</Title>
        <Paragraph>
        </Paragraph>
      </Typography>
    )
  }
}

export default Index

