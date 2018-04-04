import React from 'react';
import { connect } from 'dva';
import {routerRedux} from 'dva/router'
import PropTypes from 'prop-types'
import { Button, Row, Col } from 'antd';
import URL from 'url-parse'
import * as qs from 'query-string';
import Page from '../components/page/page'

const LoginPage = ({state, dispatch,location}) => {
	const onClick = () => {
	  	const query = qs.parse(location.search);
	  	let url = null
	  	if ('url' in query) {
	    		let urlParse = new URL(query.url);
			url = urlParse.hash
			if(url.length > 0) {
				url = url.substr(1)
			}
	  	} 
	
	  	dispatch({
        		type: 'login/login',
			payload: url
      	})	

	}
	
	const onQyClick = () => {
	    let url = new URL(window.location.href);
	    url.set('hash', "#/login_cb" + location.search);
	    window.location.assign("/api/v1/login_url" + "?" + qs.stringify({
	      url: url.toString(),
	    }));
	}
	
	return (
		<Row type="flex" justify="center" align="middle">	      
		  <Col span={1}>
			<Page><div/></Page>
		  </Col>
	      <Col span={3}>
		    <Button type="primary" icon="login" size="large"
				onClick={onClick}
			>匿名登陆</Button>
		  </Col>
	      <Col span={3}>
		    <Button type="primary" icon="login" size="large"
				onClick={onQyClick}
			>企业微信扫码登陆</Button>
		  </Col>
	    </Row>
	)
}

export default connect(({login}) => ({state: login}))(LoginPage)
