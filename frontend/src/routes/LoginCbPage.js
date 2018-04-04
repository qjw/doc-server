import React from 'react'
import PropTypes from 'prop-types'
import {connect} from 'dva'


const LoginCB = ({location, dispatch}) => {
  return (
    <div>正在登陆 ...</div>
  )
}

export default connect()(LoginCB)