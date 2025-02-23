// frontend/src/App.js
import React, { useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./index.css";
import Login from "../../../expense_tracker/frontend/expense_tracker_app/src/pages/Login";
import Home from "../../../expense_tracker/frontend/expense_tracker_app/src/pages/Home";
import Register from "../../../expense_tracker/frontend/expense_tracker_app/src/pages/Register";

const App = () => {
  const [authenticated, setAuthenticated] = useState(false);

  return (
    <Router>
      <Routes>
        <Route path="/" element={authenticated ? <Home /> : <Login setAuthenticated={setAuthenticated} />} />
        <Route path="/register" element={<Register setAuthenticated={setAuthenticated}/>} />
      </Routes>
    </Router>
  );
};

export default App;
