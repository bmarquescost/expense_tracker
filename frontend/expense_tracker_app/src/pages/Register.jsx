import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const Register = ({ setAuthenticated }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleRegister = () => {
    if (username && password) {
        axios.post(
            'http://localhost:8080/register-user', {
                username: username,
                password: password
        }).then(function (response) {
            const statusCode = response.status;
            if (statusCode == 200 || statusCode == 201) {
                console.log('Register successful:', response.data);
                setAuthenticated(true);
                navigate("/")
            } else {
                throw { message: response.message, statusCode }
            }
        }).catch(function (error) {
            console.error('Register error:', error);
            setError("Registration failed. Please try again.");
        })
    }
    
    else {
        setError("Both fields are required for registering user.")
    }
  };

  return (
    <div>
      <h2>Register</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <input type="text" placeholder="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleRegister}>Register</button>
    </div>
  );
};

export default Register;