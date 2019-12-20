import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';


const [id,setId] = React.useState(null);
const [passwd,setPasswd] = React.useState(null);
const [message,setMessage] = React.useState(null);

const login = ()=>{
  axios.request({
    url:'/login',
    method:'get',
    params:{
      id:id,
      passwd:passwd
    }
  }).then(res =>{
    console.log(res);
    if(res.data){
      setMessage("登录成功");
    }
    else{
      setMessage("账号或密码错误");
    };
  });
};

class App extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>

        <div>
            <h2>用户登录</h2>
            <p><span>用户名：</span>
              <input type="userid" onInput={(e) => {setId(e.target.value);}}/>
            </p>
            <p><span>密  码：</span>
              <input type="password" onInput={(e) => {setPasswd(e.target.value);}}/>
            </p>
            <p><button type="submit" value="登录" onClick={login}/></p>
        </div>
        <span>{message}</span>
      </div>
    );
  }
}

export default App;
