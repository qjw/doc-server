import React, { Component } from 'react'
import PropTypes from 'prop-types'
import classnames from 'classnames'
import styles from './page.less'

export default class Page extends Component {
  render () {
    const {
      className, children,
    } = this.props
	
    return (
      <div
        className={classnames(className, {
          [styles.contentInner]: true,
        })}
      >
        {children}
      </div>
    )
  }
}


Page.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node,
}
