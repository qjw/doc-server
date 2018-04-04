import { specApi } from '../services/index'

export default {

  namespace: 'index',

  state: {
  },

  subscriptions: {
    setup ({ dispatch, history }) {
      history.listen(({ pathname }) => {
        if (pathname == "/") {
          dispatch({ type: 'query'})
        }
      })
    },
  },

  effects: {
    * query ({}, { call, put }) {
      const {data} = yield call(specApi)
      yield put({
        type: 'querySuccess',
        data: data,
      })
    },
  },

  reducers: {
    querySuccess (state, { data }) {
      return {
        ...state,
        ...data,
      }
    },
  },
}