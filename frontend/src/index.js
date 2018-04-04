import dva from 'dva';
import { message } from 'antd'
import './index.css';

let lastError = null;

// 1. Initialize
const app = dva({
  onError (error) {
    if (lastError !== error.message) {
      message.error(error.message,3,(_) => lastError=null);
    }
    lastError = error.message
  },
});

// 2. Plugins
// app.use({});

// 3. Model
app.model(require('./models/index'));
app.model(require('./models/history'));
app.model(require('./models/app'));
app.model(require('./models/login'));
app.model(require('./models/login_cb'));

// 4. Router
app.router(require('./router'));

// 5. Start
app.start('#root');
