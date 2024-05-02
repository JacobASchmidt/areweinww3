import React, { useEffect, useState } from "react";
import "./App.css";

function fetchStatus() {
  return fetch("api/v1/status").then((resp) => resp.json());
}

function App() {
  const [status, setStatus] = useState({
    status: "",
    textColor: "",
    subLine: "",
    explanation: "",
  });

  useEffect(() => {
    fetchStatus().then(setStatus);
  }, []);

  const articles = [
    {
      headline: "Article Title 1",
      description: "Detailed analysis of recent global events.",
    },
    {
      headline: "Article Title 2",
      description: "Expert opinions on geopolitical tensions.",
    },
    {
      headline: "Article Title 3",
      description:
        "Historical context behind today's conflicts.",
    },
    {
      headline: "Article Title 4",
      description:
        "Predictive insights into future global trends.",
    },
  ];

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
