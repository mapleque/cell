import React from 'react'
import { Router } from '@/components/router'
import { Server } from '@/components/server'
import config from './router-config'
import './App.css'

function App() {
  return (
    <Server
      servers={[
        {
          path_rule: true,
          before: req => {
            req.path = `/api${req.path}`
            console.log(`%c[send] ${req.path} ${JSON.stringify(req.params)}`, 'color: blue')
            return req
          },
          after: req => {
            try {
              if (req.err) {
                throw req.err
              }
              if (req.resp.status !== 200) {
                throw new Error(`HTTP Status ${req.resp.status}`)
              }
              if (typeof req.resp.data === 'undefined') {
                throw new Error('HTTP Response Empty')
              }
              if (req.resp.data.status !== 0) {
                throw new Error(`Error: (${req.resp.data.status}) ${req.resp.data.data || req.resp.data.message}`)
              }
              console.log(`%c[recv] ${req.path} ${JSON.stringify(req.resp.data)}`, 'color: blue')
              return req.resp.data.data
            } catch(err) {
              console.log(`%c[recv] ${req.path} ${err.message}`, 'color: red')
              throw err
            }
          },
        }
      ]}
    >
      <Router config={config}/>
    </Server>
  )
}

export default App
