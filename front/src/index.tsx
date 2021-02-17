import React from 'react';
import ReactDOM from 'react-dom';
import RoomIndex from './pages/room/Index';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter as Router, Route } from "react-router-dom";

ReactDOM.render(
  <Router>
    <RoomIndex>
        <Route exact path="/" component={RoomIndex}></Route>
        <Route exact path="/rooms" component={RoomIndex}></Route>
    </RoomIndex>
  </Router>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
