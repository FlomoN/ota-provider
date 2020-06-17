import React, { useEffect, useState } from "react";
import Login from "./components/Login";
import "spectre.css";
import Main from "./components/Main";
import { get } from "./util/fetchHelper";

function App() {
  const [init, setInit] = useState(true);
  const [repos, setRepos] = useState([]);

  const fetchData = async () => {
    const res = await (await get("/data")).json();
    setInit(res.Init);
    setRepos(res.Watch);
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className="App">
      {!init && <Login done={setInit} />}
      {init && <Main repos={repos} refetch={fetchData} />}
    </div>
  );
}

export default App;
