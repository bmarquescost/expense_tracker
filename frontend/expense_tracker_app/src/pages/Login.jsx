// frontend/src/pages/Login.js
import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const Login = ({ setAuthenticated }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = () => {
    if (username && password) {
      axios.post(
        'http://localhost:8080/login-user', {
            username: username,
            password: password
        }).then(
          function (response) {
            const statusCode = response.status;
            if (statusCode == 200 || statusCode == 201) {
              setAuthenticated(true);
            } else {
              throw { message: response.message, statusCode }
            }
        }).catch (function (error) {
          console.error('Login error:', error);
          navigate("/register")
        });
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
    </div>
  );
};

export default Login;