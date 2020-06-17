import React, { useState } from "react";
import "./Login.sass";
import { post } from "../util/fetchHelper";

export default function Login({ done }) {
  const [ghName, setGhName] = useState("");
  const [ghPass, setGhPass] = useState("");
  const [broker, setBroker] = useState("");

  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);

  const submitHandler = async () => {
    const [host, port] = broker.split(":");
    if (!port || isNaN(port)) {
      console.error("No Port supplied");
      setError(true);
    } else {
      setError(false);
      setLoading(true);
      const res = await post("/creds", {
        Name: ghName,
        Pass: ghPass,
        MQTT: broker,
      });
      setLoading(false);
      if (res.status === 200) {
        done(true);
      }
    }
  };

  return (
    <div className="full">
      <div className="card">
        <div className="card-header">
          <div className="card-title h2">Setup</div>
          <div className="card-subtitle text-gray">
            Enter your Github Credentials and MQTT Broker
          </div>
        </div>
        <div className="card-body">
          <input
            value={ghName}
            onChange={(e) => setGhName(e.target.value)}
            className="form-input"
            type="text"
            placeholder="Github Username"
          />
          <input
            value={ghPass}
            onChange={(e) => setGhPass(e.target.value)}
            className="form-input"
            type="password"
            placeholder="Github Token"
          />
          <input
            value={broker}
            onChange={(e) => setBroker(e.target.value)}
            className={"form-input " + (error && "is-error")}
            type="text"
            placeholder="MQTT Broker Address"
          ></input>
          {error && (
            <label className="is-error">
              Please provide a port `hostname:port`
            </label>
          )}
          <button className="btn btn-primary" onClick={submitHandler}>
            Configure
          </button>
          {loading && <div className="loading" />}
        </div>
      </div>
    </div>
  );
}
