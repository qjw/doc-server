import React from 'react';
import { connect } from 'dva';
import {routerRedux} from 'dva/router'
import PropTypes from 'prop-types'
import * as qs from 'query-string';
import { List,Row, Col, Button, Badge } from 'antd';

const HistoryPage = ({state, dispatch}) => {
	const onClick = (item) => () => {
		const url = "/api/v1/file?" + qs.stringify({
			file: state.file,
			commit: item.hash,
		})
 		window.open("/swagger/index.html?" + qs.stringify({
			url: url
		}) , '_blank');
	}
	
	return (
		<div>
		<div style={{padding: "15px"}}>
			{state.data &&<Badge offset={[0,10]} count={state.data.length}>
			<span style={{ fontSize: 32, color: '#08c' }}>{state.production}{" - "}{state.name}</span>
			</Badge>}
		</div>
	    <Row>
		<Col span={0}/>
		<Col span={24}>
		<List bordered
		  size="large"
		  style={{marginLeft:10,marginRight:10}}
	      dataSource={state.data}
	      renderItem={item => (
			<List.Item style={{background: "#ffffef"}}>
			<Row style={{width: "100%"}}>
		      <Col span={11}>
				<div>
				  <h2><strong>{item.name}</strong> {"(" + item.email + " / " + item.time + ")"}</h2>
				  <h4>{item.hash}</h4>
				</div>
			  </Col>
		      <Col span={11}>
				<span>{item.log}</span>
			  </Col>
		      <Col span={2}>
				<Button type="primary" size="large" onClick={onClick(item)}> 查看文档</Button>
			  </Col>
		    </Row>
			</List.Item>
		  )}
	    />
		</Col>
		<Col span={0}/>
		</Row>
		</div>
	)
}

HistoryPage.propTypes = {
	history: PropTypes.object,
};

export default connect(({history}) => ({state: history}))(HistoryPage)
