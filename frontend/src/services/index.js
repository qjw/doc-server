import request from '../utils/request';

export async function specApi (params) {
	return request({
        url: "/api/v1/spec",
        method: 'get',
      })
}

export async function historyApi (params) {
  return request({
		url: "/api/v1/history",
	    method: 'get',
	    data: params,
	  })
}

export async function descriptionApi (params) {
  return request({
		url: "/api/v1/description",
	    method: 'get',
	    data: params,
	  })
}

export async function loginApi (params) {
  return request({
		url: "/api/v1/login",
	    method: 'get',
	    data: params,
	  })
}

export async function loginQyApi (params) {
  return request({
		url: "/api/v1/login_qy",
	    method: 'get',
	    data: params,
	  })
}