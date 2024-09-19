import axios from "axios";
import React, { useState } from "react";

const App = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handlesubmit = async () => {
    const url = `http://localhost:8080/login`;
    try {
      const response = await axios.post(url, {
        username: username,
        password: password,
      });
      // console.log(username, password);

      if (response.status === 201) {
        alert(`Variables passed`);
      }
    } catch (error) {
      console.log(`error pasring variables ${error}`);
    }
  };

  return (
    <div>
      <input
        placeholder={"username"}
        type={"text"}
        name="username"
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        placeholder={"password"}
        type={"password"}
        name="password"
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handlesubmit}>Login</button>
    </div>
  );
};

export default App;
