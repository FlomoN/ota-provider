import React, { useState } from "react";
import "./Main.sass";
import "spectre.css/dist/spectre-icons.css";
import { post } from "../util/fetchHelper";

export default function Main({ repos, refetch }) {
  console.log(repos);

  const [repo, setRepo] = useState("");
  const [device, setDevice] = useState("");

  const [adding, setAdding] = useState(false);
  const [loading, setLoading] = useState(false);

  const deleteHandler = async (key) => {
    const res = await post("/remove", { id: key });
    if (res.status === 200) {
      refetch();
    }
  };

  const addHandler = async () => {
    if (adding) {
      setLoading(true);
      const res = await post("/add", {
        Repo: repo,
        Device: device,
      });
      if (res.status === 200) {
        setLoading(false);
        setAdding(false);
        refetch();
      }
    } else {
      setAdding(true);
    }
  };

  const renderTracked = () => {
    return repos.map((elem, index) => (
      <div className="card" key={index}>
        <div className="columns">
          <div className="column col-4">
            <span className="repo">{elem.Repo}</span>
          </div>
          <div className="column col-4">
            <span className="device">{elem.Device}</span>
          </div>
          <div className="column col-2">
            <span className="version">{elem.Version}</span>
          </div>
          <div className="column col-2 right">
            <i
              className="icon icon-delete"
              onClick={() => deleteHandler(index)}
            ></i>
          </div>
        </div>
      </div>
    ));
  };

  return (
    <div className="container grid-lg">
      <h1>Tracked Repositories</h1>
      {renderTracked()}

      {adding && (
        <div className="card addCard">
          <div className="columns">
            <div className="column col-6">
              <input
                className="form-input"
                type="text"
                placeholder="Repository"
                value={repo}
                onChange={(e) => setRepo(e.target.value)}
              />
            </div>
            <div className="column col-6">
              <input
                className="form-input"
                type="text"
                placeholder="Device"
                value={device}
                onChange={(e) => setDevice(e.target.value)}
              />
            </div>
          </div>
        </div>
      )}

      <button
        className={
          "btn addElement " + (!adding ? "btn-primary" : "btn-success")
        }
        onClick={addHandler}
      >
        {!loading && (
          <i className={"icon " + (adding ? "icon-check" : "icon-plus")}></i>
        )}
        {loading && <div className="loading" />}
      </button>
    </div>
  );
}
