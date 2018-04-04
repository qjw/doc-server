import React from 'react';
import { connect } from 'dva';
import {routerRedux} from 'dva/router'
import PropTypes from 'prop-types'
import { Collapse, List, Row, Col, Button } from 'antd';
import * as qs from 'query-string';
const Panel = Collapse.Panel;
const ButtonGroup = Button.Group;


const IndexPage = ({state, dispatch}) => {
	var {description,title,groups} = state
	if(!groups) groups = []
	
	const onClickHistory = (item, obj) => () => {
	  dispatch(routerRedux.push({
        pathname: "/history",
		search: "?" + qs.stringify({
			file: item.spec,
			name: item.name,
			production: obj.name,
		}),
      }))
	}
	
	const onClick = (item) => () => {
		const url = "/api/v1/file?" + qs.stringify({
			file: item.spec,
		})
		window.open("/swagger/index.html?" + qs.stringify({
			url: url
		}) , '_blank');
	}
	
	return (
	  <Collapse style={{padding:10}}>
		{groups.map(function(obj, i){
	        return (
				<Panel 
				  header={(<div><h2><strong>{obj.name}</strong></h2>{obj.description}</div>)}
				  key={i}>
				  <div>
					  <List
						bordered
					    itemLayout="horizontal"
					    dataSource={obj.projects}
					    renderItem={item => (
					      <List.Item style={{padding: 10, paddingLeft: 40, background: "#ffffef"}}>
							<Row style={{width: "100%"}}>
						      <Col span={12}>
								<div>
								  <h2><strong>{item.name}</strong></h2>
								  <h4>{item.description}</h4>
								</div>
							  </Col>
						      <Col span={12}>
							 	<ButtonGroup style={{margin:10}}>
							      <Button type="primary" size="large" onClick={onClick(item)}>查看文档</Button>
							      <Button type="ghost" size="large" onClick={onClickHistory(item, obj)}>查看历史</Button>
							    </ButtonGroup>
							  </Col>
						    </Row>
					      </List.Item>
					    )}
					  />
				  </div>
			    </Panel>
			);
	    })}
	  </Collapse>
	)
}

IndexPage.propTypes = {
	index: PropTypes.object,
};

export default connect(({index}) => ({state: index}))(IndexPage)
