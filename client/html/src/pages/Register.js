import React, { Component } from 'react'
import { Alert, Form, Icon, Input, Button, Checkbox, Row, Col } from 'antd'
import { MD5 } from 'crypto-js'
import { Link } from 'react-router-dom'

const md5 = data => MD5(`${process.env.REACT_APP_SALT}${data}`).toString()

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {
      confirmDirty: false,
      captchaLimit: 0,
      errorMessage: null,
      successMessage: null,
    }
    this.captchaLimit = null
  }

  handleSubmit(e) {
    e.preventDefault()
    this.setState({errorMessage: null})
    const { validateFields } = this.props.form
    validateFields(async (err, values) => {
      if (!err) {
        await this.register(values)
      }
    })
  }

  handleCaptcha() {
    if (this.state.captchaLimit > 0) return
    const { validateFields } = this.props.form
    validateFields(['username'], async (err, values) => {
      if (!err) {
        if (this.state.captchaLimit) return
        this.setState({ captchaLimit: 60 }, () => {
          const captchaInterval = setInterval(() => {
            const { captchaLimit } = this.state
            if (captchaLimit <= 0) {
              clearInterval(this.state.captchaInterval)
              this.setState({ captchaInterval: null})
            } else {
              this.setState({ captchaLimit: captchaLimit - 1 })
            }
          }, 1000)
          this.setState({ captchaInterval }, () => {
            this.sendCaptcha(values)
          })
        })
      }
    })
  }

  async sendCaptcha({ username }) {
    this.setState({
      errorMessage: null,
      successMessage: null,
    })
    try {
      await this.props.request('/user/captcha', { username })
      this.setState({successMessage: 'The captcha has been sent to your Email, please check!'})
    } catch (e) {
      this.setState({errorMessage: e.message})
    }
  }

  async register({ username, password, captcha }) {
    this.setState({
      errorMessage: null,
      successMessage: null,
    })
    try {
      await this.props.request('/user/register', { username, password: md5(password), captcha })
      this.setState({successMessage: <span>Your account has been created, please <Link to='/login'>login</Link>!</span>})
    } catch (e) {
      this.setState({errorMessage: e.message})
    }
  }

  render() {
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 },
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 },
      },
    }
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0,
        },
        sm: {
          span: 16,
          offset: 8,
        },
      },
    }
    const { getFieldDecorator } = this.props.form
    return (
      <Form {...formItemLayout} onSubmit={this.handleSubmit.bind(this)} style={{ margin: '70px auto', width: 500, position: 'relative' }}>
        <h2 style={{ textAlign: 'center', marginBottom: 24 }}>
          You are now registering in Mapleque Cell:
        </h2>
        {
          this.state.successMessage !== null ? <Alert
            message={this.state.successMessage}
            type='success'
            style={{ position: 'absolute', top: -50, width: 500 }}
            closable
          /> : ''
        }
        {
          this.state.errorMessage !== null ? <Alert
            message={this.state.errorMessage}
            type='error'
            style={{ position: 'absolute', top: -50, width: 500 }}
            closable
          /> : ''
        }
        <Form.Item label='E-mail'>
          {getFieldDecorator('username', {
            rules: [
              { type: 'email', message: 'The input is not valid Email!' },
              { required: true, message: 'Please input your Email!' },
            ]
          })(
            <Input prefix={<Icon type='user' style={{ fontSize: 13 }} />} placeholder='Email' />
          )}
        </Form.Item>
        <Form.Item label='Password' hasFeedback>
          {getFieldDecorator('password', {
            rules: [
              { required: true, message: 'Please input your Password!' },
              { validator: (rule, value, callback) => {
                const { validateFields } = this.props.form
                if (value && this.state.confirmDirty) {
                  validateFields(['confirm'], { force: true })
                }
                callback()
              }}
            ]
          })(
            <Input.Password prefix={<Icon type='lock' style={{ fontSize: 13 }} />} placeholder='Password' />
          )}
        </Form.Item>
        <Form.Item label='Confirm Password' hasFeedback>
          {getFieldDecorator('confirm', {
            rules: [
              { required: true, message: 'Please confirm your Password!' },
              { validator: (rule, value, callback) => {
                const { getFieldValue } = this.props.form
                if (value && value !== getFieldValue('password')) {
                  callback('Two passwords that you enter is inconsistent!')
                } else {
                  callback()
                }
              }}
            ]
          })(
            <Input.Password
              prefix={<Icon type='lock' style={{ fontSize: 13 }} />}
              placeholder='Confirm Password'
              onBlur={ e => {
                const { value } = e.target
                this.setState({ confirmDirty: this.state.confirmDirty || !!value })
              }}
            />
          )}
        </Form.Item>
        <Form.Item label="Captcha" extra="We will send a captcha to your E-mail.">
          <Row gutter={8}>
            <Col span={12}>
              {getFieldDecorator('captcha', {
                rules: [
                  { required: true, message: 'Please input the captcha you got!' }
                ],
              })(<Input />)}
            </Col>
            <Col span={12}>
              <Button onClick={this.handleCaptcha.bind(this)} disabled={this.state.captchaLimit > 0}>
                Get captcha
              </Button>
              { this.state.captchaLimit > 0 ? ` ${this.state.captchaLimit}s` : '' }
            </Col>
          </Row>
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          {getFieldDecorator('agreement', {
            rules: [
              { required: true, message: 'Please read the agreement!' },
            ],
            valuePropName: 'checked',
          })(
            <Checkbox>I have read the <Link to='/agreement'>agreement</Link></Checkbox>
          )}
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          <Button type='primary' htmlType='submit'>
            Register
          </Button>
          <span style={{ float: 'right' }}>Already have account? <Link to='/login'>login now!</Link></span>
        </Form.Item>
      </Form>
    )
  }
}

export default Form.create({ name: 'register' })(Index)
