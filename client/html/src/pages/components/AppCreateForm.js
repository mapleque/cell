import React, { Component } from 'react'
import { Form, Button, Input } from 'antd'

class Index extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  handleSubmit(e) {
    e.preventDefault()
    const { validateFields } = this.props.form
    validateFields((err, values) => {
      if (!err) {
        this.props.onSubmit(values)
      }
    })
  }

  render() {
    const { getFieldDecorator } = this.props.form
    return (
      <Form onSubmit={this.handleSubmit.bind(this)}>
        <h2>Create a New Application:</h2>
        <Form.Item label='Name' extra='Which name you want to show your customer.'>
          {getFieldDecorator('name', {
            rules: [{ required: true, message: 'Please input application name!'}]
          })(<Input/>)}
        </Form.Item>
        <Form.Item label='Description' extra='Whick could show your customer more details.'>
          {getFieldDecorator('description', {
          })(<Input.TextArea autosize={{ minRows: 2 }}/>)}
        </Form.Item>
        <Form.Item label='OIDC Redirect URI' extra='When you want to use OIDC service, setting it.'>
          {getFieldDecorator('oidc_redirect_uri', {
            rules: [{ type: 'url', message: 'Please input the correct URI!'}]
          })(<Input/>)}
        </Form.Item>
        <Form.Item style={{ display: 'flex', flexDirection: 'row-reverse'}}>
          <Button type='primary' htmlType='submit'>
            Create
          </Button>
          <Button style={{ marginLeft: 16 }} onClick={() => {
            this.props.onCancel()
          }}>
            Cancel
          </Button>
        </Form.Item>
      </Form>
    )
  }
}

export default Form.create({ name: 'app_create' })(Index)

