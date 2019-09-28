import React, { Component, Fragment } from 'react'
import { Alert, Form, Icon, Input, Button, Checkbox } from 'antd'
import { MD5, AES, enc, pad, mode } from 'crypto-js'
import { Link } from 'react-router-dom'

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

const md5 = data => MD5(`${process.env.REACT_APP_PASSWORD_SALT}${data}`).toString()

const timestamp = () => parseInt(new Date().valueOf() / 1000, 10)

class Index extends Component {
  constructor (props) {
    super(props)
    this.state = {
      errorMessage: null
    }
  }

  handleSubmit(e) {
    e.preventDefault()
    this.setState({errorMessage: null})
    const { validateFields } = this.props.form
    validateFields(async (err, values) => {
      if (!err) {
        await this.login(values)
      }
    })
  }

  async login({ username, password, remember }) {
    try {
      const { request } = this.props
      // using kerberos here
      const { ctsk, tgt } = await request('/kerberos/auth', { username }, { throw_error: true })
      const userKey = md5(password)

      const ctKey = decrypt(ctsk, userKey)

      const authenticator = { username, timestamp: timestamp() }

      const { cssk, st } = await request('/kerberos/grant', {
        tgt,
        app_id: 'cell',
        authenticator: encrypt(authenticator, ctKey),
      }, { throw_error: true })

      const csKey = decrypt(cssk, ctKey)

      const token = await request('/user/login', {
        st,
        authenticator: encrypt(authenticator, csKey),
        remember,
      }, { throw_error: true })

      // Here if your want to deal no cookie permission scene,
      // your can store the token and put it in HTTP Header:
      //     `Authenticate: Bearer ${token}`
      //
      // Don't forget to manage the expiration time in client.
      localStorage.setItem('_cell_session', JSON.stringify({ token, timestamp: timestamp() }))
      localStorage.setItem('_cell_authenticator', JSON.stringify(authenticator))

      const { search } = window.location
      if (search.length > 1) {
        const redirect = search.substring(1).split('&').find(query => query.indexOf('redirect=') === 0)
        if (redirect) {
          window.location.replace(decodeURIComponent(redirect.substring(9)))
          return
        }
      }
      window.location.replace('/dashboard')
    } catch (e) {
      console.error(e)
      this.setState({errorMessage: 'Invalid username or password'})
    }
  }

  render() {
    const { getFieldDecorator } = this.props.form
    return (
      <Fragment>
        <h2 style={{ top: 125, width: 300, left: '50%', marginLeft: -150, textAlign: 'center', position: 'absolute' }}>
          Please login with your account
        </h2>
        <Form
          onSubmit={this.handleSubmit.bind(this)}
          style={{ margin: '150px auto', width: 300, position: 'relative', top: '50%' }}
        >
          {
            this.state.errorMessage !== null ? <Alert
              message={this.state.errorMessage}
              type='error'
              style={{ position: 'absolute', top: -50, width: 300 }}
              closable
            /> : ''
          }
          <Form.Item>
            {getFieldDecorator('username', {
              rules: [{ required: true, message: 'Please input your username!' }]
            })(
              <Input prefix={<Icon type='user' style={{ fontSize: 13 }} />} placeholder='Username' />
            )}
          </Form.Item>
          <Form.Item>
            {getFieldDecorator('password', {
              rules: [{ required: true, message: 'Please input your Password!' }]
            })(
              <Input prefix={<Icon type='lock' style={{ fontSize: 13 }} />} type='password' placeholder='Password' />
            )}
          </Form.Item>
          <Form.Item>
            {getFieldDecorator('remember', {
              valuePropName: 'checked',
              initialValue: true
            })(
              <Checkbox>Remember me</Checkbox>
            )}
            <Link to='/forgot' style={{ float: 'right' }}>Forgot password</Link>
            <Button type='primary' htmlType='submit' style={{ width: '100%' }}>
              Log in
            </Button>
            Or <Link to='/register'>register now!</Link>
          </Form.Item>
        </Form>
      </Fragment>
    )
  }
}

export default Form.create({ name: 'login' })(Index)

