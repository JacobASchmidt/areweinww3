import React, { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [status, setStatus] = useState({
    status: "",
    textColor: "",
    subLine: "",
    explanation: "",
  });
  const [articles, setArticles] = useState([]);

  const fetchStatus = async () => {
    const response = await fetch(
      "http://localhost:3000/api/v1/status"
    );
    if (!response.ok) {
      console.error("Failed to fetch status");
      console.log(await response.text());
      return;
    }
    const data = await response.json();
    console.log("Status: ", data.status);
    console.log("Articles: ", data.articles);
    setStatus(data.status);
    setArticles(data.articles);
  };

  useEffect(() => {
    fetchStatus();
  }, []);

  return (
    <div className="container">
      <div className="status-column">
        <h1
          style={{
            color:
              status.status === "YES" ? "#e63946" : "#4caf50",
          }}
        >
          {status.status}
        </h1>
        <h2>{status.subLine}</h2>
        <p>{status.explanation}</p>
      </div>
      <div className="news-column">
        {articles.slice(0, 2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.headline}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
      <div className="news-column">
        {articles.slice(2).map((article, index) => (
          <div key={index} className="article">
            <h2>{article.headline}</h2>
            <p>{article.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
