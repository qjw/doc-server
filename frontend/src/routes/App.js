import React from 'react';
import { connect } from 'dva';
import {routerRedux} from 'dva/router'
import PropTypes from 'prop-types'
import { Helmet } from 'react-helmet'
import { Layout, Menu, Breadcrumb, Icon } from 'antd';
const { Header, Content, Footer, Row, Col } = Layout;
import Page from '../components/page/page'

const App = ({ state, children, location, dispatch}) => {
  const {description} = state
  return (
	<div>
	  <Helmet>
        <title>{description && description.title}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      </Helmet>
	 <Layout>
	    <Header className="header">
			<span>
			<span style={{fontSize: 32,color: "#cccccc"}}>{description && description.title}</span>
			<span style={{fontSize: 16,color: "#cccccc"}}> {description && description.description}</span>
			</span>
	    </Header>
		<Content>
			<Page>
				{children}
	        </Page>
	    </Content>
	    <Footer style={{ textAlign: 'center' }}>
	     	CopyRight ©2018 Created by {state.description && state.description.company}
	    </Footer>
	  </Layout>
	</div>
  )
}

App.propTypes = {
  children: PropTypes.element.isRequired,
  dispatch: PropTypes.func,
}


export default connect(({app}) => ({state: app}))(App)